package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files, _ := ioutil.ReadDir("E:\\")
	for _, f := range files {
		if strings.Contains(f.Name(), ".zip") {
			err := os.Remove("E:\\" + f.Name())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(f.Name(), " removed")
		}
	}
}
