package main

import (
	"etcd-web-ui/cmd/start"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	rootCmd := &cobra.Command{
		Short: "etcd api server",
	}
	rootCmd.AddCommand(start.Command)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
