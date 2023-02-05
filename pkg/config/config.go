package Config

import (
	"os"
	"fmt"
	"flag"
	"time"

)

type Settings struct {
	APP_NAME, VERSION, PROXY_ADDRESS, TIME_ZONE string
	RIDERS, RESULTS, LAPS int
	RANDOM bool
	RACE_TIMEOUT_DURATION, MINIMAL_LAP_TIME_DURATION time.Duration
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func ParseFlags() (config Settings)  { 
	config.APP_NAME = "racetest"
	flagVersion := flag.Bool("version", false, "Output version information")
	flag.StringVar(&config.PROXY_ADDRESS, "target", "127.0.0.1:4000", "Provide IP address and port where to send data (where chicha timekeeper is running).")
	flag.IntVar(&config.RIDERS, "riders", 10, "Set ammount of riders (sportsmen) for the race.")
	flag.IntVar(&config.RESULTS, "results", 10, "Set ammount of results each competitor sends to timekeeper.")
	flag.IntVar(&config.LAPS, "laps", 5, "Set ammount of laps each race holds.")
	flag.DurationVar(&config.RACE_TIMEOUT_DURATION, "timeout", 2*time.Minute, "Set race timeout duration. After this time if nobody passes the finish line the race will be stopped. Valid time units are: 's' (second), 'm' (minute), 'h' (hour).")
	flag.DurationVar(&config.MINIMAL_LAP_TIME_DURATION, "minimal-lap-time", 60*time.Second, "Minimal lap time duration. Results smaller than this duration would be considered wrong. Valid time units are: 's' (second), 'm' (minute), 'h' (hour)." )
	flag.BoolVar(&config.RANDOM, "non-random", false, "Do not send random data. By default we send random data.")


	//process all flags
	flag.Parse()


	if isFlagPassed("non-random") {
		config.RANDOM=false
	} else {
		config.RANDOM=true
	}

	if *flagVersion  {
		if config.VERSION != "" {
			fmt.Println("Version:", config.VERSION)
		}
		os.Exit(0)
	}
	return
}
