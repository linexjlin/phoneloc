package main

import (
	"encoding/json"
	"flag"
	"github.com/xluohome/phonedata"
	"log"
	"net/http"
)

func location(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mobile := r.FormValue("mobile")
	if mobile != "" {
		if pr, e := phonedata.Find(mobile); e != nil {
			log.Println(e)
			http.NotFound(w, r)
		} else {
			if jdat, e := json.Marshal(pr); e != nil {
				log.Println(e)
				http.NotFound(w, r)
			} else {
				w.Write(jdat)
			}
		}
	}
}

func main() {
	addr := flag.String("addr", "0.0.0.0:8001", `-addr=0.0.0.0:8001`)
	flag.Parse()
	http.HandleFunc("/location", location)
	log.Println("listen on:", *addr)
	if e := http.ListenAndServe(*addr, nil); e != nil {
		panic(e)
	} else {
		log.Println("listen on:", *addr)
	}
}
