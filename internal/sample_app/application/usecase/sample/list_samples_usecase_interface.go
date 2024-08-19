//go:generate mockgen -source=list_samples_usecase_interface.go -destination=list_samples_usecase_mock.go -package=sample -mock_names=IListSamplesUseCase=MockListSamplesUseCase

package sample

import (
	"modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	sample2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
)

// IListSamplesUseCase ユースケース: サンプルデータのリストを取得 のインターフェース
// プレゼンテーション層は、ユースケース層で定義したこのインターフェースを介して処理を依頼する
type IListSamplesUseCase interface {
	Run(req *sample.ListSamplesRequest) (*sample2.ListSamplesResponse, error)
}
