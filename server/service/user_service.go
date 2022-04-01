package service

import (
	"context"
	"fmt"
	"log"

	userpc "github.com/boring-darwin/totalitycorp-demo/proto"
	"github.com/boring-darwin/totalitycorp-demo/server/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	userpc.UserServiceServer
}

func (*Server) GetUser(ctx context.Context, req *userpc.UserRequest) (*userpc.UserResponse, error) {

	log.Printf("Request for UserId: %d\n", req.Id)
	u, err := repository.FindUser(req.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find User with specified ID: %d", req.Id),
		)
	}

	return &userpc.UserResponse{
		Id:      u.Id,
		Fname:   u.Fname,
		City:    u.City,
		Phone:   u.Phone,
		Height:  u.Height,
		Married: u.Married,
	}, nil

	// return findUser(req.Id), nil
}

func (*Server) GetListOfUser(ctx context.Context, req *userpc.UserListRequest) (*userpc.UserListResponse, error) {

	var response userpc.UserListResponse

	for _, i := range req.Ids {
		u, err := repository.FindUser(i)

		if err == nil {
			r := &userpc.UserResponse{
				Id:      u.Id,
				Fname:   u.Fname,
				City:    u.City,
				Phone:   u.Phone,
				Height:  u.Height,
				Married: u.Married,
			}

			response.UserResponseList = append(response.UserResponseList, r)
		}

	}

	if len(response.UserResponseList) == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find User with specified IDs",
		)
	}

	return &response, nil

}
