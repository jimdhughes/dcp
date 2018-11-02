package main

import (
	"log"
	"strings"
)

type dictionaryMap map[string][]string

func main() {
	arr := []string{"deer", "deal", "dog"}
	var dictionaryMap map[string][]string
	dictionaryMap = make(map[string][]string)
	PrepareMap(arr, &dictionaryMap)
	log.Printf("Autocomplete d: %s", strings.Join(Autocomplete("d", &dictionaryMap), ","))
	log.Printf("Autocomplete de: %s", strings.Join(Autocomplete("de", &dictionaryMap), ","))
	log.Printf("Autocomplete dee: %s", strings.Join(Autocomplete("dee", &dictionaryMap), ","))
}

func PrepareMap(arr []string, dict *map[string][]string) {
	var key string
	for i := 0; i < len(arr); i++ {
		s := arr[i]
		for j := 1; j < len(s); j++ {
			key = s[0:j]
			_, ok := (*dict)[key]
			if ok == true {
				(*dict)[key] = append((*dict)[key], s)
			} else {
				(*dict)[key] = []string{s}
			}
		}
	}
}

func Autocomplete(s string, dict *map[string][]string) []string {
	val := (*dict)[s]
	if val != nil {
		return val
	}
	return []string{}
}
