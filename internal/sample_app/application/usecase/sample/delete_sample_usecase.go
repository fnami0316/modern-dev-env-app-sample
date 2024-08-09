package sample

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository"
	usecase2 "modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

var _ IDeleteSampleUseCase = &DeleteSampleUseCase{}

type DeleteSampleUseCase struct {
	iCon        usecase2.IConnection
	iSampleRepo usecase.ISampleRepository
}

// NewDeleteSampleUseCase DeleteSampleUsecaseのコンストラクタ
func NewDeleteSampleUseCase(iCon usecase2.IConnection, iSampleRepo usecase.ISampleRepository) (*DeleteSampleUseCase, error) {
	useCase := &DeleteSampleUseCase{
		iCon:        iCon,
		iSampleRepo: iSampleRepo,
	}
	if err := useCase.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return useCase, nil
}

// validate DeleteSampleUsecaseのバリデーション
func (l *DeleteSampleUseCase) validate() error {
	if l.iCon == nil {
		return errors.New("iCon is nil")
	}
	if l.iSampleRepo == nil {
		return errors.New("iSampleRepo is nil")
	}
	return nil
}

// Run ユースケース: サンプルデータを削除 を実行
func (l *DeleteSampleUseCase) Run(req *application.DeleteSampleRequest) (*application2.DeleteSampleResponse, error) {
	id := req.ID()

	if err := l.iCon.Transaction(func(iTx usecase2.ITransaction) error {
		samples, err := l.iSampleRepo.FindByIDs(value.SampleIDs{id}, iTx)
		if err != nil {
			return fmt.Errorf("failed to FindByIDs(): %w", err)
		}
		if len(samples) == 0 {
			return fmt.Errorf("failed to FindByIDs(): request id not exist in data store")
		}

		sampleEntity := samples[0]
		if err := l.iSampleRepo.Delete(sampleEntity, iTx); err != nil {
			return fmt.Errorf("failed to Delete(): %w", err)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to Transaction(): %w", err)
	}

	return &application2.DeleteSampleResponse{}, nil
}
