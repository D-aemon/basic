package handler

import (
	"basic/api/proto"
	"context"
)

type StreamService struct {
	DB
}

func (s StreamService) UpsertCamera(ctx context.Context, request *proto.UpsertCameraRequest) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s StreamService) DeleteCamera(ctx context.Context, request *proto.DeleteCameraRequest) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s StreamService) PreviewCamera(ctx context.Context, request *proto.PreviewCameraRequest) (*proto.PreviewCameraResponse, error) {
	//TODO implement me
	panic("implement me")
}

type DB interface {
}
