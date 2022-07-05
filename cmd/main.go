package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"prometheus-metrics-exporter/internal/configuration"
	"prometheus-metrics-exporter/internal/instrument"
	"prometheus-metrics-exporter/internal/pmeparser"
	"prometheus-metrics-exporter/internal/requester"
	"prometheus-metrics-exporter/internal/types"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var cfg *types.Config

func pushFailedRequest(jobName string) {

	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "failed_call",
		Help: "The timestamp of the last failed call",
	})
	completionTime.SetToCurrentTime()
	if err := push.New(cfg.PushGatewayUrl, "prometheus_metrics_exporter").
		Collector(completionTime).
		Grouping("jobname", jobName).
		Push(); err != nil {
		fmt.Println("Could not push failed call time to Pushgateway:", err)
	}

}

func init() {

	appName := "prometheus-metrics-exporter"
	appVersion := "v1.1.3"

	configPath := flag.String("config", "", "the path to the configuration file")
	help := flag.Bool("help", false, "prints this help text ")
	version := flag.Bool("version", false, "prints the version of the application")

	flag.Parse()

	if *version {
		str := fmt.Sprintf("%s: %s", appName, appVersion)
		println(str)
		os.Exit(0)
	} else if *help {
		flag.PrintDefaults()
		os.Exit(0)
	} else if *configPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var err error
	cfg, err = configuration.LoadConfig(*configPath)

	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	var wg1 sync.WaitGroup

	for _, scrapeTarget := range cfg.ScrapeTargets {

		registry := push.New(cfg.PushGatewayUrl, scrapeTarget.JobName)

		if registry == nil {
			errStr := fmt.Sprintf("Unable to create prometheus registry. PushGatewayUrl: %s. JobName: %s",
				cfg.PushGatewayUrl, scrapeTarget.JobName)
			log.Println(errStr)
			continue
		}

		wg1.Add(1)

		go func(sTarget types.ScrapeTarget, pusher *push.Pusher) {
			for _, l := range sTarget.Labels {
				pusher.Grouping(l.Name, l.Value)
			}

			content, contentType, err := requester.GetContent(sTarget.Url, sTarget.BasicAuth, sTarget.MimeType, sTarget.TimeoutInSecs)

			if err != nil {
				pushFailedRequest(sTarget.JobName)
				log.Println(err)
				wg1.Done()
				return
			}

			hasMetrics := false

			for _, m := range sTarget.Metrics {
				value, err := pmeparser.FetchValue(sTarget.Url, m.Path, content, contentType, m.Regex)

				if err != nil {
					log.Println(err)
					continue
				}

				i, err := instrument.CreateInstrument(m.InstrumentType, m.Path, m.Name, m.Help, value)

				if err != nil {
					log.Println(err)
					continue
				}

				pusher.Collector(i)

				if !hasMetrics {
					hasMetrics = true
				}

			}

			if hasMetrics {
				err = instrument.Push(sTarget.Url, pusher)
			}

			if err != nil {
				log.Println(err)
			}

			wg1.Done()

		}(scrapeTarget, registry)

	}

	wg1.Wait()

}
