package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	t := time.Now().UTC().UnixNano()
	arr := strings.Fields(text)
	t = report_elapse(t, "split words")
	prefix1, prefix2, suffix := NOWORD, NOWORD, NOWORD
	tab := make(map[string][]string, 10240)
	for i := range arr {
		prefix1 = prefix2
		prefix2 = suffix
		suffix = arr[i]
		insertSuffix(tab, buildKey(prefix1, prefix2), suffix)
	}
	report_elapse(t, "build hash")
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

func report_elapse(t int64, title string) int64 {
	t2 := time.Now().UTC().UnixNano()
	fmt.Fprintf(os.Stderr, "%s elapsed = %d ns\n", title, t2-t)
	return t2
}

func main() {
	filename := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	check(err)
	t := time.Now().UTC().UnixNano()
	text, err := ioutil.ReadFile(filename)
	check(err)
	t = report_elapse(t, "read file")
	prefixes := BuildIndice(string(text))
	t = report_elapse(t, "build indice")
	generate(prefixes, count)
	t = report_elapse(t, "generate output")
}
