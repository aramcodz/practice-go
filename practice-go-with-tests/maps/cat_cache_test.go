package main

/*
The Question: Build a Concurrent CacheProblem Statement:

In Go, standard maps do not support concurrent writes.
If two goroutines try to write to the same map at the same time, the program will crash with a fatal error:
*/

import (
	"testing"
)

func TestSafeMapInt(t *testing.T) {

	safeMapTests := []struct {
		name string
		key  string
		val  CatProfile
	}{
		{name: "Jasper", key: "Jasper-1",
			val: CatProfile{Id: 1, Name: "Jasper", Age: 15},
		},
		{name: "Luna", key: "Luna-2",
			val: CatProfile{Id: 2, Name: "Luna", Age: 9},
		},
		{name: "Willa", key: "Willa-3",
			val: CatProfile{Id: 3, Name: "Willa", Age: 1},
		},
	}

	sm := SafeIntMap{
		m: make(map[string]CatProfile),
	}

	for _, tt := range safeMapTests {
		sm.Set(tt.key, tt.val)
		t.Run(tt.name, func(t *testing.T) {
			actualVal, ok := sm.Get(tt.key)
			if !ok {
				t.Fatalf("map-Get returned err")
			}
			println("actual Cat: %v", actualVal.Name)
			if tt.val != actualVal {
				t.Errorf("The actual map value %v does Not match the expected %v", actualVal.Name, tt.val.Name)
			}
		})
	}
}

func TestSafeMapConcurrent(t *testing.T) {

	//TODO:

	// sm := SafeIntMap{
	// 	m: make(map[string]int),
	// }
	// var wg sync.WaitGroup
	// workers := 500

	// Concurrent Reads and Writes - launch # of workers goroutines
	// for i := 0; i < workers; i++ {
	// 	wg.Add(1)
	// 	go func(workerId int) {
	// 		sm.Set()
	// 	}

	// }

}
