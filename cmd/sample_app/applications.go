package main

import (
	"fmt"

	application "modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
)

// applications アプリケーション層の全インスタンスをまとめた構造体
type applications struct {
	iListSamplesUseCase  application.IListSamplesUseCase
	iCreateSampleUseCase application.ICreateSampleUseCase
	iUpdateSampleUseCase application.IUpdateSampleUseCase
	iDeleteSampleUseCase application.IDeleteSampleUseCase
}

// newApplications applicationsのコンストラクタ
func newApplications(
	iListSampleUseCase application.IListSamplesUseCase,
	iCreateSampleUseCase application.ICreateSampleUseCase,
	iUpdateSampleUseCase application.IUpdateSampleUseCase,
	iDeleteSampleUseCase application.IDeleteSampleUseCase,
) (*applications, error) {
	return &applications{
		iListSamplesUseCase:  iListSampleUseCase,
		iCreateSampleUseCase: iCreateSampleUseCase,
		iUpdateSampleUseCase: iUpdateSampleUseCase,
		iDeleteSampleUseCase: iDeleteSampleUseCase,
	}, nil
}

// createUsesCases applicationsのファクトリ
func createApplications(
	infras *infrastructures,
) (*applications, error) {
	// ユースケースを実行する構造体のインスタンス生成
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
	return newApplications(
		listSamplesUseCase,
		creatSampleUseCase,
		updateSampleUseCase,
		deleteSampleUseCase,
	)
}
