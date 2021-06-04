package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Store struct {
	id   string
	data map[string]string
	mux  sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) String() string {
	return s.id
}

func (s *Store) Get(key string) string {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.data[key]
}

func (s *Store) Put(key string, value string) {
	// xxxx
	// xxx
	s.mux.Lock()
	defer s.mux.Unlock()

	s.data[key] = value

	// s.mux.Lock()
	// s.data[key] = value
	// s.mux.Unlock()

	// Lock
	// data[a] = b
	// data[c] = b
	// Unlock
	//
	// id = ....
	// xxxx
	//
	// Lock
	// data[id] =
	// Unlock
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}(i)
	// }

	s := NewStore()
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			s.Put(strconv.Itoa(i), "value")
		}(i)

		go func(i int) {
			fmt.Println(i, "=", s.Get(strconv.Itoa(i)))
		}(i)
	}
	wg.Wait()

	// for k, v := range s.data {
	// 	fmt.Println(k, "=", v)
	// }

	fmt.Println("End")
	// time.Sleep(1 * time.Second)
}
