package main

import "fmt"

func main()  {
	var fleet []interface{Ship}
	fleet = append(fleet, sub{}, destroyer{})

	fmt.Println(fleet)

	for _, ship := range fleet {
		ship = ship.(Ship)
		ship.Fire()
	}
}

type Ship interface {
	Fire()
}

type position struct {
	x int
	y int
}

type sub struct {
	pos position
}

type destroyer struct {
	pos position
}

func (s sub) Fire()  {
	fmt.Println("open silo doors")
}

func (d destroyer) Fire() {
	fmt.Println("Aim battery")
}


func newSub() *sub {
	s := sub{}
	return &s
}