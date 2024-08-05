package transaction

import (
	"reflect"
	"testing"

	"modern-dev-env-app-sample/internal/sample_app/usecase/repository/transaction"

	"gorm.io/gorm"
)

func TestNewGORMTransaction(t *testing.T) {
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *GORMTransaction
	}{
		{
			name: "インスタンスを生成できる",
			args: args{
				tx: &gorm.DB{},
			},
			want: &GORMTransaction{
				tx: &gorm.DB{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGORMTransaction(tt.args.tx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGORMTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORMTransaction_RollBack(t *testing.T) {
	type fields struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORMTransaction{
				tx: tt.fields.tx,
			}
			if err := g.RollBack(); (err != nil) != tt.wantErr {
				t.Errorf("RollBack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORMTransaction_Commit(t *testing.T) {
	type fields struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORMTransaction{
				tx: tt.fields.tx,
			}
			if err := g.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORMTransaction_Tx(t *testing.T) {
	type fields struct {
		tx *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "フィールドのtxをそのまま返す",
			fields: fields{
				tx: &gorm.DB{},
			},
			want: &gorm.DB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORMTransaction{
				tx: tt.fields.tx,
			}
			if got := g.Tx(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConWithTx(t *testing.T) {
	type args struct {
		con *gorm.DB
		iTx transaction.ITransaction
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "iTxがnilの場合,conが返される",
			args: args{
				con: &gorm.DB{
					RowsAffected: 1,
				},
				iTx: nil,
			},
			want: &gorm.DB{
				RowsAffected: 1,
			},
			wantErr: false,
		},
		{
			name: "iTxがnilでない場合,iTx内部の*gorm.DBが返される",
			args: args{
				con: &gorm.DB{
					RowsAffected: 1,
				},
				iTx: &GORMTransaction{
					tx: &gorm.DB{},
				},
			},
			want:    &gorm.DB{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConWithTx(tt.args.con, tt.args.iTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConWithTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConWithTx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tx(t *testing.T) {
	type args struct {
		iTx transaction.ITransaction
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "*gorm.DB型に変換したTxを取得できる",
			args: args{
				iTx: &GORMTransaction{
					tx: &gorm.DB{},
				},
			},
			want: &gorm.DB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tx(tt.args.iTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("tx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
