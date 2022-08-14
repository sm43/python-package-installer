package installer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Installer struct {
	diskTargetLocation string
}

type response struct {
	Message string `json:"message"`
}

func NewInstaller() *Installer {
	targetLocation, _ := os.UserHomeDir()
	if targetLocation != "" {
		targetLocation = filepath.Join(targetLocation, "python-packages")

		if err := os.MkdirAll(targetLocation, os.ModePerm); err != nil {
			fmt.Println("failed to create target directory: ", err)
		}
	}
	return &Installer{
		diskTargetLocation: targetLocation,
	}
}

func (i *Installer) Handler() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("failed to read body : %v", err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		var body map[string]interface{}
		if string(payload) != "" {
			if err := json.Unmarshal(payload, &body); err != nil {
				fmt.Printf("Invalid event body format format: %s", err)
				response.WriteHeader(http.StatusBadRequest)
			}
		}

		if body["command"] == nil {
			responseWriter(http.StatusBadRequest, "invalid request, missing pip command in body!", response)
			return
		}

		command := strings.TrimSpace(fmt.Sprint(body["command"]))
		pkgs := strings.Split(command, " ")
		if !strings.HasPrefix(command, "pip install ") || len(pkgs) <= 2 {
			responseWriter(http.StatusBadRequest, "invalid pip command", response)
			return
		}

		for ctr := 2; ctr < len(pkgs); ctr++ {
			if trimmed := strings.TrimSpace(pkgs[ctr]); len(trimmed) == 0 {
				continue
			}

			if i.isAlreadyInstalled(pkgs[ctr]) {
				responseWriter(http.StatusCreated, fmt.Sprintf("package %s already installed!", pkgs[ctr]), response)
				return
			}

			go i.installPackage(pkgs[ctr])
		}
		responseWriter(http.StatusCreated, fmt.Sprintf("installing package: %v", pkgs[2:]), response)
	})
}

func responseWriter(statusCode int, message string, resp http.ResponseWriter) {
	body := response{
		Message: message,
	}
	resp.WriteHeader(statusCode)
	if err := json.NewEncoder(resp).Encode(body); err != nil {
		fmt.Println("failed to encode response : ", err)
	}
}
