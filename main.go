package main

import (
	"FanControl/fancontrol"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	http.Handle("/control", fancontrol.NewFanControl(19))
	http.Handle("/metric", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
