package atomgo

import (
	"fmt"
	"math/rand"
	"time"
)

func SmartDelay(interval float64, demo bool) float64 {
	rand.Seed(time.Now().UnixNano())
	min := interval * 4 / 5
	max := interval * 6 / 5
	r := min + rand.Float64()*(max-min)
	if demo {
		fmt.Printf("SmartDelay: %v seconds...", r)
	} else {
		time.Sleep(time.Duration(r) * time.Second)
	}
	return r
}
