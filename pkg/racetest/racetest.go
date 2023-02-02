package Racetest

import (
	"time"
	"fmt"
	"math/rand"
	"racetest/pkg/config"
	"github.com/goombaio/namegenerator"
)

func Run(config Config.Settings) () {

	var sleeptime int = 0 
	var totalSleeptime int = 0

	averageSleepSecondsInt  := int(config.MINIMAL_LAP_TIME_DURATION.Seconds())/config.RIDERS
	fmt.Println("averageSleepSecondsInt", averageSleepSecondsInt)

	//get rider names:
	competitors := riders(config)

	raceId := 0

	for {
		for lap := 0; lap <= config.LAPS; lap++ {

			if config.RANDOM {
				//shuffle riders:
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(competitors), func(i, j int) { competitors[i], competitors[j] = competitors[j], competitors[i] })
			}

			//zero total sleep time:
			totalSleeptime = 0

			for index, name := range competitors {

				if config.RANDOM {
					//not each competitor finishes each lap (75%):
					if RandBool75() == true {
						for result := 1; result <= config.RESULTS; result++ {
							//not each result registered by antenna (60%):
							if RandBool60() == true {
								//print result:
								fmt.Println(name)
							}
						}

						//if not last rider in the lap: 
						if (index + 1) != len(competitors) {

							//get random sleeptime = 1 second:
							rand.Seed(time.Now().UnixNano())
							sleeptime = rand.Intn(averageSleepSecondsInt)

							//count total sleeptime
							totalSleeptime = (totalSleeptime + sleeptime)

							fmt.Printf("next rider in %d seconds...  index=%d, total riders=%d\n", sleeptime, index, len(competitors))

							//sleep:
							time.Sleep(time.Duration(sleeptime) * time.Second)
						} else {
							//if not the last lap:
							if lap != config.LAPS {
								//if last rider in the lap:
								//calculate time left for the lap (to sleep after all riders checked in):
								if time.Duration(totalSleeptime) * time.Second < config.MINIMAL_LAP_TIME_DURATION {
									totalSleeptimeDuration := time.Duration(totalSleeptime) * time.Second
									nextLapSleepTime := config.MINIMAL_LAP_TIME_DURATION.Seconds() - totalSleeptimeDuration.Seconds() 

									fmt.Printf("Lap #%d finished, next lap in %d seconds, total laps=%d.\n", lap, int(nextLapSleepTime), config.LAPS)

									time.Sleep(time.Duration(int(nextLapSleepTime)) * time.Second)
								} else {

									fmt.Printf("Lap #%d finished, next lap in %d seconds, total laps=%d.\n", lap, 0, config.LAPS)

								}
							} else {
								//if last lap:
								fmt.Printf("Race #%d finished, next race in %d seconds.\n", raceId, int(config.RACE_TIMEOUT_DURATION.Seconds()))
								time.Sleep(config.RACE_TIMEOUT_DURATION)
								//increment raceId
								raceId++
							}
						}
					}
				} else {
					//if selected not random execution:

					for result := 1; result <= config.RESULTS; result++ {
						//print result:
						fmt.Println(name)
					}

					//if not last rider in the lap:
					if (index + 1) != len(competitors) {

						//get random sleeptime = 1 second:
						rand.Seed(time.Now().UnixNano())
						sleeptime = rand.Intn(averageSleepSecondsInt)

						//count total sleeptime
						totalSleeptime = (totalSleeptime + sleeptime)

						fmt.Printf("next rider in %d seconds...  index=%d, total riders=%d\n", sleeptime, index, len(competitors))

						//sleep:
						time.Sleep(time.Duration(sleeptime) * time.Second)
					} else {
						//if not the last lap:
						if lap != config.LAPS {
							//if last rider in the lap:
							//calculate time left for the lap (to sleep after all riders checked in):
							if time.Duration(totalSleeptime) * time.Second < config.MINIMAL_LAP_TIME_DURATION {
								totalSleeptimeDuration := time.Duration(totalSleeptime) * time.Second
								nextLapSleepTime := config.MINIMAL_LAP_TIME_DURATION.Seconds() - totalSleeptimeDuration.Seconds()

								fmt.Printf("Lap #%d finished, next lap in %d seconds, total laps=%d.\n", lap, int(nextLapSleepTime), config.LAPS)

								time.Sleep(time.Duration(int(nextLapSleepTime)) * time.Second)
							} else {

								fmt.Printf("Lap #%d finished, next lap in %d seconds, total laps=%d.\n", lap, 0, config.LAPS)

							}
						} else {
							//if last lap:
							fmt.Printf("Race #%d finished, next race in %d seconds.\n", raceId, int(config.RACE_TIMEOUT_DURATION.Seconds()))
							time.Sleep(config.RACE_TIMEOUT_DURATION)
							//increment raceId
							raceId++
						}
					}
				}
			/////
			}
		}
	}
}

func riders(config Config.Settings) (riders []string) {
	for rider:=1; rider <= config.RIDERS; rider++ {
		name := namegenerator.NewNameGenerator(time.Now().UTC().UnixNano()).Generate()
		riders = append(riders, name)
	}
	return
}

func RandBool75() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(4) != 1
}

func RandBool60() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) != 1
}

