package transaction

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/usecase/repository/transaction"

	"gorm.io/gorm"
)

var _ usecase.IConnection = &GORMConnection{}

type GORMConnection struct {
	db *gorm.DB
}

// NewGORMConnection GORMConnectionのコンストラクタ
func NewGORMConnection(db *gorm.DB) *GORMConnection {
	return &GORMConnection{
		db: db,
	}
}

// Begin トランザクションを開始
func (g *GORMConnection) Begin() (usecase.ITransaction, error) {
	tx := g.db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to Begin(): %w", tx.Error)
	}

	iTx := NewGORMTransaction(tx)
	return iTx, nil
}

// Transaction
// 与えられたfuncがエラーを返す場合、このメソッド内で自動でロールバック。
// nilを返す場合はこのメソッド内でコミット
// func(iTx)内の処理には、更新系のリポジトリメソッドをiTxを渡して使うことを期待している
func (g *GORMConnection) Transaction(f func(iTx usecase.ITransaction) error) error {
	iTx, err := g.Begin()
	if err != nil {
		return fmt.Errorf("failed to Begin(): %w", err)
	}

	if err := f(iTx); err != nil {
		if err := iTx.RollBack(); err != nil {
			return fmt.Errorf("failed to f() and failed to RollBack(): %w", err)
		}
		return fmt.Errorf("failed to f() and execute RollBack(): %w", err)
	}

	if err := iTx.Commit(); err != nil {
		return fmt.Errorf("failed to Commit(): %w", err)
	}
	return nil
}

// Con コネクションの実体(*gorm.DB)を返す
func (g *GORMConnection) Con() interface{} {
	return g.db
}

// Con コネクションの実体を*gorm.DB型に変換した状態で取得
// GORM実装のリポジトリを生成する際には、この結果をレシーバとして使うこと
func Con(iCon usecase.IConnection) (*gorm.DB, error) {
	if iCon == nil {
		return nil, nil
	}
	con, ok := iCon.Con().(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to assert iCon.Con() to *gorm.DB")
	}
	return con, nil
}
