package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ticket struct {
	Seat string
	Start_time time.Time 
}

func main()  {
	t := ticket{
		Seat : "4D",
		Start_time: time.Now(),
	} 

	var jsonData []byte
	jsonData, err := json.Marshal(t)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(jsonData)

}