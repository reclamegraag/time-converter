package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type TimeZone string

const NL TimeZone = "Europe/Berlin"
const IN TimeZone = "Asia/Calcutta"
const US TimeZone = "America/Indianapolis"
const SG TimeZone = "Asia/Singapore"

func main() {
	cliArguments := os.Args[1:]
	if len(cliArguments) != 4 {
		log.Fatal(errors.New("this only works after filling in: a FROM country code, a TO country code, hours, minutes - separated by spaces"))
	}
	fromCountryCode := cliArguments[0]
	toCountryCode := cliArguments[1]
	hours, err := strconv.Atoi(cliArguments[2])
	if err != nil {
		log.Fatal(errors.New("please insert hour of the day in number format"))
	}
	minutes, err := strconv.Atoi(cliArguments[3])
	if err != nil {
		log.Fatal(errors.New("please insert minute of the day in number format"))
	}

	fromTimeZone := getTimeZone(fromCountryCode)
	toTimeZone := getTimeZone(toCountryCode)
	nowOnLocation := convertTime(hours, minutes, fromTimeZone, toTimeZone)
	nowOnLocationString := fmt.Sprintf("%s %s", nowOnLocation.String()[:19], nowOnLocation.String()[26:])
	fmt.Print(nowOnLocationString, "\n")
}

func convertTime(hours int, minutes int, fromCountryCode TimeZone, toCountryCode TimeZone) time.Time {
	fromLocation, _ := time.LoadLocation(string(fromCountryCode))
	currentZoneTime := time.Now()
	currentTime := currentZoneTime.In(fromLocation)
	currentTimeZone, _ := currentTime.Zone()
	yearDays := strconv.Itoa(currentTime.Year())[2:]
	dateString := fmt.Sprintf("%02d %s %s", currentZoneTime.Day(), currentZoneTime.Month().String()[:3], yearDays)
	nowString := fmt.Sprintf("%s %02d:%02d %s", dateString, hours, minutes, currentTimeZone)
	nowTime, _ := time.ParseInLocation(time.RFC822, nowString, fromLocation)
	utcTime := nowTime.UTC()
	toLocation, _ := time.LoadLocation(string(toCountryCode))
	toTime := utcTime.In(toLocation)

	return toTime
}

/* Function for converting a country code into a time zone */
func getTimeZone(input string) TimeZone {
	var timeZone TimeZone
	switch strings.ToUpper(input) {
	case "NL":
		timeZone = NL
	case "SG":
		timeZone = SG
	case "US":
		timeZone = US
	case "IN":
		timeZone = IN
	default:
		log.Fatal(errors.New("this country code is not part of this tool yet"))
	}
	return timeZone
}
