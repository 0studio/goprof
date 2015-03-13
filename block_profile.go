package goprof

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"
)

var heapProfileCounter int32
var startTime = time.Now()

func DoBlockPrifile() {
	StartBlockProfile(1)
	defer StopBlockProfile()
	time.Sleep(10 * time.Second)
	fmt.Println("find do block profile")
	logInfo("find do block profile")
}
func StartBlockProfile(rate int) {
	runtime.SetBlockProfileRate(rate)
}

func StopBlockProfile() {
	filename := "block-" + strconv.Itoa(os.Getpid()) + ".pprof"
	f, err := os.Create(filename)
	if err != nil {
		logError(err)
	}
	if err = pprof.Lookup("block").WriteTo(f, 0); err != nil {
		logError(" can't write ", filename, err)
	}
	f.Close()
}
