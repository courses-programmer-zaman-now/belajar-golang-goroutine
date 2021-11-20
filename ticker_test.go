package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	var counter = 0

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time, " peringatan ke - ", counter)
		counter++
	}

	// for {
	// 	select {
	// 	case time := <-ticker.C:
	// 		fmt.Println(time, " peringatan ke - ", counter)
	// 		counter++

	// 	}
	// 	if counter == 5 {
	// 		break
	// 	}
	// }

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
