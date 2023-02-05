# racetest

[Chicha timekeeper](https://github.com/matveynator/chicha) race testing tool for any platform.

### [-> download latest version <-](http://files.matveynator.ru/racetest/latest/)

```
racetest -h
Usage of /usr/local/bin/racetest:
  -laps int
    	Set ammount of laps each race holds. (default 5)
  -minimal-lap-time duration
    	Minimal lap time duration. Results smaller than this duration would be considered wrong. Valid time units are: 's' (second), 'm' (minute), 'h' (hour). (default 1m0s)
  -results int
    	Set ammount of results each competitor sends to timekeeper. (default 10)
  -riders int
    	Set ammount of riders (sportsmen) for the race. (default 10)
  -target string
    	Provide IP address and port where to send data (where chicha timekeeper is running). (default "127.0.0.1:4000")
  -timeout duration
    	Set race timeout duration. After this time if nobody passes the finish line the race will be stopped. Valid time units are: 's' (second), 'm' (minute), 'h' (hour). (default 3m0s)
  -version
    	Output version information
```
