package main

/*
The Question: Build a Concurrent Cache

Problem Statement:
In Go, standard maps do not support concurrent writes.
Reference Dave Cheney Blog Post:
   https://dave.cheney.net/2015/12/07/are-go-maps-sensitive-to-data-races

If two goroutines try to write to the same map at the same time, the program will crash with a fatal error:

Concurrent Maps Issue References ->
  1. https://go.dev/blog/maps#concurrency
  2. https://go.dev/doc/faq#atomic_maps

TODO:
Let's build a CatProfile Cache

CatProfile(s) consist of a Cat ID, Name and Age.
*/

import (
	"sync"
)

type CatProfile struct {
	Id      int
	Name    string
	Age     int
	Pattern string
}

type SafeCatProfileMap struct {
	mu sync.Mutex
	m  map[string]CatProfile
}

//	Incremental Step for implemenation
//
// NOT safe for Concurrent Write Access
func (s *SafeCatProfileMap) Set(key string, value CatProfile) {
	s.m[key] = value
}

// NOT safe for Concurrent Read Access
func (s *SafeCatProfileMap) Get(key string) (CatProfile, bool) {
	return s.m[key], true
}

func (s *SafeCatProfileMap) Count() int {
	return len(s.m)
}
