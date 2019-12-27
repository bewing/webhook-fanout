package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bewing/webhook-fanout/pkg/proxy"
	"github.com/bewing/webhook-fanout/pkg/fanout"
)

func fanoutHandler(f fanout.Fanout) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		receivers, _ := f.Receivers()
		for _, receiver := range receivers {
			go func(recv string) {
				_, err := proxy.Do(r, recv, 10*time.Second)
				if err != nil {
					fmt.Printf("%+v\n", err)
				}
			}(receiver)
		}
	}
	return http.HandlerFunc(fn)
}
func main() {
	f, _ := fanout.NewRandomFanout()
	mux := http.NewServeMux()
	fh := fanoutHandler(f)
	mux.Handle("/", fh)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
