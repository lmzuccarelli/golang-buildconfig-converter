package main

import (
	"flag"
	"log"
	"os"

	"github.com/luigizuccarelli/golang-buildconfig-converter/pkg/service"
)

var (
	configFile string
	help       string
)

func init() {
	flag.StringVar(&configFile, "c", "", "Use config file - helps limit input parameters")
	flag.StringVar(&help, "h", " ", "Display usage")
}

func main() {
	flag.Parse()
	if help == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := service.Convert(configFile)
	if err != nil {
		log.Fatalf("ERROR: generating shipwright buid files %v", err)
		os.Exit(1)
	}
	log.Print("INFO: completed generating shipwright build files")
}
