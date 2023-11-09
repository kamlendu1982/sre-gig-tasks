/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "get a pet joke",
	Long:  `make sure the joke is not offending`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called a joke")
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	fmt.Println("Get random pet joke")
}
