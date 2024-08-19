package sample

import (
	"errors"
	"os"
	"reflect"
	"testing"

	application3 "modern-dev-env-app-sample/internal/sample_app/application/repository"
	application4 "modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	domain "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	infrastructure "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm"
	infrastructure2 "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestNewCreateSampleUseCase(t *testing.T) {
	type args struct {
		iCon        application4.IConnection
		iSampleRepo application3.ISampleRepository
	}
	tests := []struct {
		name    string
		args    args
		want    *CreateSampleUseCase
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				iCon:        &infrastructure2.GORMConnection{},
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			want: &CreateSampleUseCase{
				iCon:        &infrastructure2.GORMConnection{},
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			wantErr: false,
		},
		{
			name: "[NG]バリデーションでエラー",
			args: args{
				iCon:        nil, // エラー
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCreateSampleUseCase(tt.args.iCon, tt.args.iSampleRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCreateSampleUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateSampleUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSampleUseCase_validate(t *testing.T) {
	type fields struct {
		iCon        application4.IConnection
		iSampleRepo application3.ISampleRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]全てのバリデーション通過",
			fields: fields{
				iCon:        &infrastructure2.GORMConnection{},
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			wantErr: false,
		},
		{
			name: "[NG]iConがnil",
			fields: fields{
				iCon:        nil,
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			wantErr: true,
		},
		{
			name: "[NG]iSampleRepoがnil",
			fields: fields{
				iCon:        &infrastructure2.GORMConnection{},
				iSampleRepo: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &CreateSampleUseCase{
				iCon:        tt.fields.iCon,
				iSampleRepo: tt.fields.iSampleRepo,
			}
			if err := l.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// テストの前提条件
// - Spannerエミュレータが起動状態であり、spanner-emulator:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func TestCreateSampleUseCase_Run(t *testing.T) {
	gormDB := createConnectionForTest(t)
	con := infrastructure2.NewGORMConnection(gormDB)
	sampleRepo, err := infrastructure.CreateSampleRepository(con)
	if err != nil {
		t.Fatalf("failed to CreateSampleRepository(): %v", err)
	}
	ctrl := gomock.NewController(t)
	mockSampleRepo := application3.NewMockSampleRepository(ctrl)

	// 自動採番で生成したSampleエンティティ
	sample1 := createSampleForTest(t, "sample1")
	sample2 := createSampleForTest(t, "sample2")
	sample3 := createSampleForTest(t, "sample3")
	sample4 := createSampleForTest(t, "sample4")

	type fields struct {
		iCon        application4.IConnection
		iSampleRepo application3.ISampleRepository
	}
	type args struct {
		req *application.CreateSampleRequest
	}
	tests := []struct {
		name         string
		setupSamples domain.Samples
		fields       fields
		args         args
		wantRes      *application2.CreateSampleResponse
		wantSamples  domain.Samples
		wantErr      bool
	}{
		{
			name: "[OK]サンプルデータを作成できる",
			setupSamples: domain.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iCon:        con,
				iSampleRepo: sampleRepo,
			},
			args: args{
				req: newCreateSampleRequestForTest(t, "sample4"),
			},
			wantRes: newCreateSampleResponseForTest(t, sample4),
			wantSamples: domain.Samples{
				sample1,
				sample2,
				sample3,
				sample4,
			},
			wantErr: false,
		},
		{
			name: "[NG]Saveでエラー",
			setupSamples: domain.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iCon: con,
				iSampleRepo: func() application3.ISampleRepository {
					mockSampleRepo.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(
						func(sample *domain.Sample, iTx application4.ITransaction) error {
							return errors.New("dummy error")
						},
					)
					return mockSampleRepo
				}(),
			},
			args: args{
				req: newCreateSampleRequestForTest(t, "sample4"),
			},
			wantRes: nil,
			wantSamples: domain.Samples{
				sample1,
				sample2,
				sample3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllSamplesForTest(t)
			setupSamplesForTest(t, tt.setupSamples)
			defer deleteAllSamplesForTest(t)
			l := &CreateSampleUseCase{
				iCon:        tt.fields.iCon,
				iSampleRepo: tt.fields.iSampleRepo,
			}

			gotRes, err := l.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// レスポンスのチェック
			compareCreateSampleResponse(t, gotRes, tt.wantRes)

			// エンティティが期待通り永続化されているかチェック
			gotSamples, err := sampleRepo.FindAll(nil)
			if err != nil {
				t.Fatalf("failed to FindAll(): %v", err)
			}
			compareSamples(t, gotSamples, tt.wantSamples)
		})
	}
}

// createConnectionForTest テスト用のGORMコネクションを生成
// 利用するための前提条件
// - Spannerエミュレータが起動状態であり、spanner-emulator:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func createConnectionForTest(t *testing.T) *gorm.DB {
	err := os.Setenv("SPANNER_EMULATOR_HOST", "spanner-emulator:9010")
	if err != nil {
		t.Fatalf("failed to Setenv(): %v", err)
	}
	con, err := infrastructure.Setup(
		"local-project",
		"test-instance",
		"test-database",
	)
	if err != nil {
		t.Fatalf("failed to Setup(): %v", err)
	}
	return con
}

// newCreateSampleRequestForTest CreateSampleRequestを生成(エラーはテスト失敗として扱う)
func newCreateSampleRequestForTest(t *testing.T, name value.SampleName) *application.CreateSampleRequest {
	req, err := application.NewCreateSampleRequest(name)
	if err != nil {
		t.Fatalf("failed to NewCreateSampleRequest(): %v", err)
	}
	return req
}

// newCreateSampleResponseForTest CreateSampleResponseを生成(エラーはテスト失敗として扱う)
func newCreateSampleResponseForTest(t *testing.T, sample *domain.Sample) *application2.CreateSampleResponse {
	res, err := application2.NewCreateSampleResponse(sample)
	if err != nil {
		t.Fatalf("failed to NewCreateSampleResponse(): %v", err)
	}
	return res
}

// createSampleForTest ID自動採番でSampleエンティティを生成(エラーはテスト失敗として扱う)
func createSampleForTest(t *testing.T, name value.SampleName) *domain.Sample {
	sample, err := domain.CreateDefaultSample(name)
	if err != nil {
		t.Fatalf("failed to CreateDefaultSample(): %v", err)
	}
	return sample
}

// setupSamplesForTest テスト用データストア上に指定したSampleエンティティ群をセットアップ
// テスト用データストア: Spannerエミュレータ上のDB
// 利用するための前提条件
// - Spannerエミュレータが起動状態であり、spanner-emulator:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func setupSamplesForTest(t *testing.T, samples domain.Samples) {
	gormDB := createConnectionForTest(t)
	iCon := infrastructure2.NewGORMConnection(gormDB)
	sampleRepo, err := infrastructure.CreateSampleRepository(iCon)
	if err != nil {
		t.Fatalf("failed to CreateSampleRepository(): %v", err)
	}

	// リポジトリ経由でデータストアへ保存
	if err := iCon.Transaction(func(iTx application4.ITransaction) error {
		for _, sample := range samples {
			if err := sampleRepo.Save(sample, iTx); err != nil {
				t.Fatalf("failed to Save(): %v", err)
			}
		}
		return nil
	}); err != nil {
		t.Fatalf("failed to Transaction(): %v", err)
	}

	// 与えられたエンティティのみがデータストア上に存在することを確認
	gotSamples, err := sampleRepo.FindAll(nil)
	if err != nil {
		t.Fatalf("failed to FindAll(): %v", err)
	}
	compareSamples(t, gotSamples, samples)
}

// deleteAllSamplesForTest テスト用データストア上から全てのSampleエンティティ群を削除
// 利用するための前提条件
// - Spannerエミュレータが起動状態であり、spanner-emulator:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func deleteAllSamplesForTest(t *testing.T) {
	gormDB := createConnectionForTest(t)
	iCon := infrastructure2.NewGORMConnection(gormDB)
	sampleRepo, err := infrastructure.CreateSampleRepository(iCon)
	if err != nil {
		t.Fatalf("failed to CreateSampleRepository(): %v", err)
	}

	// データストア上のSampleエンティティを全て取得
	allSamples, err := sampleRepo.FindAll(nil)
	if err != nil {
		t.Fatalf("failed to FindAll(): %v", err)
	}

	// データストア上のSampleエンティティを全て削除
	for _, sample := range allSamples {
		err := sampleRepo.Delete(sample, nil)
		if err != nil {
			t.Fatalf("failed to Delete(): %v", err)
		}
	}

	// 削除されたことを確認
	gotSamples, err := sampleRepo.FindAll(nil)
	if err != nil {
		t.Fatalf("failed to FindAll(): %v", err)
	}
	if len(gotSamples) != 0 {
		t.Fatalf("gotSamples = %v, want nil", gotSamples)
	}
}

// compareCreateSampleResponse CreateSampleResponseの比較
func compareCreateSampleResponse(t *testing.T, got, want *application2.CreateSampleResponse) {
	opts := []cmp.Option{
		// IDはランダム採番なので比較対象外とする
		cmpopts.IgnoreFields(domain.Sample{}, "id"),
		cmp.AllowUnexported(
			application2.CreateSampleResponse{},
			domain.Sample{},
		),
	}
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Errorf("(-got +want)\n%s", diff)
	}
}

// compareSamples Sampleエンティティ群の比較
func compareSamples(t *testing.T, got, want domain.Samples) {
	opts := []cmp.Option{
		cmpopts.IgnoreFields(domain.Sample{}, "id"),
		cmp.AllowUnexported(domain.Sample{}),
		// 自動採番であるIDは結果の比較においては順番の制御が難しいため、名前でソートして比較
		cmpopts.SortSlices(func(a, b *domain.Sample) bool {
			return a.Name() < b.Name()
		}),
	}
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Errorf("(-got +want)\n%s", diff)
	}
}
