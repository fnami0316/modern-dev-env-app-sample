package sample

import (
	"errors"
	"reflect"
	"testing"

	application3 "modern-dev-env-app-sample/internal/sample_app/application/repository"
	"modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	infrastructure2 "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm"
	infrastructure "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestNewUpdateSampleUseCase(t *testing.T) {
	type args struct {
		iCon        transaction.IConnection
		iSampleRepo application3.ISampleRepository
	}
	tests := []struct {
		name    string
		args    args
		want    *UpdateSampleUseCase
		wantErr bool
	}{
		{
			name: "[OK]インスタンスが生成できる",
			args: args{
				iCon:        &infrastructure.GORMConnection{},
				iSampleRepo: &infrastructure2.SampleRepository{},
			},
			want: &UpdateSampleUseCase{
				iCon:        &infrastructure.GORMConnection{},
				iSampleRepo: &infrastructure2.SampleRepository{},
			},
			wantErr: false,
		},
		{
			name: "[NG]バリデーションでエラー",
			args: args{
				iCon:        nil,
				iSampleRepo: &infrastructure2.SampleRepository{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUpdateSampleUseCase(tt.args.iCon, tt.args.iSampleRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUpdateSampleUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateSampleUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSampleUseCase_validate(t *testing.T) {
	type fields struct {
		iCon        transaction.IConnection
		iSampleRepo application3.ISampleRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]全てのバリデーションを通過",
			fields: fields{
				iCon:        &infrastructure.GORMConnection{},
				iSampleRepo: &infrastructure2.SampleRepository{},
			},
			wantErr: false,
		},
		{
			name: "[NG]iConがnil",
			fields: fields{
				iCon:        nil,
				iSampleRepo: &infrastructure2.SampleRepository{},
			},
			wantErr: true,
		},
		{
			name: "[OK]iSampleRepoがnil",
			fields: fields{
				iCon:        &infrastructure.GORMConnection{},
				iSampleRepo: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &UpdateSampleUseCase{
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
func TestUpdateSampleUseCase_Run(t *testing.T) {
	gormDB := createConnectionForTest(t)
	con := infrastructure.NewGORMConnection(gormDB)
	sampleRepo, err := infrastructure2.CreateSampleRepository(con)
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

	updatedSample1Name, err := value.NewSampleName("updated_sample1")
	if err != nil {
		t.Fatalf("failed to NewSampleName(): %v", err)
	}
	updatedSample1 := newSampleForTest(t, sample1.ID(), updatedSample1Name)

	type fields struct {
		iCon        transaction.IConnection
		iSampleRepo application3.ISampleRepository
	}
	type args struct {
		req *application.UpdateSampleRequest
	}
	tests := []struct {
		name         string
		setupSamples entity.Samples
		fields       fields
		args         args
		wantRes      *application2.UpdateSampleResponse
		wantSamples  entity.Samples
		wantErr      bool
	}{
		{
			name: "[OK]サンプルデータを更新できる",
			setupSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iCon:        con,
				iSampleRepo: sampleRepo,
			},
			args: args{
				req: newUpdateSampleRequestForTest(t, sample1.ID(), updatedSample1Name),
			},
			wantRes: newUpdateSampleResponseForTest(t, updatedSample1),
			wantSamples: entity.Samples{
				updatedSample1,
				sample2,
				sample3,
			},
			wantErr: false,
		},
		{
			name: "[NG]データストアに存在しないSampleエンティティのIDを指定した場合",
			setupSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iCon:        con,
				iSampleRepo: sampleRepo,
			},
			args: args{
				req: newUpdateSampleRequestForTest(t, sample4.ID(), updatedSample1Name),
			},
			wantRes: nil,
			wantSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			wantErr: true,
		},
		{
			name: "[NG]Save()でエラー",
			setupSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iCon: con,
				iSampleRepo: func() application3.ISampleRepository {
					mockSampleRepo.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).Return(entity.Samples{sample1}, nil)
					mockSampleRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("save error"))
					return mockSampleRepo
				}(),
			},
			args: args{
				req: newUpdateSampleRequestForTest(t, sample1.ID(), updatedSample1Name),
			},
			wantRes: nil,
			wantSamples: entity.Samples{
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

			l := &UpdateSampleUseCase{
				iCon:        tt.fields.iCon,
				iSampleRepo: tt.fields.iSampleRepo,
			}
			gotRes, err := l.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// レスポンスのチェック
			compareUpdateSampleResponse(t, gotRes, tt.wantRes)

			// エンティティが期待通り更新されているかチェック
			gotSamples, err := sampleRepo.FindAll(nil)
			if err != nil {
				t.Fatalf("failed to FindAll(): %v", err)
			}
			compareSamples(t, gotSamples, tt.wantSamples)
		})
	}
}

// newUpdateSampleRequestForTest UpdateSampleRequestを生成(エラーはテスト失敗として扱う)
func newUpdateSampleRequestForTest(t *testing.T, id value.SampleID, name value.SampleName) *application.UpdateSampleRequest {
	req, err := application.NewUpdateSampleRequest(id, name)
	if err != nil {
		t.Fatalf("failed to NewUpdateSampleRequest(): %v", err)
	}
	return req
}

// newUpdateSampleResponseForTest UpdateSampleResponseを生成(エラーはテスト失敗として扱う)
func newUpdateSampleResponseForTest(t *testing.T, sample *entity.Sample) *application2.UpdateSampleResponse {
	res, err := application2.NewUpdateSampleResponse(sample)
	if err != nil {
		t.Fatalf("failed to NewUpdateSampleResponse(): %v", err)
	}
	return res
}

// compareUpdateSampleResponse UpdateSampleResponseの比較
func compareUpdateSampleResponse(t *testing.T, got, want *application2.UpdateSampleResponse) {
	opts := []cmp.Option{
		cmp.AllowUnexported(entity.Sample{}, application2.UpdateSampleResponse{}),
	}
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Errorf("(-got +want)\n%s", diff)
	}
}
