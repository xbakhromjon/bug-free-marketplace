package client

import (
	"context"
	"fmt"
	"golang-project-template/internal/users/ports/grpc/proto/pb"
	"log"

	"google.golang.org/grpc"
)

func RunGRPCClient() {
	conn, err := grpc.Dial("localhost:5005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	//LoginUser
	phone_number := "998931234567"
	password := "Adm8n0ass"
	loginReq := &pb.LoginRequest{
		PhoneNumber: phone_number,
		Password:    password,
	}
	phoneNumberReq := &pb.PhoneNumberRequest{
		PhoneNumber: phone_number,
	}

	res, err := client.LoginUser(context.Background(), loginReq)
	if err != nil {
		log.Fatalf("error while calling LoginUser: %v", err)
	}
	fmt.Println("login sucess: ", res.Success)
	//GetByPhoneNumberUser
	user, err := client.GetUserByPhoneNumber(context.Background(), phoneNumberReq)
	if err != nil {
		log.Fatalf("error while calling GetByPhoneNumber: %v", err)
	}
	fmt.Printf("connection is successful...\n\n")
	fmt.Printf("Login success: %v\n\n", res.Success)
	fmt.Printf("User Name: %v", user.Name)
	fmt.Printf("User Phone Number : %v", user.PhoneNumber)

}
