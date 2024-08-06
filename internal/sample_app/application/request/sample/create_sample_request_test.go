package sample

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

func TestNewCreateSampleRequest(t *testing.T) {
	type args struct {
		name value.SampleName
	}
	tests := []struct {
		name    string
		args    args
		want    *CreateSampleRequest
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				name: "name",
			},
			want: &CreateSampleRequest{
				name: "name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCreateSampleRequest(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCreateSampleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateSampleRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSampleRequest_validate(t *testing.T) {
	type fields struct {
		name value.SampleName
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "常に成功",
			fields: fields{
				name: "name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateSampleRequest{
				name: tt.fields.name,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateSampleRequest_Name(t *testing.T) {
	type fields struct {
		name value.SampleName
	}
	tests := []struct {
		name   string
		fields fields
		want   value.SampleName
	}{
		{
			name: "名前を取得できる",
			fields: fields{
				name: "name",
			},
			want: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateSampleRequest{
				name: tt.fields.name,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
