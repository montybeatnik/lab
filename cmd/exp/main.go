package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/montybeatnik/devcon"
	"github.com/montybeatnik/devcon/vendors/juniper"
	"github.com/montybeatnik/lab"
)

type Device struct {
	Hostname   string
	IP         string
	ConfigPath string
}

func main() {
	var wg sync.WaitGroup
	store := lab.NewStore()
	cfgs := store.GetConfigs()
	wg.Add(len(cfgs))
	for _, cfg := range cfgs {
		log.Printf("applying config on %v", cfg.Device.Hostname)
		go func(cfgfile string, ip string) {
			defer wg.Done()
			cfg := getConfig(cfgfile)
			dev := getDevice(ip)
			// output, err := dev.Diff(cfg)
			output, err := dev.ApplyConfig(cfg)
			if err != nil {
				log.Printf("cfg failed on %v; %v\n", ip, err)
			}
			fmt.Println(output)

		}(cfg.Path, cfg.Device.MGMTAddress)
	}
	wg.Wait()
}

func getConfig(fn string) []string {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cfg []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cfg = append(cfg, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cfg
}

func getDevice(ip string) *juniper.JuniperClient {
	khfp := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	jnprClient := juniper.NewJuniperClient(
		os.Getenv("SSH_USER"),
		ip,
		devcon.WithPassword(os.Getenv("SSH_PASSWORD")),
		devcon.WithTimeout(time.Second*1),
		devcon.WithHostKeyCallback(khfp),
	)
	return jnprClient
}
