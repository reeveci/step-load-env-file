package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	reeveAPI := os.Getenv("REEVE_API")
	if reeveAPI == "" {
		fmt.Println("This docker image is a Reeve CI pipeline step and is not intended to be used on its own.")
		os.Exit(1)
	}

	var params []string
	err := json.Unmarshal([]byte(os.Getenv("REEVE_PARAMS")), &params)
	if err != nil {
		panic(fmt.Sprintf("error parsing REEVE_PARAMS - %s", err))
	}

	envFile := os.Getenv("FILE")
	if envFile == "" {
		panic("missing env file name")
	}

	loadAll := os.Getenv("LOAD_ALL") == "true"

	var env map[string]string
	env, err = godotenv.Read(filepath.Join("/reeve/src", envFile))
	if err != nil {
		panic(fmt.Sprintf("error loading \"%s\" - %s", envFile, err))
	}

	var varNames map[string]string
	if loadAll {
		varNames = make(map[string]string, len(env))
		for name := range env {
			varNames[name] = name
		}
	} else {
		varNames = make(map[string]string, len(params))
		for _, param := range params {
			if !strings.HasPrefix(param, "ENV_") {
				continue
			}

			name := strings.TrimPrefix(param, "ENV_")
			varName := os.Getenv(param)
			if name != "" && varName != "" {
				varNames[name] = varName
			}
		}
	}

	for name, varName := range varNames {
		value := env[name]
		response, err := http.Post(fmt.Sprintf("%s/api/v1/var?key=%s", reeveAPI, url.QueryEscape(varName)), "text/plain", strings.NewReader(value))
		if err != nil {
			panic(fmt.Sprintf("setting var returned status %v", response.StatusCode))
		}
		fmt.Printf("Set %s=%s\n", varName, value)
	}
}
