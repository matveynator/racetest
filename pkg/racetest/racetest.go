package Racetest

import (
  "time"
	"fmt"
	"math/rand"
	"racetest/pkg/config"
	"github.com/goombaio/namegenerator"
)

func Run(config Config.Settings) () {
	//get rider names:
	var sleeptime int
	competitors := riders(config)

	for lap := 0; lap <= config.LAPS; lap++ {
		//shuffle riders:
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(competitors), func(i, j int) { competitors[i], competitors[j] = competitors[j], competitors[i] })

		for index, name := range competitors {
			for result := 1; result <= config.RESULTS; result++ {
				fmt.Println(name)
			}
			 if (index + 1) != len(competitors) {
				rand.Seed(time.Now().UnixNano())
				sleeptime = rand.Intn(10 - 1)
				fmt.Printf("next rider in %d seconds...  index=%d, total riders=%d\n", sleeptime, index, len(competitors))
				time.Sleep(time.Duration(sleeptime) * time.Second)
			}
		}
		if lap != config.LAPS {
			rand.Seed(time.Now().UnixNano())
			sleeptime = rand.Intn(10 - 1)
			fmt.Printf("next lap in %d seconds... lap#=%d, total laps=%d\n", sleeptime, lap, config.LAPS)
			time.Sleep(time.Duration(sleeptime) * time.Second)
		}
	}
}

func riders(config Config.Settings) (riders []string) {
	for rider:=1; rider <= config.COMPETITORS; rider++ {
		name := namegenerator.NewNameGenerator(time.Now().UTC().UnixNano()).Generate()
		riders = append(riders, name)
	}
	return
}
