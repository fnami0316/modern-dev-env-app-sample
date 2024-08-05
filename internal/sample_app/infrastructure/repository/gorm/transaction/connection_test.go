package transaction

import (
	"reflect"
	"testing"

	usecase "modern-dev-env-app-sample/internal/sample_app/usecase/repository/transaction"

	"gorm.io/gorm"
)

func TestNewGORMConnection(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *GORMConnection
	}{
		{
			name: "インスタンスを生成できる",
			args: args{
				db: &gorm.DB{},
			},
			want: &GORMConnection{
				db: &gorm.DB{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGORMConnection(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGORMConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORMConnection_Begin(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    usecase.ITransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORMConnection{
				db: tt.fields.db,
			}
			got, err := g.Begin()
			if (err != nil) != tt.wantErr {
				t.Errorf("Begin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Begin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORMConnection_Transaction(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		f func(iTx usecase.ITransaction) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORMConnection{
				db: tt.fields.db,
			}
			if err := g.Transaction(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("Transaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORMConnection_Con(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "コネクションの実体を返す",
			fields: fields{
				db: &gorm.DB{},
			},
			want: &gorm.DB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORMConnection{
				db: tt.fields.db,
			}
			if got := g.Con(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Con() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCon(t *testing.T) {
	type args struct {
		iCon usecase.IConnection
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "[OK]コネクションの実体を*gorm.DB型に変換した状態で取得",
			args: args{
				iCon: &GORMConnection{
					db: &gorm.DB{},
				},
			},
			want:    &gorm.DB{},
			wantErr: false,
		},
		{
			name: "[OK]iConがnilの場合nilを返す",
			args: args{
				iCon: nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Con(tt.args.iCon)
			if (err != nil) != tt.wantErr {
				t.Errorf("Con() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Con() got = %v, want %v", got, tt.want)
			}
		})
	}
}
