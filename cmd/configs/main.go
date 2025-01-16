package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/montybeatnik/lab"
)

func main() {
	labR1 := "lab-r1"
	cfg, err := buildCfg(labR1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cfg)
}

func buildCfg(hostname string) (string, error) {
	dev := lab.Device{
		Hostname: hostname,
		Interfaces: []lab.Interface{
			{
				IFD:     "ge-0/0/0",
				IFL:     "0",
				Address: "192.168.1.0/31",
				Role:    "infrastructure",
			},
			{
				IFD:     "ge-0/0/1",
				IFL:     "0",
				Address: "192.168.1.2/31",
				Role:    "infrastructure",
			},
		},
	}
	t, err := template.ParseGlob("./cmd/configs/templates/*")
	if err != nil {
		return "", fmt.Errorf("failed to get templates: %w", err)
	}
	var config bytes.Buffer
	t.ExecuteTemplate(&config, "base", dev)
	return config.String(), nil
}
