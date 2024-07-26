package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// EnvironmentVariables 環境変数
type EnvironmentVariables struct {
	// Port リスンポート番号
	Port int `envconfig:"GCP_PROJECT_ID" default:"8080"`
	// GCPProjectID GCPプロジェクトID
	GCPProjectID string `envconfig:"GCP_PROJECT_ID"`
	// SpannerInstanceID SpannerインスタンスID
	SpannerInstanceID string `envconfig:"SPANNER_INSTANCE_ID"`
	// SpannerDatabaseID SpannerデータベースID
	SpannerDatabaseID string `envconfig:"SPANNER_DATABASE_ID"`
}

// LoadEnvironmentVariables 環境変数を読み込む
func LoadEnvironmentVariables() EnvironmentVariables {
	var envVars EnvironmentVariables
	if err := envconfig.Process("", &envVars); err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
	}
	return envVars
}
