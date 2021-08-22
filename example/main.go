package main

import (
	"fmt"
	"time"

	timer "github.com/span42/go-timer"
)

func main() {
	lpTimer1 := timer.NewTimer(lpFunc, "lpTimer1")
	defer lpTimer1.Destory()
	lpTimer1.Start(2000) //msec

	lpTimer2 := timer.NewTimer(lpFunc, "lpTimer2")
	defer lpTimer2.Destory()
	lpTimer2.Start(2000)
	time.Sleep(time.Second)
	lpTimer2.Stop() //stop at any time

	select {}
}

//timeout callback
func lpFunc(t *timer.Timer, lpArg ...interface{}) {
	t.Start(2000)
	fmt.Println(lpArg)
}
