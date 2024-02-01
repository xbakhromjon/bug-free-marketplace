package app

import (
	"fmt"
	"log"
	"os"

	"golang-project-template/cmd/app/servers"

	"golang-project-template/internal/common"
	"golang-project-template/internal/pkg/config"
	userRepo "golang-project-template/internal/users/adapters"
	userApp "golang-project-template/internal/users/app"
	userPb "golang-project-template/internal/users/ports/grpc/proto/pb"
	userGrpcServer "golang-project-template/internal/users/ports/grpc/server"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{Use: "run-grpc"}

// command to run grpc server
var userGrpcServerCmd = &cobra.Command{
	Use: "user-grpc-server",

	Run: func(cmd *cobra.Command, args []string) {

		db, err := common.ConnectToDb(os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DATABASE"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"))

		if err != nil {
			log.Fatalf("Failed to connect to database: %s", err)
		}

		config := config.NewConfig()
		rpcPort := fmt.Sprintf(":%s", config.RpcPort)

		userRepo := userRepo.NewUserRepository(db)
		userUsecase := userApp.NewUserUsecase(userRepo)

		servers.RunGRPCServerOnAddr(rpcPort, func(server *grpc.Server) {
			userSvc := userGrpcServer.NewUserGrpcServer(userUsecase)
			userPb.RegisterUserServiceServer(server, userSvc)
		})
	},
}

// command to run http server
var httpCmd = &cobra.Command{
	Use: "http-server",
	Run: func(cmd *cobra.Command, args []string) {
		servers.RunHttpServer()
	},
}

func Execute() {
	rootCmd.AddCommand(userGrpcServerCmd)
	rootCmd.AddCommand(httpCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
