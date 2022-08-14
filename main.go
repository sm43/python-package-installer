package main

import (
	"log"
	"net/http"

	"github.com/sm43/python-package-installer/installer"
)

func main() {
	log.Println("Starting Server...")

	installer := installer.NewInstaller()
	http.HandleFunc("/", installer.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
