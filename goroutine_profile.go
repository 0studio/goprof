package goprof

import (
	"fmt"
	"github.com/0studio/goutils"
	"os"
	"runtime"
	"runtime/pprof"
)

func PrintStack() {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Printf("%s", buf)
}
func WriteGoroutineStack() {
	f, err := os.Create("goroutine-" + goutils.Int2Str(os.Getpid()) + ".pprof")
	if err != nil {
		logError("get goroutine stack error", err)
	}
	pprof.Lookup("goroutine").WriteTo(f, 1)

}
