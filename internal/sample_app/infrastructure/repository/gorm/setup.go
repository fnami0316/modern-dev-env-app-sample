package gorm

import (
	"fmt"

	spannergorm "github.com/googleapis/go-gorm-spanner"
	"gorm.io/gorm"
)

// Setup DBはSpanner、データ操作はORM GORMという依存を考慮したリポジトリを利用するためのセットアップ
// TODO: 各リポジトリのインスタンス生成までここでやっちゃってもいいかも？
func Setup(gcpProjectID, spannerInstanceID, spannerDatabaseID string) (*gorm.DB, error) {
	// TODO: Spannerではオートマイグレーションを実行してもインターリーブ部分は反映ができないことに注意
	db, err := gorm.Open(spannergorm.New(spannergorm.Config{
		DriverName: "spanner",
		DSN:        fmt.Sprintf("projects/%s/instances/%s/databases/%s", gcpProjectID, spannerInstanceID, spannerDatabaseID),
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to gorm.Open(): %w", err)
	}
	if err := db.AutoMigrate(autoMigrateGORMs...); err != nil {
		return nil, fmt.Errorf("failed to db.AutoMigrate(): %w", err)
	}
	return db, nil
}

// autoMigrateTargets GORMに自動でマイグレーションを任せるGORMを列挙
// GORM構造体(= テーブル)の定義が増えるたびに追加する必要あり
var autoMigrateGORMs = []interface{}{
	SampleGORM{},
}
