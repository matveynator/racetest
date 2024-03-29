package Racetest

import (
	"time"
	"fmt"
	"math/rand"
	"strings"
	"racetest/pkg/config"
	"racetest/pkg/data"
	"racetest/pkg/proxy"
)

var rawData Data.RawData

func Run(config Config.Settings) () {

	var sleeptime int = 0 
	var totalSleeptime int = 0
	var competitors []string

	averageSleepSecondsInt  := int(config.MINIMAL_LAP_TIME_DURATION.Seconds())/config.RIDERS
	if averageSleepSecondsInt == 0 {
		averageSleepSecondsInt = 1
	}

	raceId := 1

	if config.RANDOM == true {
		fmt.Println("Preparing tests with random data...")
	}

	for {
		//for each new race:

		//get new rider list for each race:
		competitors = riders(config)
		//fmt.Printf("created list with %d riders for race #%d \n", len(competitors), raceId)

		for lap := 0; lap <= config.LAPS; lap++ {

			if config.RANDOM {
				//shuffle riders:
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(competitors), func(i, j int) { competitors[i], competitors[j] = competitors[j], competitors[i] })
			}

			//fmt.Printf("shuffled list with %d riders for race #%d \n", len(competitors), raceId)

			//zero total sleep time:
			totalSleeptime = 0

			for index, name := range competitors {

				if config.RANDOM {
					//fmt.Printf("rider %s and raceid %d \n", name, raceId)
					//not each competitor finishes each lap (90%):
					if RandBool90() {
						for result := 1; result <= config.RESULTS; result++ {
							//not each result registered by antenna (60%):
							//fmt.Printf("rider %s and result %d \n", name, result)
							if RandBool60() {
								//fmt.Printf("rider %s and random result %d \n", name, result)
								time.Sleep(time.Duration(20) * time.Millisecond)
								rawData.TagId = name
								rawData.DiscoveryUnixTime = time.Now().UnixNano()/int64(time.Millisecond)

								rand.Seed(time.Now().UnixNano())
								rawData.Antenna = uint(rand.Intn(3))

								Proxy.ProxyTask <- rawData

								fmt.Printf("%s,%d,%d\n", rawData.TagId, rawData.DiscoveryUnixTime, rawData.Antenna )
							}
						}

						//if not last rider in the lap: 
						if index != (len(competitors)-1) {

							//get random sleeptime = 1 second:
							rand.Seed(time.Now().UnixNano())
							sleeptime = rand.Intn(averageSleepSecondsInt)

							//count total sleeptime
							totalSleeptime = (totalSleeptime + sleeptime)

							fmt.Printf("next rider in %d seconds...  index=%d, total riders=%d\n", sleeptime, index, len(competitors))

							//sleep:
							if sleeptime > 0 {
								time.Sleep(time.Duration(sleeptime) * time.Second)
							}
						} else {
							//if last rider in the lap and not the last lap:
							if lap != config.LAPS {
								//if last rider in the lap:
								//calculate time left for the lap (to sleep after all riders checked in):
								if time.Duration(totalSleeptime) * time.Second < config.MINIMAL_LAP_TIME_DURATION {
									totalSleeptimeDuration := time.Duration(totalSleeptime) * time.Second
									nextLapSleepTime := config.MINIMAL_LAP_TIME_DURATION.Seconds() - totalSleeptimeDuration.Seconds() 

									fmt.Printf("Lap - #%d finished, next lap in %d seconds, total laps=%d.\n", lap, int(nextLapSleepTime), config.LAPS)

									time.Sleep(time.Duration(int(nextLapSleepTime)) * time.Second)
								} else {

									fmt.Printf("Lap -- #%d finished, next lap in %d seconds, total laps=%d.\n", lap, 0, config.LAPS)

								}
							}
						}
					}
				} else {
					//if selected not random execution:

					for result := 1; result <= config.RESULTS; result++ {
						//print result:
						time.Sleep(time.Duration(20) * time.Millisecond)
						rawData.TagId = name
						rawData.DiscoveryUnixTime = time.Now().UnixNano()/int64(time.Millisecond)

						rand.Seed(time.Now().UnixNano())
						rawData.Antenna = uint(rand.Intn(3))
						Proxy.ProxyTask <- rawData
						fmt.Printf("%s,%d,%d\n", rawData.TagId, rawData.DiscoveryUnixTime, rawData.Antenna )

					}

					//if not last rider in the lap:
					if index != (len(competitors)-1) {

						//get random sleeptime = 1 second:
						rand.Seed(time.Now().UnixNano())
						sleeptime = rand.Intn(averageSleepSecondsInt)

						//count total sleeptime
						totalSleeptime = (totalSleeptime + sleeptime)

						fmt.Printf("next rider in %d seconds...  index=%d, total riders=%d\n", sleeptime, index, len(competitors))

						//sleep:
						time.Sleep(time.Duration(sleeptime) * time.Second)
					} else {
						//if last rider in the lap and not the last lap:
						if lap != config.LAPS {
							//if last rider in the lap:
							//calculate time left for the lap (to sleep after all riders checked in):
							if time.Duration(totalSleeptime) * time.Second < config.MINIMAL_LAP_TIME_DURATION {
								totalSleeptimeDuration := time.Duration(totalSleeptime) * time.Second
								nextLapSleepTime := config.MINIMAL_LAP_TIME_DURATION.Seconds() - totalSleeptimeDuration.Seconds()

								fmt.Printf("Lap --- #%d finished, next lap in %d seconds, total laps=%d.\n", lap, int(nextLapSleepTime), config.LAPS)

								time.Sleep(time.Duration(int(nextLapSleepTime)) * time.Second)
							} else {

								fmt.Printf("Lap ---- #%d finished, next lap in %d seconds, total laps=%d.\n", lap, 0, config.LAPS)

							}
						}
					}
				}
			}
		}

		//if last lap:
		fmt.Printf("Race #%d finished, next race in %d seconds.\n", raceId, int(config.RACE_TIMEOUT_DURATION.Seconds()))
		time.Sleep(config.RACE_TIMEOUT_DURATION)
		//increment raceId
		raceId++
	}





}

func generateRandomNameAndSurname() string {
	rand.Seed(time.Now().UnixNano())

	// Список гласных и согласных букв
	vowels := "aeiouy"
	consonants := "bcdfghjklmnpqrstvwxz"

	// Генерируем имя
	firstNameLen := rand.Intn(3) + 10 // Длина имени от 2 до 12 букв
	firstName := ""
	for i := 0; i < firstNameLen; i++ {
		if i == 0 {
			// Первая буква имени должна быть заглавной
			firstName += strings.ToUpper(string(consonants[rand.Intn(len(consonants))]))
		} else {
			// Остальные буквы имени должны быть строчными
			firstName += string(vowels[rand.Intn(len(vowels))])
		}
	}

	// Генерируем фамилию
	lastNameLen := rand.Intn(3) + 12 // Длина фамилии от 2 до 12 букв
	lastName := ""
	for i := 0; i < lastNameLen; i++ {
		if i == 0 {
			// Первая буква фамилии должна быть заглавной
			lastName += strings.ToUpper(string(consonants[rand.Intn(len(consonants))]))
		} else {
			// Остальные буквы фамилии должны быть строчными
			lastName += string(vowels[rand.Intn(len(vowels))])
		}
	}

	// Соединяем имя и фамилию в одну строку
	return firstName + lastName
}

func getRandomDriverName() string {

	drivers := []string{
		"VladimirChagin",
		"HansStacey",
		"GerardDeRooy",
		"EduardNikolaev",
		"AiratMardeev",
		"DmitrySotnikov",
		"AntonShibalov",
		"KarelLoprais",
		"JanDeRooy",
		"TobyPrice",
		"RickyBrabec",
		"CyrilDespres",
		"KevinBenavides",
		"NailHubbatulin",
		"ToniBou",
		"ValentinoRossi",
		"AirtonSenna",
		"MichaelSchumacher",
		"NikkiLauda",
		"LewisHamilton",
		"NigelMansell",
		"AlainProst",
		"CarlosSainz",
		"BillElliott",
		"MarcMarquez",
		"AdamRaga",
		"ChuckNorris",
	}
	rand.Seed(time.Now().UnixNano())
	return drivers[rand.Intn(len(drivers))]
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func riders(config Config.Settings) (riders []string) {
	var name string

	// Sleep to calm down startup prompts:
	fmt.Println("Sleeping 3 seconds to calm down startup prompts...")
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("Preparing %d random riders for the race: ", config.RIDERS)

	for {

		name = generateRandomNameAndSurname()
		if !contains(riders, name) {
			riders = append(riders, name)
		}
		for portionOfResults := 10; portionOfResults <= 100; portionOfResults=portionOfResults+10 {
			if len(riders) == (config.RIDERS / 100) * portionOfResults {
				fmt.Printf(" %d%%", portionOfResults)
			}
		}
		if len(riders) == config.RIDERS {
			fmt.Println(" DONE.")
			return
		}
	}
}

func RandBool90() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9) != 1
}

func RandBool60() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) != 1
}
