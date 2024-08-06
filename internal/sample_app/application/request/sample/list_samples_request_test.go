package sample

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

func TestNewListSamplesRequest(t *testing.T) {
	type args struct {
		ids []value.SampleID
	}
	tests := []struct {
		name    string
		args    args
		want    *ListSamplesRequest
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				ids: []value.SampleID{"1", "2"},
			},
			want: &ListSamplesRequest{
				ids: []value.SampleID{"1", "2"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewListSamplesRequest(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewListSamplesRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListSamplesRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListSamplesRequest_validate(t *testing.T) {
	type fields struct {
		ids []value.SampleID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "常に成功",
			fields: fields{
				ids: []value.SampleID{"1", "2"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListSamplesRequest{
				ids: tt.fields.ids,
			}
			if err := l.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListSamplesRequest_IDs(t *testing.T) {
	type fields struct {
		ids []value.SampleID
	}
	tests := []struct {
		name   string
		fields fields
		want   []value.SampleID
	}{
		{
			name: "[OK]IDのリストを取得できる",
			fields: fields{
				ids: []value.SampleID{"1", "2"},
			},
			want: []value.SampleID{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListSamplesRequest{
				ids: tt.fields.ids,
			}
			if got := l.IDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
