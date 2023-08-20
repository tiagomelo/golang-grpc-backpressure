// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
//
// Package main contains a utility for generating configurations based on
// templates and specific data, e.g., IP addresses and ports.
package main

import (
	"fmt"
	"net"
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/tiagomelo/golang-grpc-backpressure/config"
)

// data is a struct that holds the information used to fill the templates.
type data struct {
	IP         string
	Port       int
	ClientPort int
}

// getOutboundIpAddr returns the outbound IP address of the current machine.
func getOutboundIpAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

// parseTemplate takes in a data object, a template file, and an output file.
// It parses the template, fills it with data, and writes the resulting configuration to the output file.
func parseTemplate(data *data, templateFile, outputFile string) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return errors.Wrapf(err, `parsing template file "%s"`, templateFile)
	}
	file, err := os.Create(outputFile)
	if err != nil {
		return errors.Wrapf(err, `creating output file "%s"`, outputFile)
	}
	defer file.Close()
	if err = tmpl.Execute(file, data); err != nil {
		return errors.Wrapf(err, `executing template file "%s"`, templateFile)
	}
	return nil
}

// parsePrometheusScrapeTemplate is a specialized function to generate Prometheus scrape configurations.
// It sets up data based on provided parameters and then uses the general template parsing function.
func parsePrometheusScrapeTemplate(ip string, serverPort, clientPort int, templateFile, outputFile string) error {
	data := &data{
		IP:         ip,
		Port:       serverPort,
		ClientPort: clientPort,
	}
	return parseTemplate(data, templateFile, outputFile)
}

// parsePrometheusDataSourceTemplate is another specialized function to generate Prometheus datasource configurations.
// It sets up data based on provided parameters and then uses the general template parsing function.
func parsePrometheusDataSourceTemplate(ip string, serverPort int, templateFile, outputFile string) error {
	data := &data{
		IP:   ip,
		Port: serverPort,
	}
	return parseTemplate(data, templateFile, outputFile)
}

func run() error {
	cfg, err := config.Read()
	if err != nil {
		return errors.Wrap(err, "reading config")
	}
	ip, err := getOutboundIpAddr()
	if err != nil {
		return errors.Wrap(err, "getting ip")
	}
	if err := parsePrometheusScrapeTemplate(ip, cfg.PromTargetGrpcServerPort,
		cfg.PromTargetGrpcClientPort, cfg.PromTemplateFile, cfg.PromOutputFile); err != nil {
		return errors.Wrap(err, "parsing Prometheus scrape template")
	}
	if err := parsePrometheusDataSourceTemplate(ip, cfg.DsServerPort, cfg.DsTemplateFile, cfg.DsOutputFile); err != nil {
		return errors.Wrap(err, "parsing Prometheus datasource template")
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
