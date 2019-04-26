package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/neatflux/neatflux-backend/pkg/api/v1"
)

const (
	currentAPIVer = "v1"
)

type contentServer struct{}

// checkAPI checks if the API version requested by client is supported by server
func (s *contentServer) checkAPIversion(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 && currentAPIVer != api {
		return status.Errorf(codes.Unimplemented,
			"unsupported API version: service implements API version '%s', but asked for '%s'", currentAPIVer, api)
	}
	return nil
}

func (s *contentServer) GetVideoInfo(ctx context.Context, req *v1.VideoInfoRequest) (*v1.VideoInfoReply, error) {

	if err := s.checkAPIversion(req.ApiVersion); err != nil {
		return nil, err
	}

	vidInfo := v1.VideoInfo{
		Title: "Cute cat video",
	}

	return &v1.VideoInfoReply{
		VideoInfo: &vidInfo,
	}, nil
}
