package sample

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/application/repository"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	infrastructure "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm"
	infrastructure2 "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewListSamplesUseCase(t *testing.T) {
	type args struct {
		iSampleRepo repository.ISampleRepository
	}
	tests := []struct {
		name    string
		args    args
		want    *ListSamplesUseCase
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			want: &ListSamplesUseCase{
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			wantErr: false,
		},
		{
			name: "[NG]バリデーションでエラー",
			args: args{
				iSampleRepo: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewListSamplesUseCase(tt.args.iSampleRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewListSamplesUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListSamplesUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// テストの前提条件
// - Spannerエミュレータが起動状態であり、spanner-emulator:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func TestListSamplesUseCase_Run(t *testing.T) {
	gorm := createConnectionForTest(t)
	con := infrastructure2.NewGORMConnection(gorm)
	sampleRepo, err := infrastructure.CreateSampleRepository(con)
	if err != nil {
		t.Fatalf("failed to CreateSampleRepository(): %v", err)
	}

	// 自動採番で生成したSampleエンティティ
	sample1 := newSampleForTest(t, "sample1", "sample1")
	sample2 := newSampleForTest(t, "sample2", "sample2")
	sample3 := newSampleForTest(t, "sample3", "sample3")

	type fields struct {
		iSampleRepo repository.ISampleRepository
	}
	type args struct {
		req *application2.ListSamplesRequest
	}
	tests := []struct {
		name         string
		setupSamples entity.Samples
		fields       fields
		args         args
		wantRes      *application.ListSamplesResponse
		wantErr      bool
	}{
		{
			name: "[OK]リクエストで1つ以上のID(重複なし)を指定した場合、指定したIDのエンティティが全て取得できる",
			setupSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iSampleRepo: sampleRepo,
			},
			args: args{
				req: newListSamplesRequestForTest(t, value.SampleIDs{"sample1", "sample3"}),
			},
			wantRes: newListSamplesResponseForTest(t, entity.Samples{sample1, sample3}),
			wantErr: false,
		},
		{
			name: "[OK]リクエストで1つ以上のID(重複あり)を指定した場合、重複は1つにまとめられた上で指定したIDのエンティティが全て取得できる",
			setupSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iSampleRepo: sampleRepo,
			},
			args: args{
				req: newListSamplesRequestForTest(t, value.SampleIDs{"sample1", "sample3", "sample1"}),
			},
			wantRes: newListSamplesResponseForTest(t, entity.Samples{sample1, sample3}),
			wantErr: false,
		},
		{
			name: "[OK]リクエストで1つ以上のID(データストアに存在しないものを含む)を指定した場合、存在しないものは除外された上で指定したIDのエンティティが全て取得できる",
			setupSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iSampleRepo: sampleRepo,
			},
			args: args{
				req: newListSamplesRequestForTest(t, value.SampleIDs{"sample1", "sample3", "sample4"}),
			},
			wantRes: newListSamplesResponseForTest(t, entity.Samples{sample1, sample3}),
			wantErr: false,
		},
		{
			name: "[NG]リクエストで1つもIDを指定していない場合、エラーが返る",
			setupSamples: entity.Samples{
				sample1,
				sample2,
				sample3,
			},
			fields: fields{
				iSampleRepo: sampleRepo,
			},
			args: args{
				req: newListSamplesRequestForTest(t, nil),
			},
			wantRes: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllSamplesForTest(t)
			setupSamplesForTest(t, tt.setupSamples)
			defer deleteAllSamplesForTest(t)

			l := &ListSamplesUseCase{
				iSampleRepo: tt.fields.iSampleRepo,
			}
			got, err := l.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			compareListSamplesResponse(t, got, tt.wantRes)
		})
	}
}

func TestListSamplesUseCase_validate(t *testing.T) {
	type fields struct {
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
				iSampleRepo: &infrastructure.SampleRepository{},
			},
			wantErr: false,
		},
		{
			name: "[NG]iSampleRepoがnil",
			fields: fields{
				iSampleRepo: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListSamplesUseCase{
				iSampleRepo: tt.fields.iSampleRepo,
			}
			if err := l.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// newSampleForTest Sampleエンティティを生成(エラーはテスト失敗として扱う)
func newSampleForTest(t *testing.T, id value.SampleID, name value.SampleName) *entity.Sample {
	sample, err := entity.NewSample(id, name)
	if err != nil {
		t.Fatalf("failed to NewSample(): %v", err)
	}
	return sample
}

// newListSamplesRequestForTest ListSamplesRequestを生成(エラーはテスト失敗として扱う)
func newListSamplesRequestForTest(t *testing.T, ids value.SampleIDs) *application2.ListSamplesRequest {
	req, err := application2.NewListSamplesRequest(ids)
	if err != nil {
		t.Fatalf("failed to NewListSamplesRequest(): %v", err)
	}
	return req
}

// newListSamplesResponseForTest ListSamplesResponseを生成(エラーはテスト失敗として扱う)
func newListSamplesResponseForTest(t *testing.T, samples entity.Samples) *application.ListSamplesResponse {
	res, err := application.NewListSamplesResponse(samples)
	if err != nil {
		t.Fatalf("failed to NewListSamplesResponse(): %v", err)
	}
	return res
}

// compareListSamplesResponse ListSamplesResponseの比較
func compareListSamplesResponse(t *testing.T, got, want *application.ListSamplesResponse) {
	opts := []cmp.Option{
		cmp.AllowUnexported(
			application.ListSamplesResponse{},
			entity.Sample{},
		),

		cmpopts.SortSlices(func(a, b *entity.Sample) bool {
			return a.ID() < b.ID()
		}),
	}
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Errorf("(-got +want)\n%s", diff)
	}
}
