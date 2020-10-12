package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello!")
	})

	fileServe := http.FileServer(http.Dir("../static"))
	http.Handle("/", fileServe)
	port := os.Args[1]
	i, _ := strconv.Atoi(port)
	fmt.Print("Server Running at port: ", i)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}

}
