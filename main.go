package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var filePath, searchString, output string

// Reads program flags
// input || i, output || o, search || s
func readFlags() {
	flag.StringVar(&filePath, "input", "./examples/dictionary.txt", "Input file for dictionary")
	flag.StringVar(&filePath, "i", "./examples/dictionary.txt", "Input file for dictionary [short form]")
	flag.StringVar(&output, "output", "./examples/output.txt", "Output file for closest words")
	flag.StringVar(&output, "o", "./examples/output.txt", "Output file for closest words [short form]")
	flag.StringVar(&searchString, "search", "", "Search string")
	flag.StringVar(&searchString, "s", "", "Search string [short form]")

	flag.Parse()

	if searchString == "" {
		flag.PrintDefaults()
		log.Fatal("Please add a search string")
	}
}

// Reads File and adds words to dictionary
// Input -> File Reader and Trie Pointer
// Returns a channel of strings
func parseWordsToTrie(dictFile io.Reader, trie *Trie) <-chan string {
	dictScanner := bufio.NewScanner(dictFile)

	out := make(chan string)
	go func() {
		for dictScanner.Scan() {
			word := dictScanner.Text()
			trie.Insert(word)
			out <- word
		}
		close(out)
	}()

	return out
}

func main() {
	readFlags()

	dictFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Could not open file: ", filePath, err)
	}
	defer dictFile.Close()

	trie := Trie{Root: &Node{}}
	lines := parseWordsToTrie(dictFile, &trie)
	for {
		str := <-lines
		if str == "" {
			break
		}
	}
	log.Println("Finished Creating Trie...")

	closestStrings := trie.FindClosestStrings(searchString)
	if len(closestStrings) == 0 {
		log.Println("No Closest strings found")
		return
	}

	outFile, err := os.Create(output)
	if err != nil {
		log.Fatal("Could not open file: ", filePath, err)
	}
	defer outFile.Close()

	outFile.Write([]byte(fmt.Sprintf("Closest Strings to %s: \n", searchString)))
	for _, item := range closestStrings {
		outFile.Write([]byte(item + "\n"))
	}

	log.Println("Finished Writing Closest strings to ", output)
}
