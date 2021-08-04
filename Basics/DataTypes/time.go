package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		speed int
		now   time.Time
	)

	speed, now = 200, time.Now()

	fmt.Println(speed, now)

	var(
		newSpeed = 100
		prevSpeed = 50
	)
   newSpeed, prevSpeed = prevSpeed, newSpeed
   fmt.Println(newSpeed, prevSpeed)
}
