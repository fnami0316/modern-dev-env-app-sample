package main

import (
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
)

// useCases 全ユースケースインスタンスをまとめた構造体
type useCases struct {
	listSamplesUseCase *sample.ListSamplesUseCase
}

// newUseCases コンストラクタ
func newUseCases(
	listSamplesUseCase *sample.ListSamplesUseCase,
) (*useCases, error) {
	return &useCases{
		listSamplesUseCase: listSamplesUseCase,
	}, nil
}

// createUsesCases 全ユースケースインスタンスのファクトリ
func createUseCases(
	infras *infrastructures,
) (*useCases, error) {
	// ユースケース層のインスタンス生成
	listSamplesUseCase, err := sample.NewListSamplesUseCase(infras.iSampleRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to NewListSamplesUseCase(): %w", err)
	}
	return newUseCases(listSamplesUseCase)
}
