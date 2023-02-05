package Proxy

import (
	"log"
	"fmt"
	"net"
	"time"

	"racetest/pkg/mylog"
	"racetest/pkg/data"
	"racetest/pkg/config"
)
//оставляем только один процесс который будет брать задачи и передавать их далее на другой сервер (чтобы сохранялась последовательность)
var proxyWorkersMaxCount int = 1
var ProxyTask chan Data.RawData
var respawnLock chan int

func init() {
	//initialise channel with tasks:
	ProxyTask = make(chan Data.RawData)

	//initialize unblocking channel to guard respawn tasks
	respawnLock = make(chan int, proxyWorkersMaxCount)
}
func Run(config Config.Settings) {

	if config.PROXY_ADDRESS != "" {

		log.Println("connecting to 0", config.PROXY_ADDRESS)

		go func() {
			for {
				// will block if there is proxyWorkersMaxCount ints in respawnLock
				respawnLock <- 1 
				//sleep 1 second
				time.Sleep(1 * time.Second)
				go proxyWorkerRun(len(respawnLock), config)
			}
		}()
	}
}


//close connection on programm exit
func deferCleanup(connection net.Conn) {
	<-respawnLock
	if connection != nil {
		err := connection.Close() 
		if err != nil {
			log.Println("Error closing proxy connection:", err)
		}
	}
}

func proxyWorkerRun(workerId int, config Config.Settings) {

	connection, err := net.DialTimeout("tcp", config.PROXY_ADDRESS, 15 * time.Second)
	defer deferCleanup(connection)
	if err != nil  {
		MyLog.Printonce(fmt.Sprintf("Proxy destination %s unreachable. Error: %s",  config.PROXY_ADDRESS, err))
		return

	} else {
		MyLog.Println(fmt.Sprintf("Proxy worker #%d connected to destination %s", workerId, config.PROXY_ADDRESS))
	}
	//initialise connection error channel
	connectionErrorChannel := make(chan error, 1)
	go func() {
		buffer := make([]byte, 1024)
		for {
			numberOfLines, err := connection.Read(buffer)
			if err != nil {
				connectionErrorChannel <- err
				return
			}
			if numberOfLines > 0 {
				log.Printf("Proxy worker received unexpected data back: %s", buffer[:numberOfLines])
			}
		}
	}()

	for {
		select {
			//в случае если есть задание в канале ProxyTask
		case currentProxyTask := <- ProxyTask :
			_, networkSendingError := fmt.Fprintf(connection, "%s, %d, %d\n", currentProxyTask.TagId, currentProxyTask.DiscoveryUnixTime, currentProxyTask.Antenna)
			if err != nil {
				//в случае потери связи во время отправки мы возвращаем задачу обратно в канал ProxyTask
				ProxyTask <- currentProxyTask
				log.Printf("Proxy worker %d exited due to sending error: %s\n", workerId, networkSendingError)
				//и завершаем работу гоурутины
				return
			}
		case networkError := <-connectionErrorChannel :
			//обнаружена сетевая ошибка - завершаем гоурутину
			log.Printf("Proxy worker %d exited due to connection error: %s\n", workerId, networkError)
			return
		}
	}
}

