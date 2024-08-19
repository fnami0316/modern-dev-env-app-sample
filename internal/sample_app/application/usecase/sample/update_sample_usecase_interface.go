//go:generate mockgen -source=update_sample_usecase_interface.go -destination=update_sample_usecase_mock.go -package=sample -mock_names=IUpdateSampleUseCase=MockUpdateSampleUseCase

package sample

import (
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
)

// IUpdateSampleUseCase ユースケース: サンプルデータを更新 のインターフェース
type IUpdateSampleUseCase interface {
	Run(req *application.UpdateSampleRequest) (*application2.UpdateSampleResponse, error)
}
