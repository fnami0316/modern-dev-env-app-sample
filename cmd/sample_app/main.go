package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 環境変数を読み込む
	envVars := LoadEnvironmentVariables()

	// インフラ層のインスタンスを生成
	infrastructures, err := createInfrastructuresWithGORMSpanner(
		envVars.GCPProjectID,
		envVars.SpannerInstanceID,
		envVars.SpannerDatabaseID,
	)
	if err != nil {
		log.Fatalf("failed to createInfrastructuresWithGORMSpanner(): %v", err)
	}

	// ユースケース層のインスタンスを生成
	useCases, err := createUseCases(infrastructures)
	if err != nil {
		log.Fatalf("failed to createUseCases(): %v", err)
	}

	// プレゼンテーション層のインスタンスを生成
	presentations, err := createPresentations(useCases)
	if err != nil {
		log.Fatalf("failed to createPresentations(): %v", err)
	}

	// gRPCサーバーの起動
	if err := startGrpcServer(envVars.Port, presentations); err != nil {
		log.Fatalf("failed to startGrpcServer(): %v", err)
	}
}

// startGrpcServer gRPCサーバーの起動処理
func startGrpcServer(port int, presentations *presentations) error {
	// 1. 指定したプロトコル・ポートのListenerを作成
	// (net.Listenerが返される。Listenerとはポートに対して聞き耳を立てる人である)。
	listenProtocol := "tcp"
	portListener, err := net.Listen(listenProtocol, fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("プロトコル:%v, ポート:%v のListener作成に失敗しました: %w", listenProtocol, port, err)
	}

	// 2. gRPCサーバのインスタンスを生成
	// (grpc.Serverインスタンスのポインタが返ってくる)
	grpcServer := grpc.NewServer()

	// 3. gRPCサーバにサービスを登録
	presentations.registerProtocServices(grpcServer)

	// 4. gRPCサーバのServer Reflectionを有効にする
	// (「grpcurl」コマンドで、gRPCサーバに登録したサービスのRPCメソッドをシリアライズなしで実行可能になる)
	reflection.Register(grpcServer)

	// 5. gRPCサーバーを起動(指定したプロトコル・ポートのListenも開始)
	if err := grpcServer.Serve(portListener); err != nil {
		return fmt.Errorf("failed to grpcServer.Serve(): %w", err)
	}
	return nil
}
