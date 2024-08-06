package value

import (
	"reflect"
	"slices"
	"testing"
)

func TestCreateRandomSampleID(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "[OK]ランダムなSampleIDを生成できる",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRandomSampleIDs := make([]SampleID, 10)
			for i := 0; i < 10; i++ {
				got, err := CreateRandomSampleID()
				if (err != nil) != tt.wantErr {
					t.Errorf("CreateRandomSampleID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				gotRandomSampleIDs[i] = got
			}

			slices.Sort(gotRandomSampleIDs)
			compactedIDs := slices.Compact(gotRandomSampleIDs)
			if len(compactedIDs) != 10 {
				t.Errorf("CreateRandomSampleID() is not Random got = %v,", gotRandomSampleIDs)
			}
		})
	}
}

func TestNewSampleID(t *testing.T) {
	type args struct {
		id string
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
				id: "1",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "[NG]validate()でエラー",
			args: args{
				id: "",
			},
			want:    "",
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
			name:    "[NG]サイズが0の場合",
			s:       "",
			wantErr: true,
		},
		{
			name:    "[OK]サイズが0より大きい場合OK",
			s:       "x",
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

func TestSampleID_ToString(t *testing.T) {
	tests := []struct {
		name string
		s    SampleID
		want string
	}{
		{
			name: "[OK]SampleIDをstringに変換できる",
			s:    "x",
			want: "x",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleIDs_ToString(t *testing.T) {
	tests := []struct {
		name string
		s    SampleIDs
		want []string
	}{
		{
			name: "[OK]SampleIDsをstringのスライスに変換できる",
			s:    SampleIDs{"x", "y", "z"},
			want: []string{"x", "y", "z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
