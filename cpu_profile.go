package goprof

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

func DoCpuProfile() {
	StartCPUProfile()
	defer StopCPUProfile()
	time.Sleep(30 * time.Second)
	fmt.Println("Finish cpu profile. Please check dump file.")
}

func StartCPUProfile() {
	f, err := os.Create("cpu-" + strconv.Itoa(os.Getpid()) + ".pprof")
	if err != nil {
		logError(err)
	}
	pprof.StartCPUProfile(f)
}

func StopCPUProfile() {
	pprof.StopCPUProfile()
}
