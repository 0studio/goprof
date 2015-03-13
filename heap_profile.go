package goprof

import (
	"github.com/0studio/goutils"
	"os"
	"runtime/debug"
)

// go get github.com/randall77/hprof/dumptodot
//  dumptodot heapdump mybinary > heap.dot
// dot file.dot -Tpng -o image.png

//  go get github.com/randall77/hprof/dumptohprof
//  dumptohprof heapdump heap.hprof
//  jhat heap.hprof
// and navigate your browser to http://localhost:7000.
func WriteHeapDump() {
	f, err := os.Create("runtime-heap-" + goutils.Int2Str(os.Getpid()) + ".heap")
	if err != nil {
		logError("get goroutine heapdump error", err)
	}

	debug.WriteHeapDump(f.Fd())
	if f != nil {
		f.Close()
		f = nil
	}
}
