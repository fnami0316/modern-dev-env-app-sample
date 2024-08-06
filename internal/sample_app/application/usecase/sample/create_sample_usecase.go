package sample

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository"
	usecase2 "modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
)

var _ ICreateSampleUseCase = &CreateSampleUseCase{}

type CreateSampleUseCase struct {
	iCon        usecase2.IConnection
	iSampleRepo usecase.ISampleRepository
}

// NewCreateSampleUseCase CreateSampleUsecaseのコンストラクタ
func NewCreateSampleUseCase(iCon usecase2.IConnection, iSampleRepo usecase.ISampleRepository) (*CreateSampleUseCase, error) {
	useCase := &CreateSampleUseCase{
		iCon:        iCon,
		iSampleRepo: iSampleRepo,
	}
	if err := useCase.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return useCase, nil
}

// validate CreateSampleUsecaseのバリデーション
func (l *CreateSampleUseCase) validate() error {
	if l.iCon == nil {
		return errors.New("iCon is nil")
	}
	if l.iSampleRepo == nil {
		return errors.New("iSampleRepo is nil")
	}
	return nil
}

// Run ユースケース: サンプルデータを作成 を実行
func (l *CreateSampleUseCase) Run(req *application.CreateSampleRequest) (*application2.CreateSampleResponse, error) {
	createdSample, err := entity.CreateDefaultSample(req.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to CreateDefaultSample(): %w", err)
	}

	if err := l.iCon.Transaction(func(iTx usecase2.ITransaction) error {
		if err := l.iSampleRepo.Save(createdSample, iTx); err != nil {
			return fmt.Errorf("failed to Save(): %w", err)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to Transaction(): %w", err)
	}

	res, err := application2.NewCreateSampleResponse(createdSample)
	if err != nil {
		return nil, fmt.Errorf("failed to NewCreateSampleResponse(): %w", err)
	}
	return res, nil
}
