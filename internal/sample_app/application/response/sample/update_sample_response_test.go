package sample

import (
	"reflect"
	"testing"

	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
)

func TestNewUpdateSampleResponse(t *testing.T) {
	type args struct {
		sample *entity.Sample
	}
	tests := []struct {
		name    string
		args    args
		want    *UpdateSampleResponse
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				sample: &entity.Sample{},
			},
			want: &UpdateSampleResponse{
				sample: &entity.Sample{},
			},
			wantErr: false,
		},
		{
			name: "[NG]バリデーションでエラー",
			args: args{
				sample: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUpdateSampleResponse(tt.args.sample)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUpdateSampleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateSampleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSampleResponse_validate(t *testing.T) {
	type fields struct {
		sample *entity.Sample
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]全てのバリデーションを通過",
			fields: fields{
				sample: &entity.Sample{},
			},
			wantErr: false,
		},
		{
			name: "[NG]sampleがnil",
			fields: fields{
				sample: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UpdateSampleResponse{
				sample: tt.fields.sample,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateSampleResponse_Sample(t *testing.T) {
	type fields struct {
		sample *entity.Sample
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.Sample
	}{
		{
			name: "[OK]sampleを取得できる",
			fields: fields{
				sample: &entity.Sample{},
			},
			want: &entity.Sample{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UpdateSampleResponse{
				sample: tt.fields.sample,
			}
			if got := c.Sample(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sample() = %v, want %v", got, tt.want)
			}
		})
	}
}
