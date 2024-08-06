package value

import "testing"

func TestNewSampleName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    SampleName
		wantErr bool
	}{
		{
			name: "[OK]バリデーションでエラーがない場合、インスタンスを生成できる",
			args: args{
				name: "name",
			},
			want:    SampleName("name"),
			wantErr: false,
		},
		{
			name: "[NG]バリデーションでエラーがある場合、エラーを返す",
			args: args{
				name: "",
			},
			want:    SampleName(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSampleName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSampleName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewSampleName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleName_validate(t *testing.T) {
	tests := []struct {
		name    string
		s       SampleName
		wantErr bool
	}{
		{
			name:    "[OK]サイズが1以上の場合エラーが返らない",
			s:       SampleName(" "),
			wantErr: false,
		},
		{
			name:    "[NG]サイズが0の場合エラーが返る",
			s:       SampleName(""),
			wantErr: true,
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

func TestSampleName_ToString(t *testing.T) {
	tests := []struct {
		name string
		s    SampleName
		want string
	}{
		{
			name: "[OK]SampleNameを文字列に変換できる",
			s:    SampleName("name"),
			want: "name",
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
