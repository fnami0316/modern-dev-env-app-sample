package sample

import (
	"reflect"
	"testing"

	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
)

func TestNewCreateSampleResponse(t *testing.T) {
	type args struct {
		sample *entity.Sample
	}
	tests := []struct {
		name    string
		args    args
		want    *CreateSampleResponse
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				sample: newSampleEntityForTest(t, "1", "sample_name"),
			},
			want: &CreateSampleResponse{
				sample: newSampleEntityForTest(t, "1", "sample_name"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCreateSampleResponse(tt.args.sample)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCreateSampleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateSampleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSampleResponse_validate(t *testing.T) {
	type fields struct {
		sample *entity.Sample
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]常にバリデーション通過",
			fields: fields{
				sample: newSampleEntityForTest(t, "1", "sample_name"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateSampleResponse{
				sample: tt.fields.sample,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateSampleResponse_Sample(t *testing.T) {
	type fields struct {
		sample *entity.Sample
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.Sample
	}{
		{
			name: "[OK]サンプルデータを取得できる",
			fields: fields{
				sample: newSampleEntityForTest(t, "1", "sample_name"),
			},
			want: newSampleEntityForTest(t, "1", "sample_name"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateSampleResponse{
				sample: tt.fields.sample,
			}
			if got := c.Sample(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sample() = %v, want %v", got, tt.want)
			}
		})
	}
}
