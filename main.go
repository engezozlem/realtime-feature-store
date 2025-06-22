package main

import (
	"log"
	"net/http"

	"github.com/engezozlem/realtime-feature-store/api"
	"github.com/engezozlem/realtime-feature-store/store"
)

func main() {
	store.InitRedis()
	http.HandleFunc("/features/", api.FeatureHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
