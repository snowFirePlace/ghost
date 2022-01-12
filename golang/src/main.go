package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	dictionary map[string]map[string]map[string]string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trans(path, source string) {
	var temp string

	if strings.ContainsAny(source, string(os.PathSeparator)) {
		temp = source
		s := strings.Split(source, string(os.PathSeparator))
		path = path + string(os.PathSeparator) + s[0]
		source = s[1]

	}
	fmt.Println(path, source)
	var sourcePathFile string

	files, err := ioutil.ReadDir(path)
	if err != nil {

		log.Fatal(err)
	}
	for _, f := range files {
		if name, _ := regexp.MatchString(source, f.Name()); name {
			sourcePathFile = path + string(os.PathSeparator) + f.Name()
			break
		}
	}
	dat, err := os.ReadFile(sourcePathFile)
	out := string(dat)
	check(err)
	if len(temp) != 0 {
		for o, m := range dictionary[temp] {
			for key, word := range m {
				// fmt.Println(fmt.Sprintf(o, key), "->", fmt.Sprintf(o, word))
				out = strings.Replace(out, fmt.Sprintf(o, key), fmt.Sprintf(o, word), -1)
				// out = strings.Replace(out, fmt.Sprintf(`,name:"%s"`, key), fmt.Sprintf(`,name:"%s"`, word), -1)
			}
			// out = strings.Replace(string(dat), word, `[2,"Материалы"]`, -1)
		}
	} else {
		for o, m := range dictionary[source] {
			for key, word := range m {
				// fmt.Println(fmt.Sprintf(o, key), "->", fmt.Sprintf(o, word))
				out = strings.Replace(out, fmt.Sprintf(o, key), fmt.Sprintf(o, word), -1)
				// out = strings.Replace(out, fmt.Sprintf(`,name:"%s"`, key), fmt.Sprintf(`,name:"%s"`, word), -1)
			}
			// out = strings.Replace(string(dat), word, `[2,"Материалы"]`, -1)
		}
	}
	err = ioutil.WriteFile(sourcePathFile, []byte(out), 0644)
	if err != nil {
		fmt.Println("Error creating", sourcePathFile)
		return
	}

}
func main() {
	dictionary = make(map[string]map[string]map[string]string)
	// * For windows
	// fileGhostJSON, err := ioutil.ReadFile("C:\\docker\\KB\\golang\\src\\ru.json")
	// * For docker
	fileGhostJSON, err := ioutil.ReadFile("./ru.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileGhostJSON, &dictionary)
	if err != nil {
		fmt.Println(dictionary)
		log.Fatal(err)
	}
	for file, _ := range dictionary {
		// * For windows

		// * For docker
		trans("./source", file)
	}
}
