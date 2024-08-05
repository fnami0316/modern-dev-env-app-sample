package transaction

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/usecase/repository/transaction"

	"gorm.io/gorm"
)

var _ usecase.ITransaction = &GORMTransaction{}

type GORMTransaction struct {
	tx *gorm.DB
}

// NewGORMTransaction GORMTransactionのコンストラクタ
func NewGORMTransaction(tx *gorm.DB) *GORMTransaction {
	return &GORMTransaction{
		tx: tx,
	}
}

// RollBack トランザクションをロールバック
func (g *GORMTransaction) RollBack() error {
	tx := g.tx.Rollback()
	if tx.Error != nil {
		return fmt.Errorf("failed to RollBack(): %w", tx.Error)
	}
	return nil
}

// Commit トランザクションをコミット
func (g *GORMTransaction) Commit() error {
	tx := g.tx.Commit()
	if tx.Error != nil {
		return fmt.Errorf("failed to Commit(): %w", tx.Error)
	}
	return nil
}

// Tx トランザクションの実体を返す
func (g *GORMTransaction) Tx() interface{} {
	return g.tx
}

// ConWithTx トランザクションを考慮した*gorm.DBを取得
// GORMを使ったリポジトリの各データ操作系メソッドを実行する際は、この結果をレシーバとして使うこと
func ConWithTx(con *gorm.DB, iTx usecase.ITransaction) (*gorm.DB, error) {
	// トランザクションオブジェクトの実体を取得
	conWithTx, err := tx(iTx)
	if err != nil {
		return nil, fmt.Errorf("failed to tx(): %w", err)
	}

	// トランザクションオブエジェクトの実体が取得できないならばトランザクション外の処理に
	if conWithTx == nil {
		conWithTx = con
	}
	return conWithTx, nil
}

// tx *gorm.DB型に変換したTxを取得する
func tx(iTx usecase.ITransaction) (*gorm.DB, error) {
	if iTx == nil {
		return nil, nil
	}
	tx, ok := iTx.Tx().(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to assert iTx to *gorm.DB")
	}
	return tx, nil
}
