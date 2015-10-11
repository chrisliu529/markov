package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

const NOWORD string = " "

func insertSuffix(tab map[string][]string, key string, suffix string) {
	if v, found := tab[key]; found {
		tab[key] = append(v, suffix)
	} else {
		tab[key] = []string{suffix}
	}
}

func buildKey(prefix1 string, prefix2 string) string {
	return fmt.Sprintf("%s %s", prefix1, prefix2)
}

func BuildIndice(text string) map[string][]string {
	words := regexp.MustCompile("\\w+")
	arr := words.FindAllString(text, -1)
	prefix1, prefix2, suffix := NOWORD, NOWORD, NOWORD
	tab := make(map[string][]string, 10240)
	for i := range arr {
		prefix1 = prefix2
		prefix2 = suffix
		suffix = arr[i]
		insertSuffix(tab, buildKey(prefix1, prefix2), suffix)
	}
	return tab
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func choose(suffix []string) string {
	return suffix[rand.Intn(len(suffix))]
}

func generate(prefixes map[string][]string, n int) {
	prefix1, prefix2 := NOWORD, NOWORD
	rand.Seed(time.Now().UTC().UnixNano())
	num_words := 0
	for num_words < n {
		if suffix, found := prefixes[buildKey(prefix1, prefix2)]; found {
			word := choose(suffix)
			fmt.Printf("%s ", word)
			prefix1 = prefix2
			prefix2 = word
		} else {
			return
		}
		num_words++
	}
}

func main() {
	filename := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	check(err)
	text, err := ioutil.ReadFile(filename)
	check(err)
	prefixes := BuildIndice(string(text))
	generate(prefixes, count)
}
