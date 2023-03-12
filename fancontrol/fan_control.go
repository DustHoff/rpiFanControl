package fancontrol

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
)

var pin rpio.Pin
var speed int

func InitFanControl(pinNo int) {
	rpio.Open()
	pin = rpio.Pin(pinNo)
	pin.Mode(rpio.Pwm)
	pin.Freq(25000)
	pin.DutyCycleWithPwmMode(0, 100, rpio.Balanced)
}

func SetSpeed(speedValue int) {
	pin.DutyCycleWithPwmMode(uint32(speedValue), 100, rpio.Balanced)
	speed = speedValue
	log.Println("changed speed to duty cycle " + fmt.Sprint(speed) + "%")
}

func GetSpeed() float64 {
	log.Println("debug" + " " + fmt.Sprint(speed) + "" + fmt.Sprint(float64(speed)))
	return float64(speed)
}
