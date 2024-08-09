package gorm

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

// テストの前提条件
// - Spannerエミュレータが起動状態であり、spanner-emulator:9010でアクセス可能であること
// - DB projects/local-project/instances/test-instance/databases/test-database が作成されていること
func TestSetup(t *testing.T) {
	err := os.Setenv("SPANNER_EMULATOR_HOST", "spanner-emulator:9010")
	if err != nil {
		t.Fatalf("failed to Setenv(): %v", err)
	}
	type args struct {
		gcpProjectID      string
		spannerInstanceID string
		spannerDatabaseID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantNil bool
		wantErr bool
	}{
		{
			name: "[OK]引数で指定した設定を反映したSpannerドライバのgorm.DBインスタンスを生成できる",
			args: args{
				gcpProjectID:      "local-project",
				spannerInstanceID: "test-instance",
				spannerDatabaseID: "test-database",
			},
			wantNil: false,
			wantErr: false,
		},
		{
			name: "[NG]DBが存在しない場合エラー",
			args: args{
				gcpProjectID:      "local-project",
				spannerInstanceID: "test", // 存在しないインスタンスID
				spannerDatabaseID: "test-database",
			},
			wantNil: true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Setup(tt.args.gcpProjectID, tt.args.spannerInstanceID, tt.args.spannerDatabaseID)
			if (got != nil) == tt.wantNil {
				t.Errorf("Setup() got = %v, wantNil %v", got, tt.wantNil)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
