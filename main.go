package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func main() {
	conf := flag.String("config", "etc/go/conf.yaml", "configuration file path")
	port := flag.Int("port", 3333, "server port")
	flag.Parse()

	file, dir := filepath.Base(*conf), filepath.Dir(*conf)

	fp, err := os.Open(fmt.Sprintf("%s/%s", dir, file))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	b, _ := ioutil.ReadAll(fp)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(b)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
