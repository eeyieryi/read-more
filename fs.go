package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	data_dir := os.Getenv("DATA_DIR")
	port := os.Getenv("FS_PORT")
	fs := http.FileServer(http.Dir(data_dir))
	http.Handle("/", fs)
	fmt.Println("Serving files from", data_dir)
	fmt.Println("Starting server on port :", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
