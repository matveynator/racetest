package main

import (
 "time"
 "racetest/pkg/config"
 "racetest/pkg/mylog"
 "racetest/pkg/proxy"
 "racetest/pkg/racetest"

)

func main() {

	settings := Config.ParseFlags()
	go MyLog.ErrorLogWorker()
  go Proxy.Run(settings)
	go Racetest.Run(settings)

for {
	    time.Sleep(10 * time.Second)
	}
}
