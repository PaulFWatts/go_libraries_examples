package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "hello",
        Short: "Prints Hello World",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Hello, world!")
        },
    }

    rootCmd.Execute()
}
