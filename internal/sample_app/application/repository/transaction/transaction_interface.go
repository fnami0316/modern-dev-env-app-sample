package transaction

type ITransaction interface {
	RollBack() error
	Commit() error
	Tx() interface{}
}

// Transaction Use Sample with interface
// iConnection, iTransaction, iRepository の実装は使うデータベース操作ライブラリ、ORMによる実装のセットを使うことを前提としている。
// つまり、iConnectionでGORMの実装を注入するなら、iTransaction, iRepositoryの実装にもGORMのものを注入する必要がある。
// ユースケース層にはiConnectionを満たすDBコネクションを注入すれば芋づる式にうまくいくように設計したつもり。
//
// 実装例1.
// {
//	iCon := NewGORMConnection(db)
//    iSampleRepo := NewSampleGORMRepository(db)
//
//	// トランザクション開始
//	iTx, err := iCon.Begin()
//	if err != nil {
//		return fmt.Errorf("failed to Begin(): %w", err)
//	}
//
//	// トランザクション内での処理
//	// 検索
//	sample, err := iSampleRepo.FindByID(id, itx)
//	if err != nil {
//		// トランザクションをロールバック
//		if err := iTx.RollBack(); err != nil {
//			return fmt.Errorf("failed to RollBack(): %w", err)
//		}
//		return fmt.Errorf("failed to FindByID(): %w", err)
//	}
//
//	// 保存
//	err := iSampleRepo.Save(sample, itx)
//	if err != nil {
//		// トランザクションをロールバック
//		if err := iTx.RollBack(); err != nil {
//			return fmt.Errorf("failed to RollBack(): %w", err)
//		}
//		return fmt.Errorf("failed to Save(): %w", err)
//	}
//
//	// トランザクションをコミット
//	if err := iTx.Commit(); err != nil {
//		return fmt.Errorf("failed to Commit(): %w", err)
//	}
// }

// 実装例2.
// {
//  // ユースケースのインスタンス生成時に外部から注入
//	iCon := NewGORMConnection(db)
//  iSampleRepo := NewSampleGORMRepository(db)
//
//	// トランザクション開始
//  // func(iTx) 内に記述した処理でエラーが発生した場合は、自動でロールバックしてくれる。
//  // エラーが発生しなかった場合は、自動でコミットしてくれる。
//  // 各IXxxRepositoryのメソッドを利用する際にはちゃんとiTxを渡さないとダメ。
//  // 逆にiTxがnilでない場合、各IXxxRepositoryはTransaction内でのデータ操作をする実装でなければならない
//	iCon.Transaction(func(iTx ITransaction) error {
//      // 検索
//      if sample, err := iSampleRepo.FindByID(id, iTx); err != nil {
//		  return fmt.Errorf("failed to FindByID(): %w", err)
//      }
//      // 保存
//	    if err := iSampleRepo.Save(sample, iTx); err != nil {
//		    return fmt.Errorf("failed to Save(): %w", err)
//	    }
//      return nil
//	}
// }
