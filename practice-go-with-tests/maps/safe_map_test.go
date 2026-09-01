package main

/*
The Question: Build a Concurrent CacheProblem Statement:

In Go, standard maps do not support concurrent writes.
If two goroutines try to write to the same map at the same time, the program will crash with a fatal error:

create a SafeMap implemenetation.
Implement a thread-safe Cache struct that allows multiple goroutines to safely read and write data concurrently without crashing.

*/

import "testing"

func TestSafeMapInt(t *testing.T) {
	safeMapTests := []struct {
		name string
		key  string
		val  int
	}{
		{name: "A entry", key: "A", val: 3},
		{name: "B entry", key: "B", val: 14},
		{name: "C entry", key: "B", val: 42},
	}

	sm := SafeIntMap{
		m: make(map[string]int),
	}

	for _, tt := range safeMapTests {
		sm.Set(tt.key, tt.val)
		t.Run(tt.name, func(t *testing.T) {
			actualVal, ok := sm.Get(tt.key)
			if !ok {
				t.Fatalf("map-Get returned err")
			}
			if tt.val != actualVal {
				t.Errorf("The actual map value %d does Not match the expected %d", actualVal, tt.val)
			}
		})
	}
}

//func TestSafeMapConcurrent()
