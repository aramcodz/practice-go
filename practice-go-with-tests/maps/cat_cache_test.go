package main

/*
The Question: Build a Concurrent CacheProblem Statement:

In Go, standard maps do not support concurrent writes.
If two goroutines try to write to the same map at the same time, the program will crash with a fatal error:
*/

import (
	"math/rand/v2"
	"strconv"
	"testing"
)

func TestSafeMapInt(t *testing.T) {

	safeMapTests := []struct {
		name string
		key  string
		val  CatProfile
	}{
		{name: "Jasper", key: "Jasper-1",
			val: CatProfile{Id: 1, Name: "Jasper", Age: 15, Pattern: "tuxedo"},
		},
		{name: "Luna", key: "Luna-2",
			val: CatProfile{Id: 2, Name: "Luna", Age: 9, Pattern: "solid"},
		},
		{name: "Willa", key: "Willa-3",
			val: CatProfile{Id: 3, Name: "Willa", Age: 1, Pattern: "bicolor"},
		},
	}

	sm := SafeCatProfileMap{
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

	// sm := catProfileMap{
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

func createRandomCatProfile(id int) CatProfile {
	name := "furry-" + strconv.Itoa(id)
	ages := []int{1, 3, 5, 8, 15}
	// rand.N generates a random number from 0 to len(items)-1
	randIndex1 := rand.N(len(ages))
	age := ages[randIndex1]

	coatPatterns := []string{"solid", "bicolor", "tabby", "calico", "tortoiseshell"}
	randIndex2 := rand.N(len(coatPatterns))
	pattern := coatPatterns[randIndex2]

	return CatProfile{
		Id:      id,
		Name:    name,
		Age:     age,
		Pattern: pattern,
	}

}
