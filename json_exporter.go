package main

import (
	"fmt"
	"github.com/kawamuray/prometheus-exporter-harness/harness"
	"github.com/AllocateSoftware/prometheus-allocate-exporter/jsonexporter"
	"github.com/AllocateSoftware/prometheus-allocate-exporter/allocateexporter"
)

func main() {

	fmt.Println("Start me up")

	allocateexporter.Find_services

	fmt.Println("Started me up")


	opts := harness.NewExporterOpts("json_exporter", jsonexporter.Version)
	opts.Usage = "[OPTIONS] HTTP_ENDPOINT CONFIG_PATH"
	opts.Init = jsonexporter.Init
	harness.Main(opts)
}
