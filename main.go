package main

import (
	"FanControl/fancontrol"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	reg := prometheus.NewRegistry()
	fancontrol.InitFanControl(13)
	// Add go runtime metrics and process collectors.
	gauge := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name:      "fan_cycle_duty",
		Namespace: "rpi",
		Help:      "duty cycle of fan in percent",
	}, fancontrol.GetSpeed)
	reg.MustRegister(gauge)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		EnableOpenMetrics: false,
	}))

	api := fancontrol.Api{}
	http.Handle("/control", api)

	log.Println("start listen port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
