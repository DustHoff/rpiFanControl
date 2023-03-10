package main

import (
	"FanControl/fancontrol"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	http.Handle("/control", fancontrol.NewFanControl(13))
	http.Handle("/metric", promhttp.Handler())
	fmt.Println("start listen port 8080")
	http.ListenAndServe(":8080", nil)
}
