// SPDX-License-Identifier: Apache-2.0
// Copyright 2020 Intel Corporation
// Copyright 2022-present Open Networking Foundation

package main

import (
	"flag"

	"github.com/omec-project/upf-epc/pfcpiface"
	log "github.com/sirupsen/logrus"
)

var (
	configPath = flag.String("config", "upf.jsonc", "path to upf config")
)

func init() {
	// Set up logger
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Info("UPF by CDAC - Jayaram R @ 04-04-2024, 14:15:00")
}

func main() {
	// cmdline args
	flag.Parse()
	log.Info("PFCP UPF - Jayaram R")
	// Read and parse json startup file.
	conf, err := pfcpiface.LoadConfigFile(*configPath)
	if err != nil {
		log.Fatalln("Error reading conf file:", err)
	}

	log.SetLevel(conf.LogLevel)

	log.Infof("%+v", conf)

	pfcpi := pfcpiface.NewPFCPIface(conf)

	// blocking
	pfcpi.Run()
}
