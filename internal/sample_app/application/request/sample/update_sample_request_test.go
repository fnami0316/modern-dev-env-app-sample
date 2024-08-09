package sample

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

func TestNewUpdateSampleRequest(t *testing.T) {
	type args struct {
		id   value.SampleID
		name value.SampleName
	}
	tests := []struct {
		name    string
		args    args
		want    *UpdateSampleRequest
		wantErr bool
	}{
		{
			name: "インスタンスを生成できる",
			args: args{
				id:   value.SampleID("1"),
				name: value.SampleName("name"),
			},
			want: &UpdateSampleRequest{
				id:   value.SampleID("1"),
				name: value.SampleName("name"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUpdateSampleRequest(tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUpdateSampleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateSampleRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSampleRequest_validate(t *testing.T) {
	type fields struct {
		id   value.SampleID
		name value.SampleName
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:   "[OK]常にバリデーション通過",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UpdateSampleRequest{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateSampleRequest_ID(t *testing.T) {
	type fields struct {
		id   value.SampleID
		name value.SampleName
	}
	tests := []struct {
		name   string
		fields fields
		want   value.SampleID
	}{
		{
			name:   "IDを取得できる",
			fields: fields{id: value.SampleID("1")},
			want:   value.SampleID("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UpdateSampleRequest{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := c.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSampleRequest_Name(t *testing.T) {
	type fields struct {
		id   value.SampleID
		name value.SampleName
	}
	tests := []struct {
		name   string
		fields fields
		want   value.SampleName
	}{
		{
			name:   "名前を取得できる",
			fields: fields{name: value.SampleName("name")},
			want:   value.SampleName("name"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UpdateSampleRequest{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
