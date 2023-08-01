package main

import (
	"fmt"
	"github.com/duckcache/cache"
	http2 "github.com/duckcache/http"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		},
	))

	addr := "localhost:8080"
	peers := http2.NewHTTPPool(addr)
	log.Println("duck-cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
