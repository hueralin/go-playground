package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *safeCounter) Inc(key string) {
	// 不加锁会报错
	// fatal error: concurrent map writes
	// 在 Go 语言中，map 不是线程安全的数据结构
	s.mu.Lock()
	s.m[key]++
	s.mu.Unlock()
}

func (s *safeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[key]
}

func ShowMutex() {
	sc := safeCounter{m: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go sc.Inc("xxx")
	}
	time.Sleep(time.Second)
	fmt.Println(sc.Value("xxx"))
}
