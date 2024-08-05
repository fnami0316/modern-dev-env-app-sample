package main

import (
	"fmt"
	"log"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository"
	usecase2 "modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	infra "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm"
	infra2 "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"
)

// infrastructures インフラ層の全インスタンスをまとめた構造体
type infrastructures struct {
	iCon        usecase2.IConnection
	iSampleRepo usecase.ISampleRepository
}

// newInfrastructures コンストラクタ
func newInfrastructures(
	iCon usecase2.IConnection,
	iSampleRepo usecase.ISampleRepository,
) *infrastructures {
	return &infrastructures{
		iCon:        iCon,
		iSampleRepo: iSampleRepo,
	}
}

// createInfrastructures インフラ層の全インスタンスをDB:Spanner,ORM:GORMに依存したリポジトリ,トランザクション制御の実装を使う前提で生成するファクトリ
func createInfrastructuresWithGORMSpanner(
	gcpProjectID string,
	spannerInstanceID string,
	spannerDatabaseID string,
) (*infrastructures, error) {
	con, err := infra.Setup(gcpProjectID, spannerInstanceID, spannerDatabaseID)
	if err != nil {
		log.Fatalf("failed to Setup(): %v", err)
	}
	iCon := infra2.NewGORMConnection(con)
	iSampleRepo, err := infra.CreateSampleRepository(iCon)
	if err != nil {
		return nil, fmt.Errorf("failed to CreateSampleRepository(): %w", err)
	}
	return newInfrastructures(iCon, iSampleRepo), nil
}
