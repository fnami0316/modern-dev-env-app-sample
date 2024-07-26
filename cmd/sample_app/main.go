package main

import (
	"fmt"
	"log"
	"net"

	infrarepo "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm"
	"modern-dev-env-app-sample/internal/sample_app/presentation"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenProtocol := "tcp"
	envVars := LoadEnvironmentVariables()

	// 1. 指定したプロトコル・ポートのListenerを作成
	// (net.Listenerが返される。Listenerとはポートに対して聞き耳を立てる人である)。
	portListener, err := net.Listen(listenProtocol, fmt.Sprintf(":%d", envVars.Port))
	// Listenに失敗したならばプログラムを即終了
	if err != nil {
		log.Fatalf("プロトコル:%v, ポート:%v のListener作成に失敗しました: %v", listenProtocol, envVars.Port, err)
	}

	// 2. gRPCサーバのインスタンスを生成
	// (grpc.Serverインスタンスのポインタが返ってくる)
	grpcServer := grpc.NewServer()

	// インフラ層にリポジトリ利用のためのセットアップを任せる
	db, err := infrarepo.Setup(envVars.GCPProjectID, envVars.SpannerInstanceID, envVars.SpannerDatabaseID)
	if err != nil {
		log.Fatalf("failed to Setup(): %v", err)
	}

	// 3. gRPCサーバにサービスを登録
	// Register<サービス名>Server(grpc.Serverのポインタ, <サービス名>Serverのポインタ)関数は、
	// protocコマンド実行で生成された「<元になった.proroファイル名>_grpc.pb.go」内に自動で定義されている
	// ここで登録されたサービスについてのみAPIが使えるようになる
	if err := presentation.RegisterRPCServices(grpcServer, db); err != nil {
		log.Fatalf("failed to RegisterRPCServices(): %v", err)
	}

	// 4. gRPCサーバのServer Reflectionを有効にする
	// (「grpcurl」コマンドで、gRPCサーバに登録したサービスのRPCメソッドをシリアライズなしで実行可能になる)
	reflection.Register(grpcServer)

	// 5. gRPCサーバーを起動(指定したプロトコル・ポートのListenも開始)
	if err := grpcServer.Serve(portListener); err != nil {
		log.Fatalf("failed to grpcServer.Serve(): %v", err)
	}
}

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
