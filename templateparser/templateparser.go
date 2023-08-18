package main

import (
	"fmt"
	"net"
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/tiagomelo/golang-grpc-backpressure/config"
)

type data struct {
	IP         string
	Port       int
	ClientPort int
}

func getOutboundIpAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

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

func parsePrometheusScrapeTemplate(ip string, serverPort, clientPort int, templateFile, outputFile string) error {
	data := &data{
		IP:         ip,
		Port:       serverPort,
		ClientPort: clientPort,
	}
	return parseTemplate(data, templateFile, outputFile)
}

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
