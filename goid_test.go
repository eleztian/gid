package gls

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
	_ "unsafe"
)

func TestGoID(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := GoID()
			if !isThisGIDInStackTrace(id) {
				t.Error("goroutine id error ", id)
			}
		}()
	}
	wg.Wait()
}

func TestWithID(t *testing.T) {
	var id int64
	wg := sync.WaitGroup{}
	wg.Add(1)
	go WithGoID(&id, func() {
		defer wg.Done()
		id := GoID()
		if !isThisGIDInStackTrace(id) {
			t.Error("goroutine id error ", id)
		}
	})
	wg.Wait()
}

func isThisGIDInStackTrace(id int64) bool {
	lines := strings.Split(stackTrace(), "\n")
	for i, line := range lines {
		if !strings.HasPrefix(line, fmt.Sprintf("goroutine %d ", id)) {
			continue
		}
		if i+1 == len(lines) {
			break
		}
		if strings.Contains(lines[i+1], ".stackTrace") {
			return true
		}
	}
	return false
}

func stackTrace() string {
	var n int
	for n = 4096; n < 16777216; n *= 2 {
		buf := make([]byte, n)
		ret := runtime.Stack(buf, true)
		if ret != n {
			return string(buf[:ret])
		}
	}
	panic(n)
}
