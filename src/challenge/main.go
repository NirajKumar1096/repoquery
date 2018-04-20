package main

import (
	"challenge/util"
	"log"
	"net/http"
)

func main() {

	router := util.NewRouter()
	util.Init()
	log.Fatal(http.ListenAndServe(":8088", router))

}
