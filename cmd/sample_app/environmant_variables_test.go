package main

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadEnvironmentVariables(t *testing.T) {
	tests := []struct {
		name    string
		prepare func()
		want    *EnvironmentVariables
		wantErr bool
	}{
		{
			name: "[OK]環境変数を読み込むことができる",
			prepare: func() {
				_ = os.Setenv("PORT", "8080")
				_ = os.Setenv("GCP_PROJECT_ID", "test")
				_ = os.Setenv("SPANNER_INSTANCE_ID", "test")
				_ = os.Setenv("SPANNER_DATABASE_ID", "test")
			},
			want: &EnvironmentVariables{
				Port:              8080,
				GCPProjectID:      "test",
				SpannerInstanceID: "test",
				SpannerDatabaseID: "test",
			},
			wantErr: false,
		},
		{
			name: "[NG]バリデーションエラー",
			prepare: func() {
				_ = os.Setenv("PORT", "8080")
				_ = os.Setenv("GCP_PROJECT_ID", "") // エラー
				_ = os.Setenv("SPANNER_INSTANCE_ID", "test")
				_ = os.Setenv("SPANNER_DATABASE_ID", "test")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := LoadEnvironmentVariables()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadEnvironmentVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadEnvironmentVariables() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariables_validate(t *testing.T) {
	type fields struct {
		Port              int
		GCPProjectID      string
		SpannerInstanceID string
		SpannerDatabaseID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]バリデーション全てを通過",
			fields: fields{
				Port:              1,
				GCPProjectID:      "test",
				SpannerInstanceID: "test",
				SpannerDatabaseID: "test",
			},
			wantErr: false,
		},
		{
			name: "[NG]Portがゼロ値の場合",
			fields: fields{
				Port:              0,
				GCPProjectID:      "test",
				SpannerInstanceID: "test",
				SpannerDatabaseID: "test",
			},
			wantErr: true,
		},
		{
			name: "[NG]GCPProjectIDがゼロ値の場合",
			fields: fields{
				Port:              1,
				GCPProjectID:      "",
				SpannerInstanceID: "test",
				SpannerDatabaseID: "test",
			},
			wantErr: true,
		},
		{
			name: "[NG]SpannerInstanceIDがゼロ値の場合",
			fields: fields{
				Port:              1,
				GCPProjectID:      "test",
				SpannerInstanceID: "",
				SpannerDatabaseID: "test",
			},
			wantErr: true,
		},
		{
			name: "[NG]SpannerDatabaseIDがゼロ値の場合",
			fields: fields{
				Port:              1,
				GCPProjectID:      "test",
				SpannerInstanceID: "test",
				SpannerDatabaseID: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EnvironmentVariables{
				Port:              tt.fields.Port,
				GCPProjectID:      tt.fields.GCPProjectID,
				SpannerInstanceID: tt.fields.SpannerInstanceID,
				SpannerDatabaseID: tt.fields.SpannerDatabaseID,
			}
			if err := e.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
