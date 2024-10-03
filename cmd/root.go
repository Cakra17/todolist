package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "basic",
  Short: "short description about this fuckin cli",
  Long: "long description about this fuckin cli that contains multiple lines and examples or usage of your application",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("basic manja")
  },
}

func Execute() {
  err := rootCmd.Execute()
  if err != nil {
    os.Exit(1)
  }
}

func init() {
  rootCmd.Flags().BoolP("toggle", "t", false, "coba manja aja")
}
