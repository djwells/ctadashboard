package cta

import (
	"time"
)

type Route struct {
	Id    string
	Name  string
	CtaId int
}

type TrainLine struct {
	Id     string
	Name   string
	CtaId  string
	Routes []Route
}

type Station struct {
	Id         int
	Name       string
	RouteStops []Route
}

type ArrivalInfo struct {
	StationId   int
	StationName string
	Timestamp   time.Time
}

var redRoutes = [...]Route{
	Route{"123", "Howard", 123},
	Route{"123", "95th / Dan Ryan", 123},
}

var blueRoutes = [...]Route{
	Route{"123", "O'Hare", 123},
	Route{"123", "Forest Park", 123},
}

var TrainLines = [...]TrainLine{
//TrainLine{"A", "Red", "R", redRoutes},
//TrainLine{"A", "Blue", "B", blueRoutes},
}
