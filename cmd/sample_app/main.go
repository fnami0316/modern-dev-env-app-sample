package main

import (
	"fmt"
	"log"
	"net"

	pb "modern-dev-env-app-sample/internal/sample_app/pb/api/proto"
	"modern-dev-env-app-sample/internal/sample_app/service/sample"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenProtocol := "tcp"
	listenPortNum := 53000 // ポート番号

	// 1. 指定したプロトコル・ポートのListenerを作成
	// (net.Listenerが返される。Listenerとはポートに対して聞き耳を立てる人である)。
	portListener, err := net.Listen(listenProtocol, fmt.Sprintf(":%d", listenPortNum))
	// Listenに失敗したならばプログラムを即終了
	if err != nil {
		log.Fatalf("プロトコル:%v, ポート:%v のListener作成に失敗しました: %v", listenProtocol, listenPortNum, err)
	}

	// 2. gRPCサーバのインスタンスを生成
	// (grpc.Serverインスタンスのポインタが返ってくる)
	grpcServer := grpc.NewServer()

	// 3. gRPCサーバにサービスを登録
	// Register<サービス名>Server(grpc.Serverのポインタ, <サービス名>Serverのポインタ)関数は、
	// protocコマンド実行で生成された「<元になった.proroファイル名>_grpc.pb.go」内に自動で定義されている
	// ここで登録されたサービスについてのみAPIが使えるようになる
	pb.RegisterSampleServiceServer(grpcServer, &sample.SampleServiceServer{})

	// 4. gRPCサーバのServer Reflectionを有効にする
	// (「grpc_cli」コマンドで、gRPCサーバに登録したサービスのRPCメソッドをシリアライズなしで実行可能になる)
	reflection.Register(grpcServer)

	// 5. gRPCサーバーを起動(指定したプロトコル・ポートのListenも開始)
	grpcServer.Serve(portListener)
}
