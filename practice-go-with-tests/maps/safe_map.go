package main

/*
The Question: Build a Concurrent Cache

Problem Statement:
In Go, standard maps do not support concurrent writes.
Reference Dave Cheney Blog Post:
   https://dave.cheney.net/2015/12/07/are-go-maps-sensitive-to-data-races

If two goroutines try to write to the same map at the same time, the program will crash with a fatal error:

TODO: create a SafeMap implemenetation.
Implement a thread-safe Cache struct that allows multiple goroutines to safely read and write data concurrently without crashing.

*/

import (
	"sync"
)

type SafeIntMap struct {
	mu sync.Mutex
	m  map[string]int
}

//	Incremental Step for implemenation
//
// NOT safe for Concurrent Write Access
func (s *SafeIntMap) Set(key string, value int) {
	s.m[key] = value
}

// NOT safe for Concurrent Read Access
func (s *SafeIntMap) Get(key string) (int, bool) {
	return s.m[key], true
}
