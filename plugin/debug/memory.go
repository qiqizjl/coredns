package debug

import (
	"runtime"
	"time"

	"github.com/coredns/coredns/plugin/pkg/log"
)

// reportMemory reports the memory used by the current process by logging to standard error every second.
func reportMemory(stop chan struct{}) error {
	go func() {
		tick := time.NewTicker(1 * time.Second)
		var m runtime.MemStats
		for {
			select {
			case <-tick.C:
				runtime.ReadMemStats(&m)
				log.Debugf("alloc: %10.4f MiB, total alloc: %10.4f MiB, system: %10.4f MiB",
					float32(m.Alloc)/1024/1024, float32(m.TotalAlloc)/1024/1024, float32(m.Sys)/1024/1024)
			case <-stop:
				return
			}
		}
	}()

	return nil
}
