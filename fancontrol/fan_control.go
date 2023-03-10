package fancontrol

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)
import "github.com/prometheus/client_golang/prometheus"
import "github.com/prometheus/client_golang/prometheus/promauto"

var (
	fanspeed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "case_fan_speed",
		Help: "The total speed of case fan",
	})
)

var _ http.Handler = &FanControl{}

type FanControl struct {
	maxSpeed int
	minSpeed int
	pin      rpio.Pin
}

func (fc FanControl) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, request.Body)
		if err != nil {
			log.Fatalln(err)
		}
		speed, err := strconv.Atoi(buf.String())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("received speed control request value:" + fmt.Sprint(speed))
		fc.SetSpeed(speed)
	}

	fmt.Fprintf(writer, "ok")
}

func NewFanControl(pinNo int) FanControl {
	rpio.Open()
	pin := rpio.Pin(pinNo)
	pin.Mode(rpio.Pwm)
	pin.Freq(25000)
	pin.DutyCycleWithPwmMode(0, 100, rpio.Balanced)
	return FanControl{
		pin: pin,
	}
}

func (fc FanControl) SetMaxSpeed(max int) {
	fc.maxSpeed = max
}

func (fc FanControl) SetMinSpeed(min int) {
	fc.minSpeed = min
}

func (fc FanControl) SetSpeed(speed int) {
	fc.pin.DutyCycle(uint32(speed), 100)
	fmt.Println("changed speed to dutycylce " + fmt.Sprint(speed) + "/100")
}
