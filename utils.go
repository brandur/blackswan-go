package main

import (
	"fmt"
	"time"
)

var (
	appName string
)

func init() {
	appName = "black-swan-api"
}

func MeasureT(t time.Time, name string) {
	name = appName + "." + name
	fmt.Printf("measure=%q val=%v\n", name, time.Since(t))
}
