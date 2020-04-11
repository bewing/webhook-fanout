package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bewing/webhook-fanout/pkg/fanout"
	"github.com/bewing/webhook-fanout/pkg/proxy"
)

func fanoutHandler(f fanout.Fanout) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		receivers, _ := f.Receivers()
		for _, receiver := range receivers {
			go func(recv string) {
				_, err := proxy.Do(r, recv, 10*time.Second)
				if err != nil {
					fmt.Printf("%+v\n", err)
				} else {
					fmt.Printf("Relayed to %+v\n", recv)
				}
			}(receiver)
		}
	}
	return http.HandlerFunc(fn)
}
func main() {
	selector := flag.String("selector", "", "Label selector used to filter for pods")
	namespace := flag.String("namespace", "", "Namespace to search for pods (default all)")
	targetPort := flag.Int("targetPort", 80, "Listen port on Pods to fanout requests to")
	flag.Parse()

	f, err := fanout.NewPodFanout(*namespace, *selector, *targetPort)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	fh := fanoutHandler(f)
	mux.Handle("/", fh)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
