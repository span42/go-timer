go-timer
======== 
A very simple and easy-to-use timer

Details
-------


Usage
-----
​```go
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
	select {}
}

//timeout callback
func lpFunc(t *timer.Timer, lpArg ...interface{}) {
	t.Start(3000)
	fmt.Println(lpArg)
}
​```

Possible Future Work
-------------------- 
