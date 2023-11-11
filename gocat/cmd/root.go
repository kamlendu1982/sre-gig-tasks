/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Number, NumberNonblank bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocat",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {

		processArgs(args)
	},
}

func processArgs(args []string) {
	var lineNum uint = 0
	for _, arg := range args {
		fileInfo, err := os.Stat(arg)
		if err != nil {
			fmt.Printf("gocat cat %s:, No such file or directory", arg)
			return
		}

		if fileInfo.IsDir() {
			fmt.Printf("gocat cat %s: is a directory", arg)
			return
		}
		lineNum = printFile(arg, lineNum)
	}
}

func printFile(arg string, lineNum uint) uint {
	file, err := os.Open(arg)
	if err != nil {
		fmt.Printf("gocat: %s: Permission denied\n", arg)
		return lineNum
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		blankLines := len(strings.TrimSpace(line)) == 0

		if NumberNonblank && !blankLines {
			fmt.Println(line)
		} else if Number {
			prefix := fmt.Sprintf("\t%5s\t", strconv.FormatUint(uint64(lineNum), 10))
			fmt.Println(prefix, line)
		} else if !Number && !NumberNonblank {
			fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
		}
		lineNum++
	}
	return lineNum
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&Number, "number", "n", false, "number all output lines")
	rootCmd.Flags().BoolVarP(&NumberNonblank, "number-nonblank", "b", false, "number nonempty output lines, overrides -n")
}
