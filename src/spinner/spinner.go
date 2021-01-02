package spinner

import (
	"fmt"
	"time"
)

func RotateSpinner() {
	var spinner = [2]string{"♡", "♥"}
	var count int
	for {
		fmt.Printf("\r  \033[36mWaiting next lesson ...\033[m %s ", spinner[count%2])
		time.Sleep(time.Millisecond * 300)
		count++
	}
}
