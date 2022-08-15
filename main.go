package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sm43/python-package-installer/installer"
	"github.com/sm43/python-package-installer/workerqueue"
)

func main() {
	NWorkers := flag.Int("number-of-workers", 2, "The number of workers to start")
	flag.Parse()

	log.Println("Initializing Installer...")
	installer := installer.NewInstaller()

	log.Println("Starting Dispatcher...")
	workerqueue.StartDispatcher(installer, *NWorkers)

	http.HandleFunc("/", installer.Handler())

	log.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
