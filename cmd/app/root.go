package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang-project-template/internal/common"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "grpc-server",
	Run: func(cmd *cobra.Command, args []string) {
		// Application entrypoint...

		common.SetUpEnv()
		fmt.Println("hello world")

		// remove it when you run http/grpc server
		c := make(chan string)
		<-c
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
