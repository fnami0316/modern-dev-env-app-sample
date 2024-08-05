package sample

import (
	"reflect"
	"testing"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

func TestNewSampleServiceServer(t *testing.T) {
	type args struct {
		iListSamplesUseCase usecase.IListSamplesUseCase
	}
	tests := []struct {
		name    string
		args    args
		want    *SampleServiceServer
		wantErr bool
	}{
		{
			name: "[OK]全てのチェックを通過",
			args: args{
				iListSamplesUseCase: &usecase.ListSamplesUseCase{},
			},
			want: &SampleServiceServer{
				iListSamplesUseCase: &usecase.ListSamplesUseCase{},
			},
			wantErr: false,
		},
		{
			name: "[NG]バリデーションエラー",
			args: args{
				iListSamplesUseCase: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSampleServiceServer(tt.args.iListSamplesUseCase)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSampleServiceServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSampleServiceServer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_validate(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              usecase.IListSamplesUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "[OK]全てのチェックを通過",
			fields: fields{
				iListSamplesUseCase: &usecase.ListSamplesUseCase{},
			},
			wantErr: false,
		},
		{
			name: "[NG]iListSamplesUseCaseがnilである場合エラー",
			fields: fields{
				iListSamplesUseCase: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleServiceServer{
				iListSamplesUseCase:              tt.fields.iListSamplesUseCase,
				UnimplementedSampleServiceServer: tt.fields.UnimplementedSampleServiceServer,
			}
			if err := s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
