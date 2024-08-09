package sample

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/application/repository"
	"modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	infrastructure2 "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm"
	infrastructure "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"

	"github.com/google/go-cmp/cmp"
)

func TestNewDeleteSampleUseCase(t *testing.T) {
	type args struct {
		iCon        transaction.IConnection
		iSampleRepo repository.ISampleRepository
	}
	tests := []struct {
		name    string
		args    args
		want    *DeleteSampleUseCase
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				iCon:        &infrastructure.GORMConnection{},
				iSampleRepo: &infrastructure2.SampleRepository{},
			},
			want: &DeleteSampleUseCase{
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
			got, err := NewDeleteSampleUseCase(tt.args.iCon, tt.args.iSampleRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDeleteSampleUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteSampleUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSampleUseCase_validate(t *testing.T) {
	type fields struct {
		iCon        transaction.IConnection
		iSampleRepo repository.ISampleRepository
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
			name: "[NG]iSampleRepoがnil",
			fields: fields{
				iCon:        &infrastructure.GORMConnection{},
				iSampleRepo: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DeleteSampleUseCase{
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
func TestDeleteSampleUseCase_Run(t *testing.T) {
	gormDB := createConnectionForTest(t)
	con := infrastructure.NewGORMConnection(gormDB)
	sampleRepo, err := infrastructure2.CreateSampleRepository(con)
	if err != nil {
		t.Fatalf("failed to CreateSampleRepository(): %v", err)
	}

	// 自動採番で生成したSampleエンティティ
	sample1 := createSampleForTest(t, "sample1")
	sample2 := createSampleForTest(t, "sample2")
	sample3 := createSampleForTest(t, "sample3")
	sample4 := createSampleForTest(t, "sample4")

	type fields struct {
		iCon        transaction.IConnection
		iSampleRepo repository.ISampleRepository
	}
	type args struct {
		req *application.DeleteSampleRequest
	}
	tests := []struct {
		name         string
		setupSamples entity.Samples
		fields       fields
		args         args
		wantRes      *application2.DeleteSampleResponse
		wantSamples  entity.Samples
		wantErr      bool
	}{
		{
			name: "[OK]サンプルデータを削除できる",
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
				req: newDeleteSampleRequestForTest(t, sample1.ID()),
			},
			wantRes: &application2.DeleteSampleResponse{},
			wantSamples: entity.Samples{
				sample2,
				sample3,
			},
			wantErr: false,
		},
		{
			name: "[NG]データストアに存在しないSampleエンティティのIDを指定した場合エラー",
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
				req: newDeleteSampleRequestForTest(t, sample4.ID()),
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

			l := &DeleteSampleUseCase{
				iCon:        tt.fields.iCon,
				iSampleRepo: tt.fields.iSampleRepo,
			}
			gotRes, err := l.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// レスポンスのチェック
			if diff := cmp.Diff(gotRes, tt.wantRes); diff != "" {
				t.Errorf("(-got +want)\n%s", diff)
			}

			// エンティティが期待通り削除されているかチェック
			gotSamples, err := sampleRepo.FindAll(nil)
			if err != nil {
				t.Fatalf("failed to FindAll(): %v", err)
			}
			compareSamples(t, gotSamples, tt.wantSamples)
		})
	}
}

// newDeleteSampleRequestForTest テスト用のDeleteSampleRequestを生成(エラーはテスト失敗として扱う)
func newDeleteSampleRequestForTest(t *testing.T, id value.SampleID) *application.DeleteSampleRequest {
	req, err := application.NewDeleteSampleRequest(id)
	if err != nil {
		t.Fatalf("failed to NewDeleteSampleRequest(): %v", err)
	}
	return req
}
