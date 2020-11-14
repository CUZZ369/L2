package main

import (
	"fmt"
	"time"
)

func main()  {
	// func Now() Time
	fmt.Println(time.Now())
	// func Parse(layout, value string) (Time, error)
	time.Parse("2016-01-02 15:04:05", "2020-11-4 21:24:51")
	// func ParseInLocation(layout, value string, loc *Location) (Time, error) (layout已带时区时可直接用Parse)
	time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)
	// func Unix(sec int64, nsec int64) Time
	time.Unix(1e9, 0)
	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	time.Date(2018, 1, 2, 15, 30, 10, 0, time.Local)
	//func (t Time) In(loc *Location) Time 当前时间对应指定时区的时间
	loc, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Now().In(loc))
	// func (t Time) Local() Time

}
