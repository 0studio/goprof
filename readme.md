# Demo

```
goprof.Init(fmt.Println,fmt.Println)
} else if value.Type == defs.REDIS_QUEUE_TYPE_DEBUG_PRINT_STACK {
goprof.WriteGoroutineStack()
} else if value.Type == defs.REDIS_QUEUE_TYPE_DEBUG_PROF_HEAP {
goprof.DoHeapProfile()
} else if value.Type == defs.REDIS_QUEUE_TYPE_DEBUG_PROF_CPU {
goprof.DoCpuProfile()
} else if value.Type == defs.REDIS_QUEUE_TYPE_DEBUG_PROF_BLOCK {
goprof.DoBlockPrifile()
} 

```