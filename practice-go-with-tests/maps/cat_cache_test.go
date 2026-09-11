package main

/*
The Question: Build a Concurrent CacheProblem Statement:

In Go, standard maps do not support concurrent writes.
If two goroutines try to write to the same map at the same time, the program will crash with a fatal error:

*/

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
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

	catProfiles := SafeCatProfileMap{
		m: make(map[string]CatProfile),
	}

	workers := 500
	iterations := 500
	var wg sync.WaitGroup
	// 1. Concurrent Writes - launch # of workers goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				// create the test data --> CatProfiles map, and adds to the "SafeMap"
				key := strconv.Itoa(i)
				cp := createRandomCatProfile(i)
				catProfiles.Set(key, cp)
			}

		}(i)
	}

	time.Sleep(time.Second)
	println("map count: " + strconv.Itoa(catProfiles.Count()))

	// 2. Concurrent Reads
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			key := strconv.Itoa(i)
			actualProfile, ok := catProfiles.Get(key)
			if !ok {
				t.Errorf("Could Not Get profile ID %v", i)
			}
			//Val should Exist - Name must start with "furry-"
			if !strings.HasPrefix(actualProfile.Name, "furry-") {
				t.Errorf("The actual Profile Name %v does Not match the expected %v", actualProfile.Name, "furry-")
			}
			if actualProfile.Id != i {
				t.Errorf("The actual Profile ID %v does Not match the expected %v", actualProfile.Id, i)
			}

			//fmt.Printf("Cat with ID %v is Named: %s \n", i, actualProfile.Name)
		}(1)
	}

	wg.Wait()

}

func createRandomCatProfile(id int) CatProfile {
	nameSuffixes := []string{"Luna", "Mittens", "Willa", "Jasper", "Floof", "Benny", "Benjamin", "Thomas", "KittenCorn", "Princess-P", "Willa-Pilla", "Buna-Cat"}
	// rand.N generates a random number from 0 to len(items)-1
	randIdx1 := rand.N(len(nameSuffixes))

	name := "furry-" + nameSuffixes[randIdx1]
	ages := []int{1, 3, 5, 8, 15}

	randIdx2 := rand.N(len(ages))
	age := ages[randIdx2]

	coatPatterns := []string{"solid", "bicolor", "tabby", "calico", "tortoise", "sealpoint", "tuxedo"}
	randIdx3 := rand.N(len(coatPatterns))
	pattern := coatPatterns[randIdx3]

	return CatProfile{
		Id:      id,
		Name:    name,
		Age:     age,
		Pattern: pattern,
	}
}
