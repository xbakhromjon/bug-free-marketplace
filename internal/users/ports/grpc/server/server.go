package server

import (
	"context"
	"golang-project-template/internal/pkg/jwt"
	"golang-project-template/internal/users/app"
	"golang-project-template/internal/users/domain"
	"golang-project-template/internal/users/ports/grpc/proto/pb"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedUserServiceServer
	userUsecase app.UserUsecase
}

func NewUserGrpcServer(userUsecase app.UserUsecase) *server {
	return &server{
		userUsecase: userUsecase,
	}
}

func (s *server) RegisterMechantUser(ctx context.Context, req *pb.NewUser) (*pb.MerchantReply, error) {
	var merchantUser = domain.NewUser{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	id, err := s.userUsecase.RegisterMerchantUser(&merchantUser)
	if err != nil {
		log.Printf("Error registering merchant user: %v", err.Error())
		return &pb.MerchantReply{}, status.Error(codes.Internal, "Error registering merchant user")
	}

	res := &pb.MerchantReply{
		Id: int64(id),
	}

	return res, nil
}

func (s *server) RegisterCustomer(ctx context.Context, req *pb.NewUser) (*pb.CustomerReply, error) {
	var customerUser = domain.NewUser{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
	id, err := s.userUsecase.RegisterCustomer(&customerUser)
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return &pb.CustomerReply{}, status.Error(codes.Internal, "Error registering customer user")
	}

	res := &pb.CustomerReply{
		Id: int64(id),
	}

	return res, nil
}

func (s *server) RegisterAdmin(ctx context.Context, req *pb.NewUser) (*pb.AdminReply, error) {
	var adminUser = domain.NewUser{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	id, err := s.userUsecase.RegisterAdmin(&adminUser)
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return &pb.AdminReply{}, status.Error(codes.Internal, "Error registering admin user")
	}

	res := &pb.AdminReply{
		Id: int64(id),
	}

	return res, nil
}

func (s *server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {

	success, err := s.userUsecase.LoginUser(req.PhoneNumber, req.Password)
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return &pb.LoginReply{}, status.Error(codes.Internal, "Internal error: "+err.Error())
	}

	if success {
		token, err := jwt.CreateToken(req.PhoneNumber)
		if err != nil {
			log.Println("phone number not found: " + err.Error())
			return &pb.LoginReply{}, status.Error(codes.NotFound, "Not found: "+err.Error())
		}
		res := &pb.LoginReply{
			Success: true,
			Token:   token,
		}
		log.Println("Login success: ", res.Success)
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "User is not registered yet")
}

func (s *server) GetUserByPhoneNumber(ctx context.Context, req *pb.PhoneNumberRequest) (*pb.User, error) {
	phone_number := req.PhoneNumber
	log.Println("phone_number: ", phone_number)
	user, err := s.userUsecase.GetUserByPhoneNumber(phone_number)
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return &pb.User{}, status.Error(codes.Internal, "Internal error: "+err.Error())
	}
	createdAtTimestamp := timestamppb.New(user.GetCreatedAt())
	updatedAtTimestamp := timestamppb.New(user.GetUpdatedAt())
	res := &pb.User{
		Id:          int64(user.GetID()),
		PhoneNumber: user.GetPhoneNumber(),
		Name:        user.GetName(),
		Role:        user.GetRole(),
		CreatedAt:   createdAtTimestamp,
		UpdatedAt:   updatedAtTimestamp,
	}
	return res, nil
}

func (s *server) GetUserByID(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	id := int(req.Id)

	user, err := s.userUsecase.GetUserByID(id)
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return &pb.User{}, status.Error(codes.Internal, "Internal error: "+err.Error())
	}
	log.Println("Be careful, Asadbek is connecting to your server now....")
	createdAtTimestamp := timestamppb.New(user.GetCreatedAt())
	updatedAtTimestamp := timestamppb.New(user.GetUpdatedAt())
	time.Sleep(time.Second * 6)
	res := &pb.User{
		Id:          int64(user.GetID()),
		PhoneNumber: user.GetPhoneNumber(),
		Name:        user.GetName(),
		Role:        user.GetRole(),
		CreatedAt:   createdAtTimestamp,
		UpdatedAt:   updatedAtTimestamp,
	}
	return res, nil
}

func (s *server) UserExists(ctx context.Context, req *pb.UserID) (*pb.UserExistsReply, error) {
	exists, err := s.userUsecase.UserExists(int(req.Id))
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return &pb.UserExistsReply{}, status.Error(codes.Internal, "")
	}
	res := &pb.UserExistsReply{
		Exists: exists,
	}
	return res, nil
}
