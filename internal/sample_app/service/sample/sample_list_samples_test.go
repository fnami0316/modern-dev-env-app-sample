package sample

import (
	"context"
	"reflect"
	"testing"

	pb "modern-dev-env-app-sample/internal/sample_app/pb/api/proto"
)

func TestSampleServiceServer_ListSamples(t *testing.T) {
	type fields struct {
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.ListSamplesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListSamplesResponse
		wantErr bool
	}{
		{
			name:   "固定値のレスポンスが返ってくること",
			fields: fields{},
			args:   args{},
			want: &pb.ListSamplesResponse{
				Samples: []*pb.Sample{
					{
						Id:   1,
						Name: "name1",
					},
					{
						Id:   2,
						Name: "name2",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleServiceServer{
				UnimplementedSampleServiceServer: tt.fields.UnimplementedSampleServiceServer,
			}
			got, err := s.ListSamples(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSamples() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSamples() got = %v, want %v", got, tt.want)
			}
		})
	}
}
