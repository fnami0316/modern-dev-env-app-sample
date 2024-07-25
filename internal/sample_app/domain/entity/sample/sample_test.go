package sample

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

func TestNewSample(t *testing.T) {
	type args struct {
		id   value.SampleID
		name value.SampleName
	}
	tests := []struct {
		name    string
		args    args
		want    *Sample
		wantErr bool
	}{
		{
			name: "[OK]バリデーションでエラーがない場合、インスタンスを生成できる",
			args: args{
				id:   1,
				name: "name",
			},
			want: &Sample{
				id:   1,
				name: "name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSample(tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSample() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSample_validate(t *testing.T) {
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
			name:    "[OK]常にエラーは発生しない",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if err := s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSample_ID(t *testing.T) {
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
			name: "[OK]IDを取得できる",
			fields: fields{
				id: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := s.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSample_Name(t *testing.T) {
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
			name: "[OK]名前を取得できる",
			fields: fields{
				name: "name",
			},
			want: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSample_Update(t *testing.T) {
	type fields struct {
		id   value.SampleID
		name value.SampleName
	}
	type args struct {
		name value.SampleName
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Sample
		wantErr bool
	}{
		{
			name: "[OK]バリデーションでエラーがない場合、更新後のフィールド値を持つインスタンスを生成できる",
			fields: fields{
				id:   1,
				name: "name",
			},
			args: args{
				name: "updated name",
			},
			want: &Sample{
				id:   1,
				name: "updated name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			got, err := s.Update(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSample_update(t *testing.T) {
	type fields struct {
		id   value.SampleID
		name value.SampleName
	}
	type args struct {
		name value.SampleName
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Sample
		wantErr bool
	}{
		{
			name: "[OK]バリデーションでエラーがない場合、更新後のフィールド値を持つインスタンスを生成できる",
			fields: fields{
				id:   1,
				name: "name",
			},
			args: args{
				name: "updated name",
			},
			want: &Sample{
				id:   1,
				name: "updated name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sample{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			got, err := s.update(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
