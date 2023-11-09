/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"

	"github.com/spf13/cobra"
)

//var ShowNumLines = flag.Bool("n", false, "flat to output the line numbers")

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "implementation like bash cat",
	Long:  `implementation like bash cat`,
	Run: func(cmd *cobra.Command, args []string) {
		//nFlag, _ := cmd.Flags().GetString("n")
		flag.Parse()
		for i := 1; i < flag.NArg(); i++ {
			//fmt.Println(flag.Arg(i))
			Cat(flag.Arg(i))
		}
		/*if nFlag != "" {
			fmt.Println("cat called")
		} else {
			Cat(nFlag)
		}*/
	},
}

func init() {
	catCmd.Flags().String("n", "", "Name of the file to display content of file")
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
