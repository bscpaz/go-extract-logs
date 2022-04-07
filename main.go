package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile := "logs.log"
	outputFile := "output.log"
	searchKey := "default task-63"

	txtLines := findLinesWithSearchKey(inputFile, searchKey)
	generateFileWithSearchKey(outputFile, txtLines)

	fmt.Println("Done!")
}

func findLinesWithSearchKey(fileName string, keySearch string) []string {
	var txtlines []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keySearch) {
			txtlines = append(txtlines, line)
		}
	}

	file.Close()
	return txtlines
}

func generateFileWithSearchKey(outputFile string, txtLines []string) {
	//Check if file exists. If true, then remove it.
	if _, err := os.Stat(outputFile); err == nil {
		err := os.Remove(outputFile)
		if err != nil {
			fmt.Println(err)
		}
	}

	//Create a new file.
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	datawriter := bufio.NewWriter(file)

	for _, data := range txtLines {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
