package main

import (
	"context"
	"fmt"
	"log"

	userpc "github.com/boring-darwin/totalitycorp-demo/proto"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("User Client Started")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:4041", opts)

	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}
	defer cc.Close()

	c := userpc.NewUserServiceClient(cc)

	res, err := c.GetUser(context.Background(), &userpc.UserRequest{Id: 4})

	if err != nil {
		log.Printf("Unexpected error : %v\n", err)
	} else {
		fmt.Printf("User Response : %v\n", res)
	}

	b := []int32{1, 2, 6}
	listRes, err := c.GetListOfUser(context.Background(), &userpc.UserListRequest{Ids: b})

	if err != nil {
		log.Printf("Unexpected error : %v\n", err)
	} else {
		fmt.Printf("User Response : %v\n", listRes)
	}

}
