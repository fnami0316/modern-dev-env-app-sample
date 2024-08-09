package sample

import (
	"reflect"
	"testing"

	application "modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

func TestNewSampleServiceServer(t *testing.T) {
	type args struct {
		iListSamplesUseCase  application.IListSamplesUseCase
		iCreateSampleUseCase application.ICreateSampleUseCase
		iUpdateSampleUseCase application.IUpdateSampleUseCase
		iDeleteSampleUseCase application.IDeleteSampleUseCase
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
				iListSamplesUseCase:  &application.ListSamplesUseCase{},
				iCreateSampleUseCase: &application.CreateSampleUseCase{},
				iUpdateSampleUseCase: &application.UpdateSampleUseCase{},
				iDeleteSampleUseCase: &application.DeleteSampleUseCase{},
			},
			want: &SampleServiceServer{
				iListSamplesUseCase:  &application.ListSamplesUseCase{},
				iCreateSampleUseCase: &application.CreateSampleUseCase{},
				iUpdateSampleUseCase: &application.UpdateSampleUseCase{},
				iDeleteSampleUseCase: &application.DeleteSampleUseCase{},
			},
			wantErr: false,
		},
		{
			name: "[NG]バリデーションエラー",
			args: args{
				iListSamplesUseCase:  nil, // エラー
				iCreateSampleUseCase: &application.CreateSampleUseCase{},
				iUpdateSampleUseCase: &application.UpdateSampleUseCase{},
				iDeleteSampleUseCase: &application.DeleteSampleUseCase{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSampleServiceServer(
				tt.args.iListSamplesUseCase,
				tt.args.iCreateSampleUseCase,
				tt.args.iUpdateSampleUseCase,
				tt.args.iDeleteSampleUseCase,
			)
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
		iListSamplesUseCase              application.IListSamplesUseCase
		iCreateSampleUseCase             application.ICreateSampleUseCase
		iUpdateSampleUseCase             application.IUpdateSampleUseCase
		iDeleteSampleUseCase             application.IDeleteSampleUseCase
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
				iListSamplesUseCase:  &application.ListSamplesUseCase{},
				iCreateSampleUseCase: &application.CreateSampleUseCase{},
				iUpdateSampleUseCase: &application.UpdateSampleUseCase{},
				iDeleteSampleUseCase: &application.DeleteSampleUseCase{},
			},
			wantErr: false,
		},
		{
			name: "[NG]iListSamplesUseCaseがnilである場合エラー",
			fields: fields{
				iListSamplesUseCase:  nil,
				iCreateSampleUseCase: &application.CreateSampleUseCase{},
				iUpdateSampleUseCase: &application.UpdateSampleUseCase{},
				iDeleteSampleUseCase: &application.DeleteSampleUseCase{},
			},
			wantErr: true,
		},
		{
			name: "[NG]iCreateSampleUseCaseがnilである場合エラー",
			fields: fields{
				iListSamplesUseCase:  &application.ListSamplesUseCase{},
				iCreateSampleUseCase: nil,
				iUpdateSampleUseCase: &application.UpdateSampleUseCase{},
				iDeleteSampleUseCase: &application.DeleteSampleUseCase{},
			},
			wantErr: true,
		},
		{
			name: "[NG]iUpdateSampleUseCaseがnilである場合エラー",
			fields: fields{
				iListSamplesUseCase:  &application.ListSamplesUseCase{},
				iCreateSampleUseCase: &application.CreateSampleUseCase{},
				iUpdateSampleUseCase: nil,
				iDeleteSampleUseCase: &application.DeleteSampleUseCase{},
			},
			wantErr: true,
		},
		{
			name: "[NG]iDeleteSampleUseCaseがnilである場合エラー",
			fields: fields{
				iListSamplesUseCase:  &application.ListSamplesUseCase{},
				iCreateSampleUseCase: &application.CreateSampleUseCase{},
				iUpdateSampleUseCase: &application.UpdateSampleUseCase{},
				iDeleteSampleUseCase: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleServiceServer{
				iListSamplesUseCase:              tt.fields.iListSamplesUseCase,
				iCreateSampleUseCase:             tt.fields.iCreateSampleUseCase,
				iUpdateSampleUseCase:             tt.fields.iUpdateSampleUseCase,
				iDeleteSampleUseCase:             tt.fields.iDeleteSampleUseCase,
				UnimplementedSampleServiceServer: tt.fields.UnimplementedSampleServiceServer,
			}
			if err := s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
