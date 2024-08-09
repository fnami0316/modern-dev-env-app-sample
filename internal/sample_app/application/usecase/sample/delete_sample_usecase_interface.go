package sample

import (
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
)

// IDeleteSampleUseCase ユースケース: サンプルデータを削除 のインターフェース
type IDeleteSampleUseCase interface {
	Run(req *application.DeleteSampleRequest) (*application2.DeleteSampleResponse, error)
}
