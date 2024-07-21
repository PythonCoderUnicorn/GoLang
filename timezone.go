package main

import (
	"fmt"
	"time"
)

// https://www.timeanddate.com/worldclock/
// https://golangbyexample.com/convert-time-timezones-go/


func main() {

	fmt.Println("----------[ Time ]-----------")


	// Get the current time
  // currentTime := time.Now()
	now := time.Now()

	fmt.Printf("Current time : %02d:%02d:%02d <-- \n", now.Hour(), now.Minute(), now.Second() )

	Central_time, _ := time.LoadLocation("America/Chicago")
	fmt.Printf("Central Time : %02d:%02d:%02d \n", now.In(Central_time).Hour(),now.In(Central_time).Minute(), now.In(Central_time).Second())



	NYC_time, _ := time.LoadLocation("America/New_York")
	fmt.Printf("NYC Time     : %02d:%02d:%02d \n", now.In(NYC_time).Hour(),now.In(NYC_time).Minute(), now.In(NYC_time).Second())

	UTC_time, _ := time.LoadLocation("UTC")
  // fmt.Printf("UTC Time: %s \n", now.In(UTC_time) )
  fmt.Printf("UTC Time     : %02d:%02d:%02d \n", now.In(UTC_time).Hour(),now.In(UTC_time).Minute(), now.In(UTC_time).Second())

	London_time, _ := time.LoadLocation("Europe/London")
  fmt.Printf("London Time  : %02d:%02d:%02d \n", now.In(London_time).Hour(),now.In(London_time).Minute(), now.In(London_time).Second())


	Paris_time, _ := time.LoadLocation("Europe/Paris")
  fmt.Printf("Paris Time   : %02d:%02d:%02d \n", now.In(Paris_time).Hour(),now.In(Paris_time).Minute(), now.In(Paris_time).Second())


	// Moscow_time, _ := time.LoadLocation("Europe/Moscow")
  // fmt.Printf("UTC Time: %s \n", now.In(UTC_time) )
  // fmt.Printf("Moscow Time   : %02d:%02d:%02d \n", now.In(India_time).Hour(),now.In(India_time).Minute(), now.In(India_time).Second())



	fmt.Println("")
}