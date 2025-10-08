package weather

import "time"

type LocationInfo struct {
	City    string
	Region  string
	Country string
	Time    time.Duration
}

type WeatherLocationInfo struct {
	TempC      int16
	FeelslikeC int16 //Ощущается как
	Humidity   int8  //Влажность
	Condition  string
	WindKph    int16 //Скорость ветра км/ч
	Location   LocationInfo
}
