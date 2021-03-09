package grpc

import (
	"context"

	transportPhonebook "github.com/sapawarga/proto-file/phonebook"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	phonebookGetList   kitgrpc.Handler
	phonebookGetDetail kitgrpc.Handler
}

func (g *grpcServer) GetList(ctx context.Context, req *transportPhonebook.GetListRequest) (*transportPhonebook.GetListResponse, error) {
	_, resp, err := g.phonebookGetList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportPhonebook.GetListResponse), nil
}

func (g *grpcServer) GetDetail(ctx context.Context, req *transportPhonebook.GetDetailRequest) (*transportPhonebook.GetDetailResponse, error) {
	_, resp, err := g.phonebookGetDetail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportPhonebook.GetDetailResponse), nil
}
