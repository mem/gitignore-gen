package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const localGitignore = ".gitignore.local"

type config struct {
	Inputs       []string
	localDataDir string
	configFn     string
}

func main() {
	var config config

	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.StringVar(&config.configFn, "config-filename", ".gitignore.yaml", "configuration filename")
	fs.StringVar(&config.localDataDir, "local-data-dir", ".", "directory containing gitignore data")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Println("E: invalid arguments.")
		fs.Usage()
		os.Exit(1)
	}

	if config.localDataDir == "" {
		fs.Usage()
		os.Exit(1)
	}

	cfg, err := os.ReadFile(config.configFn)
	if err != nil {
		log.Fatalf("E: Cannot read configuration file %s: %s", config.configFn, err)
	}

	if err := yaml.Unmarshal(cfg, &config); err != nil {
		log.Fatalf("E: Cannot unmarshal configuration file %s: %s", config.configFn, err)
	}

	var output bytes.Buffer

	output.WriteString("# FILE GENERATED USING gitignore-gen, DO NOT EDIT\n\n")

	for i, ignoreId := range config.Inputs {
		if strings.HasPrefix(ignoreId, "https://") || strings.HasPrefix(ignoreId, "http://") {
			processURL(&output, ignoreId, i > 0)
		} else {
			ignoreFn := filepath.Join(config.localDataDir, ignoreId)

			processFile(&output, ignoreFn, ignoreId, i > 0)
		}
	}

	if _, err := os.Stat(localGitignore); err == nil {
		processFile(&output, localGitignore, localGitignore, len(config.Inputs) > 0)
	}

	os.Stdout.Write(output.Bytes())
}

func processURL(output *bytes.Buffer, u string, outputSep bool) {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatalf("E: Cannot process URL %q: %s", u, err)
	}

	defer resp.Body.Close()

	processData(output, u, resp.Body, outputSep)
}

func processFile(output *bytes.Buffer, fn, id string, outputSep bool) {
	data, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalf("E: Cannot read ignore file %q: %s", fn, err)
	}

	processData(output, id, bytes.NewBuffer(data), outputSep)
}

func processData(output *bytes.Buffer, id string, r io.Reader, outputSep bool) {
	if outputSep {
		output.WriteRune('\n')
	}

	output.WriteString("# ")
	output.WriteString(id)
	output.WriteRune('\n')
	io.Copy(output, r)
	output.WriteRune('\n')
}
