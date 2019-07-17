package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		exit(errors.New("usage: catenv <filename>"))
	}
	fileName := os.Args[1:2][0]
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		exit(err)
	}
	str := string(b)
	s := new(sync.Mutex)
	env := getEnv()
	for k, v := range env {
		s.Lock()
		str = strings.Replace(str, k, v, -1)
		s.Unlock()
	}
	fmt.Print(str)
}

func getEnv() map[string]string {
	env := make(map[string]string, 0)
	for _, val := range os.Environ() {
		v := strings.Split(val, "=")
		if len(v) < 2 || v[0] == ""{
			continue
		}
		env["$"+v[0]] = v[1]
		env["${"+v[0]+"}"] = v[1]
	}
	return env
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
