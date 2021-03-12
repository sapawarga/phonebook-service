package grpc

import (
	"context"

	transportPhonebook "github.com/sapawarga/proto-file/phonebook"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	phonebookGetList   kitgrpc.Handler
	phonebookGetDetail kitgrpc.Handler
	phonebookAdd       kitgrpc.Handler
	phonebookUpdate    kitgrpc.Handler
	phonebookDelete    kitgrpc.Handler
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

func (g *grpcServer) AddPhonebook(ctx context.Context, req *transportPhonebook.AddPhonebookRequest) (*transportPhonebook.StatusResponse, error) {
	_, resp, err := g.phonebookAdd.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportPhonebook.StatusResponse), nil
}

func (g *grpcServer) UpdatePhonebook(ctx context.Context, req *transportPhonebook.UpdatePhonebookRequest) (*transportPhonebook.StatusResponse, error) {
	_, resp, err := g.phonebookUpdate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportPhonebook.StatusResponse), nil
}

func (g *grpcServer) DeletePhonebook(ctx context.Context, req *transportPhonebook.GetDetailRequest) (*transportPhonebook.StatusResponse, error) {
	_, resp, err := g.phonebookDelete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportPhonebook.StatusResponse), nil
}
