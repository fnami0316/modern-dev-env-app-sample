package sample

import (
	"reflect"
	"testing"

	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

func TestNewListSamplesResponse(t *testing.T) {
	sampleEntity1 := newSampleEntityForTest(t, "1", "sample_name")
	sampleEntity2 := newSampleEntityForTest(t, "2", "sample_name")

	type args struct {
		samples entity.Samples
	}
	tests := []struct {
		name    string
		args    args
		want    *ListSamplesResponse
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				samples: entity.Samples{
					sampleEntity1,
					sampleEntity2,
				},
			},
			want: &ListSamplesResponse{
				samples: entity.Samples{
					sampleEntity1,
					sampleEntity2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewListSamplesResponse(tt.args.samples)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewListSamplesResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListSamplesResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListSamplesResponse_validate(t *testing.T) {
	type fields struct {
		samples entity.Samples
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]常にバリデーション通過",
			fields: fields{
				samples: entity.Samples{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListSamplesResponse{
				samples: tt.fields.samples,
			}
			if err := l.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListSamplesResponse_Samples(t *testing.T) {
	sampleEntity1 := newSampleEntityForTest(t, "1", "sample_name")
	sampleEntity2 := newSampleEntityForTest(t, "2", "sample_name")

	type fields struct {
		samples entity.Samples
	}
	tests := []struct {
		name   string
		fields fields
		want   entity.Samples
	}{
		{
			name: "[OK]サンプルデータのリストを取得",
			fields: fields{
				samples: entity.Samples{
					sampleEntity1,
					sampleEntity2,
				},
			},
			want: entity.Samples{
				sampleEntity1,
				sampleEntity2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListSamplesResponse{
				samples: tt.fields.samples,
			}
			if got := l.Samples(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Samples() = %v, want %v", got, tt.want)
			}
		})
	}
}

// newSampleEntityForTest Sampleエンティティの生成(エラーは返さずテスト失敗扱いとする)
func newSampleEntityForTest(t *testing.T, id value.SampleID, name value.SampleName) *entity.Sample {
	sampleEntity, err := entity.NewSample(id, name)
	if err != nil {
		t.Fatalf("failed to NewSample(): %v", err)
	}
	return sampleEntity
}
