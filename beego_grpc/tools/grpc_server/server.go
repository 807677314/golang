package grpc_server

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServerControllers struct {
	beego.Controller
}

const (
	port = ":80"
)

type UserServer struct {
}

func (u *UserServer) GetUserById(c context.Context, user *UserRequest) (*UserIDResponse, error) {

	userID := user.UserId

	if userID == 0 {
		return nil, nil
	}

	return &UserIDResponse{UserId: userID}, nil

}

func (u *UserServer) GetUserName(c context.Context, user *UserRequest) (*UserNameResponse, error) {
	userName := user.UserName

	if userName == "" {
		return nil, nil
	}

	return &UserNameResponse{UserName: userName}, nil
}


// @router /grpc_server/ [GET,POST]
func (this *GrpcServerControllers) GrpcServer() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterUserServerServer(s, &UserServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
