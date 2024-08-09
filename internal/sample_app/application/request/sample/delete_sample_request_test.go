package sample

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

func TestNewDeleteSampleRequest(t *testing.T) {
	type args struct {
		id value.SampleID
	}
	tests := []struct {
		name    string
		args    args
		want    *DeleteSampleRequest
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				id: value.SampleID("1"),
			},
			want: &DeleteSampleRequest{
				id: value.SampleID("1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDeleteSampleRequest(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDeleteSampleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteSampleRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSampleRequest_validate(t *testing.T) {
	type fields struct {
		id value.SampleID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]常にバリデーション通過",
			fields: fields{
				id: value.SampleID("1"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DeleteSampleRequest{
				id: tt.fields.id,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteSampleRequest_ID(t *testing.T) {
	type fields struct {
		id value.SampleID
	}
	tests := []struct {
		name   string
		fields fields
		want   value.SampleID
	}{
		{
			name: "[OK]IDを取得できる",
			fields: fields{
				id: value.SampleID("1"),
			},
			want: value.SampleID("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DeleteSampleRequest{
				id: tt.fields.id,
			}
			if got := c.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
