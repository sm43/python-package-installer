package main

import (
	"log"
	"net/http"

	"github.com/sm43/python-package-installer/installer"
	"github.com/sm43/python-package-installer/workerqueue"
)

func main() {
	log.Println("Starting Server...")

	installer := installer.NewInstaller()

	workerqueue.StartDispatcher(installer, 2)

	http.HandleFunc("/", installer.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
