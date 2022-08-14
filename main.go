package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	log.Println("Starting Server...")

	http.HandleFunc("/", installPackage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func installPackage(w http.ResponseWriter, r *http.Request) {
	cmd, err := exec.Command("/bin/sh", "installer.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	output := string(cmd)
	fmt.Println("cmd output: ", output)

	fmt.Fprint(w, "package installed!")
}
