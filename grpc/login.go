package grpc

import (
	"golang.org/x/net/context"

	"github.com/charlesfan/go-api/pb"
	"github.com/charlesfan/go-api/service/rsi"
)

type login struct{}

func (l *login) EmailChecking(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	// Implement: Use backend service
	svc := rsi.LoginService
	e := &rsi.EmailLoginBody{
		Email:    in.Email,
		Password: in.Password,
	}

	result := "success"

	err := svc.EmailChecking(e)
	if err != nil {
		result = "failed"
	}

	// Build with Protobuf and return
	return &pb.LoginReply{Result: result}, err
}
