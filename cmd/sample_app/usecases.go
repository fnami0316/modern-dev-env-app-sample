package main

import (
	"fmt"

	application "modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
)

// useCases 全ユースケースインスタンスをまとめた構造体
type useCases struct {
	iListSamplesUseCase  application.IListSamplesUseCase
	iCreateSampleUseCase application.ICreateSampleUseCase
	iUpdateSampleUseCase application.IUpdateSampleUseCase
	iDeleteSampleUseCase application.IDeleteSampleUseCase
}

// newUseCases コンストラクタ
func newUseCases(
	iListSampleUseCase application.IListSamplesUseCase,
	iCreateSampleUseCase application.ICreateSampleUseCase,
	iUpdateSampleUseCase application.IUpdateSampleUseCase,
	iDeleteSampleUseCase application.IDeleteSampleUseCase,
) (*useCases, error) {
	return &useCases{
		iListSamplesUseCase:  iListSampleUseCase,
		iCreateSampleUseCase: iCreateSampleUseCase,
		iUpdateSampleUseCase: iUpdateSampleUseCase,
		iDeleteSampleUseCase: iDeleteSampleUseCase,
	}, nil
}

// createUsesCases 全ユースケースインスタンスのファクトリ
func createUseCases(
	infras *infrastructures,
) (*useCases, error) {
	// ユースケース層のインスタンス生成
	listSamplesUseCase, err := application.NewListSamplesUseCase(infras.iSampleRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to NewListSamplesUseCase(): %w", err)
	}
	creatSampleUseCase, err := application.NewCreateSampleUseCase(infras.iCon, infras.iSampleRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to NewCreateSampleUseCase(): %w", err)
	}
	updateSampleUseCase, err := application.NewUpdateSampleUseCase(infras.iCon, infras.iSampleRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to NewUpdateSampleUseCase(): %w", err)
	}
	deleteSampleUseCase, err := application.NewDeleteSampleUseCase(infras.iCon, infras.iSampleRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to NewDeleteSampleUseCase(): %w", err)
	}
	return newUseCases(
		listSamplesUseCase,
		creatSampleUseCase,
		updateSampleUseCase,
		deleteSampleUseCase,
	)
}
