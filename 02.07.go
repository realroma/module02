package main

import (
	"os"
	"fmt"
	"bufio"
)

func write (in string, out string) {
	InFile, _ := os.Open(in)
        OutFile, _ := os.Create(out)
	defer InFile.Close()
	defer OutFile.Close()
	reader := bufio.NewScanner(InFile)
        writer := bufio.NewWriter(OutFile)
	defer writer.Flush()
        for reader.Scan() {
		if reader.Text() == "" {
			panic ("Plase is empty")
		}
                writer.WriteString(reader.Text() + "\n")
        }
}

func main() {
	write("InFile.txt", "OutFile")
	fmt.Println("Done")
}
