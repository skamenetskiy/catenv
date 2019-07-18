package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	var (
		str string
	)
	if len(os.Args) != 2 {
		usage()
	}
	if os.Args[1] == "-in" {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil && err != io.EOF {
			exit(err)
		}
		str = string(b)
	} else {
		fileName := os.Args[1:2][0]
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			exit(err)
		}
		str = string(b)
	}
	if len(str) == 0 {
		exit(errors.New("no contents"))
	}
	env := getEnv()
	for k, v := range env {
		str = strings.Replace(str, k, v, -1)
	}
	fmt.Print(str)
}

func getEnv() map[string]string {
	env := make(map[string]string, 0)
	keys := make([]string, 0)
	for _, val := range os.Environ() {
		v := strings.Split(val, "=")
		if len(v) < 2 || v[0] == "" {
			continue
		}
		keys = append(keys, v[0])
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})
	for _, k := range keys {
		v := os.Getenv(k)
		env["$"+k] = v
		env["${"+k+"}"] = v
	}
	return env
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func usage() {
	fmt.Println(fmt.Sprintf("Usage: %s <filename>", os.Args[0]))
	fmt.Println(fmt.Sprintf("Usage: cat <filename> | %s -in", os.Args[0]))
	os.Exit(1)
}
