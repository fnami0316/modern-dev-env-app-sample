package gorm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository"
	usecase2 "modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	infrastructure "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gorm.io/gorm"
)

func TestCreateSampleRepository(t *testing.T) {
	db := &gorm.DB{}
	type args struct {
		iCon usecase2.IConnection
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.ISampleRepository
		wantErr bool
	}{
		{
			name: "[OK]正常系",
			args: args{
				iCon: infrastructure.NewGORMConnection(db),
			},
			want: usecase.ISampleRepository(&SampleRepository{
				con: db,
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSampleRepository(tt.args.iCon)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSampleRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSampleRepository() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newSampleRepository(t *testing.T) {
	type args struct {
		con *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    *SampleRepository
		wantErr bool
	}{
		{
			name: "[OK]正常系",
			args: args{
				con: &gorm.DB{},
			},
			want: &SampleRepository{
				con: &gorm.DB{},
			},
			wantErr: false,
		},
		{
			name: "[NG]conがnil",
			args: args{
				con: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newSampleRepository(tt.args.con)
			if (err != nil) != tt.wantErr {
				t.Errorf("newSampleRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSampleRepository() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleRepository_validate(t *testing.T) {
	type fields struct {
		con *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]正常系",
			fields: fields{
				con: &gorm.DB{},
			},
			wantErr: false,
		},
		{
			name: "[NG]conがnil",
			fields: fields{
				con: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleRepository{
				con: tt.fields.con,
			}
			if err := s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestSampleRepository_Save Saveのテスト
// テストの前提条件
// - Spannerエミュレータが起動状態であり、localhost:9010でアクセス可能であること
// - Spannerレミュレータ上にDB projects/local-project/instances/test-instance/databases/test-database が作成されていること
// TODO: Spannerエミュレータにそもそもアクセスできるかをチェックする工程がほしい
func TestSampleRepository_Save(t *testing.T) {
	// GORMConnectionのdbフィールドと、リポジトリのdbフィールドは同じものがセットされている必要がある。
	gormCon := createGORMConForTest(t)
	con := infrastructure.NewGORMConnection(gormCon)

	type fields struct {
		con *gorm.DB
	}
	type args struct {
		sampleEntity *entity.Sample
	}
	tests := []struct {
		name         string
		setupSamples []*SampleGORM
		fields       fields
		args         args
		wantSamples  []*entity.Sample
		wantErr      bool
	}{
		{
			name: "[OK]指定したエンティティに相当するレコードがSampleテーブルに存在しない場合",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "2", "name2"),
			},
			fields: fields{
				con: gormCon,
			},
			args: args{
				sampleEntity: newSampleEntityForTest(t, "1", "name1"),
			},
			wantSamples: []*entity.Sample{
				newSampleEntityForTest(t, "1", "name1"),
			},
			wantErr: false,
		},
		{
			name: "[OK]指定したエンティティに相当するレコードがSampleテーブルに存在する場合",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
			},
			fields: fields{
				con: gormCon,
			},
			args: args{
				sampleEntity: newSampleEntityForTest(t, "1", "nameX"),
			},
			wantSamples: []*entity.Sample{
				newSampleEntityForTest(t, "1", "nameX"),
				newSampleEntityForTest(t, "2", "name2"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllSamplesForTest(t)
			setupSamplesForTest(t, []*SampleGORM{})
			defer deleteAllSamplesForTest(t)
			s := &SampleRepository{
				con: tt.fields.con,
			}

			// 書き込み操作はトランザクション内で実行しないとエラーになる
			err := con.Transaction(func(iTx usecase2.ITransaction) error {
				if err := s.Save(tt.args.sampleEntity, iTx); err != nil {
					return fmt.Errorf("failed to Save(): %w", err)
				}
				return nil
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Save実行後に全てのSamplesを取得して、結果が期待通りであるかチェック
			gotAllSamples, err := s.FindAll(nil)
			if err != nil {
				t.Errorf("FindAll() error = %v", err)
			}
			compareSampleEntityList(t, gotAllSamples, []*entity.Sample{tt.args.sampleEntity})
		})
	}
}

// TestSampleRepository_FindByIDs FindByIDのテスト
// テストの前提条件
// - Spannerエミュレータが起動状態であり、localhost:9010でアクセス可能であること
// - Spannerレミュレータ上にDB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func TestSampleRepository_FindByIDs(t *testing.T) {
	type fields struct {
		con *gorm.DB
	}
	type args struct {
		ids []value.SampleID
		iTx usecase2.ITransaction
	}
	tests := []struct {
		name         string
		setupSamples []*SampleGORM
		fields       fields
		args         args
		wantSamples  []*entity.Sample
		wantErr      bool
	}{
		{
			name: "[OK]Sampleテーブルにレコードが存在かつ指定したIDのレコードが存在する場合、指定したIDのSampleエンティティ群を過不足なく取得できる",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
				newSampleGORMForTest(t, "3", "name3"),
			},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				ids: []value.SampleID{"1", "2"},
				iTx: nil,
			},
			wantSamples: []*entity.Sample{
				newSampleEntityForTest(t, "1", "name1"),
				newSampleEntityForTest(t, "2", "name2"),
			},
			wantErr: false,
		},
		{
			name: "[OK]Sampleテーブルにレコードが存在し、指定したIDのレコードがすべて存在しない場合、空スライスが返ってくる",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
				newSampleGORMForTest(t, "3", "name3"),
			},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				ids: []value.SampleID{"4", "5"},
				iTx: nil,
			},
			wantSamples: []*entity.Sample{},
			wantErr:     false,
		},
		{
			name: "[OK]Sampleテーブルにレコードが存在し、指定したIDのレコードの一部が存在しない場合、存在した分の空スライスが返ってくる",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
				newSampleGORMForTest(t, "3", "name3"),
			},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				ids: []value.SampleID{"3", "4"},
				iTx: nil,
			},
			wantSamples: []*entity.Sample{
				newSampleEntityForTest(t, "3", "name3"),
			},
			wantErr: false,
		},
		{
			name: "[OK]Sampleテーブルにレコードが存在し、指定したIDが存在するレコードのものかつ重複している場合、重複は1つにまとめられて結果が返ってくる",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
				newSampleGORMForTest(t, "3", "name3"),
			},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				ids: []value.SampleID{"2", "3", "3"}, // 3が重複
				iTx: nil,
			},
			wantSamples: []*entity.Sample{
				newSampleEntityForTest(t, "2", "name2"),
				newSampleEntityForTest(t, "3", "name3"), // id3で返ってくるのは1レコードのみ
			},
			wantErr: false,
		},
		{
			name:         "[OK]Sampleテーブルにレコードが存在しない場合、空スライスが返ってくる",
			setupSamples: []*SampleGORM{},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				ids: []value.SampleID{"1"},
				iTx: nil,
			},
			wantSamples: []*entity.Sample{},
			wantErr:     false,
		},
		{
			name: "[NG]idsのサイズがゼロの場合エラー",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
				newSampleGORMForTest(t, "3", "name3"),
			},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				ids: []value.SampleID{},
				iTx: nil,
			},
			wantSamples: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllSamplesForTest(t)
			setupSamplesForTest(t, tt.setupSamples)
			defer deleteAllSamplesForTest(t)

			s := &SampleRepository{
				con: tt.fields.con,
			}
			got, err := s.FindByIDs(tt.args.ids, tt.args.iTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			compareSampleEntityList(t, got, tt.wantSamples)
		})
	}
}

// TestSampleRepository_FindAll FindAllのテスト
// テストの前提条件
// - Spannerエミュレータが起動状態であり、localhost:9010でアクセス可能であること
// - Spannerレミュレータ上にDB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func TestSampleRepository_FindAll(t *testing.T) {
	type fields struct {
		con *gorm.DB
	}
	type args struct {
		iTx usecase2.ITransaction
	}
	tests := []struct {
		name         string
		setupSamples []*SampleGORM
		fields       fields
		args         args
		wantSamples  []*entity.Sample
		wantErr      bool
	}{
		{
			name: "[OK]Sampleテーブルにレコードが存在する場合、全てのSampleエンティティ群を過不足なく取得できる",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
				newSampleGORMForTest(t, "3", "name3"),
			},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				iTx: nil,
			},
			wantSamples: []*entity.Sample{
				newSampleEntityForTest(t, "1", "name1"),
				newSampleEntityForTest(t, "2", "name2"),
				newSampleEntityForTest(t, "3", "name3"),
			},
			wantErr: false,
		},
		{
			name:         "[OK]Sampleテーブルにレコードが存在しない場合、空スライスが返ってくる",
			setupSamples: []*SampleGORM{},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				iTx: nil,
			},
			wantSamples: []*entity.Sample{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllSamplesForTest(t)
			setupSamplesForTest(t, tt.setupSamples)
			defer deleteAllSamplesForTest(t)
			s := &SampleRepository{
				con: tt.fields.con,
			}
			got, err := s.FindAll(tt.args.iTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			compareSampleEntityList(t, got, tt.wantSamples)
		})
	}
}

// TestSampleRepository_Delete Deleteのテスト
// テストの前提条件
// - Spannerエミュレータが起動状態であり、localhost:9010でアクセス可能であること
// - Spannerレミュレータ上にDB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func TestSampleRepository_Delete(t *testing.T) {
	type fields struct {
		con *gorm.DB
	}
	type args struct {
		sample *entity.Sample
		iTx    usecase2.ITransaction
	}
	tests := []struct {
		name         string
		setupSamples []*SampleGORM
		fields       fields
		args         args
		wantSamples  []*entity.Sample
		wantErr      bool
	}{
		{
			name: "[OK]Sampleテーブルにレコードが存在し、指定したIDのレコードが存在する場合、指定したIDのSampleエンティティを削除できる",
			setupSamples: []*SampleGORM{
				newSampleGORMForTest(t, "1", "name1"),
				newSampleGORMForTest(t, "2", "name2"),
			},
			fields: fields{
				con: createGORMConForTest(t),
			},
			args: args{
				sample: newSampleEntityForTest(t, "1", "name1"),
			},
			wantSamples: []*entity.Sample{
				newSampleEntityForTest(t, "2", "name2"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllSamplesForTest(t)
			setupSamplesForTest(t, tt.setupSamples)
			defer deleteAllSamplesForTest(t)

			s := &SampleRepository{
				con: tt.fields.con,
			}
			if err := s.Delete(tt.args.sample, tt.args.iTx); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Delete実行後の全てのSamplesを取得して、結果が期待通りであるかチェック
			gotAllSamples, err := s.FindAll(tt.args.iTx)
			if err != nil {
				t.Errorf("FindAll() error = %v", err)
			}
			compareSampleEntityList(t, gotAllSamples, tt.wantSamples)
		})
	}
}

func TestSampleRepository_convEntityToGORM(t *testing.T) {
	type fields struct {
		con *gorm.DB
	}
	type args struct {
		sampleEntity *entity.Sample
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SampleGORM
		wantErr bool
	}{
		{
			name:   "[OK]SampleエンティティをSampleGORMへ変換できる",
			fields: fields{},
			args: args{
				sampleEntity: newSampleEntityForTest(t, "1", "name"),
			},
			want: &SampleGORM{
				ID:   "1",
				Name: "name",
			},
			wantErr: false,
		},
		{
			name:   "[NG]Sampleエンティティがnilの場合エラー",
			fields: fields{},
			args: args{
				sampleEntity: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleRepository{
				con: tt.fields.con,
			}
			got, err := s.convEntityToGORM(tt.args.sampleEntity)
			if (err != nil) != tt.wantErr {
				t.Errorf("convEntityToGORM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convEntityToGORM() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleRepository_convGORMToEntity(t *testing.T) {
	type fields struct {
		con *gorm.DB
	}
	type args struct {
		sampleGORM *SampleGORM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Sample
		wantErr bool
	}{
		{
			name:   "[OK]正常系",
			fields: fields{},
			args: args{
				sampleGORM: newSampleGORMForTest(t, "1", "name"),
			},
			want:    newSampleEntityForTest(t, "1", "name"),
			wantErr: false,
		},
		{
			name:   "[NG]sampleGORMがnilの場合エラー",
			fields: fields{},
			args: args{
				sampleGORM: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleRepository{
				con: tt.fields.con,
			}
			got, err := s.convGORMToEntity(tt.args.sampleGORM)
			if (err != nil) != tt.wantErr {
				t.Errorf("convGORMToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convGORMToEntity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleRepository_convGORMListToEntityList(t *testing.T) {
	type fields struct {
		con *gorm.DB
	}
	type args struct {
		sampleGORMs []*SampleGORM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entity.Sample
		wantErr bool
	}{
		{
			name:   "[OK]正常系",
			fields: fields{},
			args: args{
				sampleGORMs: []*SampleGORM{
					newSampleGORMForTest(t, "1", "name1"),
					newSampleGORMForTest(t, "2", "name2"),
				},
			},
			want: []*entity.Sample{
				newSampleEntityForTest(t, "1", "name1"),
				newSampleEntityForTest(t, "2", "name2"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleRepository{
				con: tt.fields.con,
			}
			got, err := s.convGORMListToEntityList(tt.args.sampleGORMs)
			if (err != nil) != tt.wantErr {
				t.Errorf("convGORMListToEntityList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convGORMListToEntityList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSampleGORM(t *testing.T) {
	type args struct {
		id   string
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *SampleGORM
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				id:   "1",
				name: "name",
			},
			want: &SampleGORM{
				ID:        "1",
				Name:      "name",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSampleGORM(tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSampleGORM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSampleGORM() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleGORM_validate(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "[OK]正常系",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleGORM{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			if err := s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// setupSamplesForTest SpannerエミュレータのSampleテーブルへレコードをセットアップ
// 前提条件
// - Spannerエミュレータが起動状態であり、localhost:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func setupSamplesForTest(t *testing.T, sampleGORMs []*SampleGORM) {
	if len(sampleGORMs) == 0 {
		return
	}
	con := createGORMConForTest(t)
	err := con.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(sampleGORMs).Error; err != nil {
			return fmt.Errorf("failed to Create(): %w", err)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to Transaction(): %v", err)
	}
}

// deleteSampleRecordAllWithGORM SpannerエミュレータのSampleテーブルの全レコードを削除
// 前提条件
// - Spannerエミュレータが起動状態であり、localhost:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func deleteAllSamplesForTest(t *testing.T) {
	con := createGORMConForTest(t)
	// Unscoped(): 指定することで物理削除になる
	// Where(): GORMの縛りで、Delete()実行前のWhere()に何かしらの条件を指定しないとエラーになる。全レコード削除したいため、意味のない条件をセットして実行している
	con.Unscoped().Where("1 = 1").Delete(&SampleGORM{})
}

// createGORMConForEmulator テスト用のGORMコネクションを生成
// 利用するための前提条件
// - Spannerエミュレータが起動状態であり、localhost:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func createGORMConForTest(t *testing.T) *gorm.DB {
	err := os.Setenv("SPANNER_EMULATOR_HOST", "localhost:9010")
	if err != nil {
		t.Fatalf("failed to Setenv(): %v", err)
	}
	con, err := Setup(
		"local-project",
		"test-instance",
		"test-database",
	)
	if err != nil {
		t.Fatalf("failed to Setup(): %v", err)
	}
	return con
}

// newSampleEntityForTest Sampleエンティティを生成(エラーをテスト失敗として扱う)
func newSampleEntityForTest(t *testing.T, id value.SampleID, name value.SampleName) *entity.Sample {
	sampleEntity, err := entity.NewSample(id, name)
	if err != nil {
		t.Fatalf("failed to NewSample(): %v", err)
	}
	return sampleEntity
}

// newSampleGORMForTest SampleGORMを生成(エラーはテスト失敗として扱う)
func newSampleGORMForTest(t *testing.T, id string, name string) *SampleGORM {
	sampleGORM, err := NewSampleGORM(id, name)
	if err != nil {
		t.Fatalf("failed to NewSampleGORM(): %v", err)
	}
	return sampleGORM
}

// compareSampleEntityList サンプルエンティティリストの比較
func compareSampleEntityList(t *testing.T, got, want []*entity.Sample) {
	opts := []cmp.Option{
		cmp.AllowUnexported(entity.Sample{}),
		cmpopts.SortSlices(func(a, b *entity.Sample) bool {
			return a.ID() < b.ID()
		}),
	}
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Errorf("(-got +want)\n%s", diff)
	}
}
