package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

//var showNumLines = flag.Bool("n", false, "flat to output the line numbers")

func readLines(r io.Reader) {

	rd := bufio.NewReader(r)
	lineNum := 1
	for {
		line, err := rd.ReadString('\n')
		fmt.Printf("\r%s", line)

		if err == io.EOF {
			fmt.Printf("%s\n", line)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		/*if *ShowNumLines {
			fmt.Printf("\r%s", line)
		}*/
		lineNum++
	}
}

func Cat(filename string) {
	file, err := os.Open(filename)
	//fmt.Println(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	readLines(file)
}
