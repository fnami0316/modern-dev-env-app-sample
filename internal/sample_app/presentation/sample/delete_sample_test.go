package sample

import (
	"context"
	"errors"
	"reflect"
	"testing"

	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	"modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"

	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSampleServiceServer_DeleteSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := sample.NewMockDeleteSampleUseCase(ctrl)

	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		in0 context.Context
		req *pb.DeleteSampleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteSampleResponse
		wantErr bool
	}{
		{
			name: "[OK]ユースケースを実行できる",
			fields: fields{
				iDeleteSampleUseCase: func() sample.IDeleteSampleUseCase {
					usecase.EXPECT().Run(gomock.Any()).DoAndReturn(
						func(req *application.DeleteSampleRequest) (*application2.DeleteSampleResponse, error) {
							return &application2.DeleteSampleResponse{}, nil
						},
					)
					return usecase
				}(),
			},
			args: args{
				req: &pb.DeleteSampleRequest{
					Id: "id1",
				},
			},
			want: &pb.DeleteSampleResponse{
				Empty: &emptypb.Empty{},
			},
			wantErr: false,
		},
		{
			name: "[OK]ユースケースの実行でエラー",
			fields: fields{
				iDeleteSampleUseCase: func() sample.IDeleteSampleUseCase {
					usecase.EXPECT().Run(gomock.Any()).DoAndReturn(
						func(req *application.DeleteSampleRequest) (*application2.DeleteSampleResponse, error) {
							return nil, errors.New("run error")
						},
					)
					return usecase
				}(),
			},
			args: args{
				req: &pb.DeleteSampleRequest{
					Id: "id1",
				},
			},
			want:    nil,
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
			got, err := s.DeleteSample(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSample() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToDeleteSampleRequestForUseCase(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		pbReq *pb.DeleteSampleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *application.DeleteSampleRequest
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				pbReq: &pb.DeleteSampleRequest{
					Id: "id1",
				},
			},
			want:    newDeleteSampleRequestForTest(t, "id1"),
			wantErr: false,
		},
		{
			name: "[NG]値オブジェクトSampleID生成時にエラー",
			args: args{
				pbReq: &pb.DeleteSampleRequest{
					Id: "", // エラー
				},
			},
			want:    nil,
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
			got, err := s.convertToDeleteSampleRequestForUseCase(tt.args.pbReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToDeleteSampleRequestForUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToDeleteSampleRequestForUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToDeleteSampleResponseForProtoc(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		in0 *application2.DeleteSampleResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteSampleResponse
		wantErr bool
	}{
		{
			name:   "[OK]インスタンスを生成できる",
			fields: fields{},
			args:   args{},
			want: &pb.DeleteSampleResponse{
				Empty: &emptypb.Empty{},
			},
			wantErr: false,
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
			got, err := s.convertToDeleteSampleResponseForProtoc(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToDeleteSampleResponseForProtoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToDeleteSampleResponseForProtoc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// newDeleteSampleRequestForTest DeleteSampleRequestを生成(エラー発生時はテスト失敗扱い)
func newDeleteSampleRequestForTest(t *testing.T, id value.SampleID) *application.DeleteSampleRequest {
	req, err := application.NewDeleteSampleRequest(id)
	if err != nil {
		t.Fatalf("failed to NewDeleteSampleRequest(): %v", err)
	}
	return req
}
