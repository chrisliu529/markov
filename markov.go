package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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
	arr := strings.Split(strings.Trim(text, " \n"), " ")
	if len(arr) > 2 {
		lim := len(arr) - 2
		i := 0
		tab := make(map[string][]string, 10240)
		for i < lim {
			insertSuffix(tab, buildKey(arr[i], arr[i+1]), arr[i+2])
			i++
		}
		return tab
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	text, err := ioutil.ReadFile("psalm.txt")
	check(err)
	BuildIndice(string(text))
}
