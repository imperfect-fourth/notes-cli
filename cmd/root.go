package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "notes",
	Short: "A cli app to manage all your notes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world");
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err);
		os.Exit(1);
	}
}
