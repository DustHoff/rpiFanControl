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

var _ http.Handler = &FanControl{}

type FanControl struct {
	maxSpeed int
	minSpeed int
	speed    int
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
			log.Println(err)
		}
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
	fc.pin.DutyCycleWithPwmMode(uint32(speed), 100, rpio.Balanced)
	fc.speed = speed
	log.Println("changed speed to duty cycle " + fmt.Sprint(fc.speed) + "%")
}

func (fc FanControl) GetSpeed() float64 {
	log.Println("debug" + " " + fmt.Sprint(fc.speed) + "" + fmt.Sprint(float64(fc.speed)))
	return float64(fc.speed)
}
