package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// mem := flag.Int("m", 50, "-m 50M")
	all := flag.String("a", "n", "all memory (-a y)")
	flag.Parse()
	m := make([]int32, 0)
	if *all == "y" {
		fmt.Println("Eat all your memory")
		for {
			m = append(m, 1)
		}
	} else {
		// fmt.Printf("Eat %dM memory\n", *mem)
		// for i := 1; i < *mem*1000*100; i++ {
		// 	m = append(m, 1)
		// }
	}
	time.Sleep(time.Hour * 24)
}
