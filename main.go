package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	conf := flag.String("config", "etc/go/conf.yaml", "configuration file path")
	flag.Parse()

	file, dir := filepath.Base(*conf), filepath.Dir(*conf)

	fp, err := os.Open(fmt.Sprintf("%s/%s", dir, file))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	b, _ := ioutil.ReadAll(fp)
	fmt.Println(string(b))
}
