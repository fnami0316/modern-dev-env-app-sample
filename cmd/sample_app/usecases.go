package main

import (
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
)

// useCases 全ユースケースインスタンスをまとめた構造体
type useCases struct {
	iListSamplesUseCase  sample.IListSamplesUseCase
	iCreateSampleUseCase sample.ICreateSampleUseCase
}

// newUseCases コンストラクタ
func newUseCases(
	iListSampleUseCase sample.IListSamplesUseCase,
	iCreateSampleUseCase sample.ICreateSampleUseCase,
) (*useCases, error) {
	return &useCases{
		iListSamplesUseCase:  iListSampleUseCase,
		iCreateSampleUseCase: iCreateSampleUseCase,
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
	creatSampleUseCase, err := sample.NewCreateSampleUseCase(infras.iCon, infras.iSampleRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to NewCreateSampleUseCase(): %w", err)
	}
	return newUseCases(
		listSamplesUseCase,
		creatSampleUseCase,
	)
}
