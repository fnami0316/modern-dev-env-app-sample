package value

import (
	"reflect"
	"testing"
)

func TestNewSampleID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    SampleID
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				id: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "[NG]validate()でエラー",
			args: args{
				id: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSampleID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSampleID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewSampleID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleID_validate(t *testing.T) {
	tests := []struct {
		name    string
		s       SampleID
		wantErr bool
	}{
		{
			name:    "[NG]SampleID <= 0の場合",
			s:       SampleID(0),
			wantErr: true,
		},
		{
			name:    "[OK]SampleID > 0 の場合OK",
			s:       SampleID(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSampleID_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		s    SampleID
		want int64
	}{
		{
			name: "[OK]SampleIDをint64に変換できる",
			s:    1,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToInt64(); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleIDs_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		s    SampleIDs
		want []int64
	}{
		{
			name: "[OK]SampleIDsをint64のスライスに変換できる",
			s:    SampleIDs{1, 2, 3},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToInt64(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
