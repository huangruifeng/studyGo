package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Printf("current time %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println(year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("timestamp1 is %v\n", timestamp1)
	fmt.Printf("timestamp1 is %v\n", timestamp2)

	timeobj := time.Unix(timestamp1, timestamp2-timestamp1)
	fmt.Println(timeobj)

	year := timeobj.Year()
	month := timeobj.Month()
	day := timeobj.Day()
	hour := timeobj.Hour()
	minute := timeobj.Minute()
	second := timeobj.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	fmt.Println(12 * time.Hour)

}

func timeOperation() {
	now := time.Now()
	later := now.Add(time.Hour * 2)
	fmt.Println(now)
	fmt.Println(later)

	fmt.Println(later.Sub(now))

	fmt.Println(later.Equal(now.Add(2 * time.Hour)))

	fmt.Println(later.Before(now), later.After(now))
}

func formatDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("2006-01-02 3:04:05 PM Jan Min"))
	fmt.Println(now.Format("15:04 2006/01/02"))

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	timeObj, err := time.ParseInLocation("15:04 2006/01/02", "13:38 2020/11/14", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(now.Sub(timeObj))

}

func ticDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

func main() {
	formatDemo()
	//ticDemo()
	timeOperation()
	timestampDemo()
	timeDemo()
}
