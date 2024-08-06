package sample

import (
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
)

// ICreateSampleUseCase ユースケース: サンプルデータを追加 のインターフェース
// プレゼンテーション層は、ユースケース層で定義したこのインターフェースを介して処理を依頼する
type ICreateSampleUseCase interface {
	Run(req *application.CreateSampleRequest) (*application2.CreateSampleResponse, error)
}
