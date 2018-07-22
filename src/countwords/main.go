package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	"sort"
)

type Word struct {
	name string
	count int
}

type Collection struct {
	words []Word
}

func (collection* Collection) add(word string) {
	filter := func() (*Word) {
		for i, v := range collection.words {
			if v.name == word {
				return &collection.words[i]
			}
		}
		return nil
	}
	foundWord := filter()
	if foundWord == nil {
		collection.words = append(collection.words, Word{word,1})
	} else {
		foundWord.count++
	}
}

func (collection* Collection) sortDesc() {
	sort.Slice(collection.words, func(i, j int) bool {
		return collection.words[i].count > collection.words[j].count
	})
}

func main()  {
	path := flag.String("path","textfile","Path to input file that contains text")
	flag.Parse()
	data, error := ioutil.ReadFile(*path)
	if error != nil {
		println("Error: File could not be read")
		os.Exit(1)
	}
	words := createWordSlice(string(data))
	coll := Collection{make([]Word,0)}
	for _,v := range words {
		coll.add(v)
	}
	coll.sortDesc()
	for _,v := range coll.words {
		fmt.Println(v.name,v.count)
	}
}

func createWordSlice(input string) []string{
	replaced := strings.Replace(input,"\n", " ", -1)
	replaced = strings.Trim(replaced, " ")
	return strings.Split(replaced," ")
}