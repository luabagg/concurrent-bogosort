package concurrent_bogosort

import (
	"context"
	"errors"
	"log"
	"sync"

	"golang.org/x/exp/rand"
)

var (
	wg            sync.WaitGroup
	m             sync.Mutex
	r             *rand.Rand
	ctx           context.Context
	cancel        context.CancelFunc
	goroutinesQty int64
)

// Sort uses a fast sorting technique called bogosort.
//
// Bogosort shuffles the slice randomly and verify if the it is sorted.
// This code opens a new goroutine until one matches a sorted array.
func Sort(slice []int) ([]int, error) {
	sliceLen := len(slice)
	if sliceLen >= 10 {
		return nil, errors.New("weeew why so many numbers")
	}
	log.Printf("Sorting slice of len: %d...\n", sliceLen)

	// If is sorted, returns itself.
	if isSortedAsc(slice) {
		return slice, nil
	}

	sliceChan := make(chan []int, sliceLen)
	ctx, cancel = context.WithCancel(context.Background())
	r = rand.New(new(rand.LockedSource))

	wg.Add(1)
out:
	for {
		// Checks if the context has been canceled.
		select {
		case <-ctx.Done():
			wg.Done()
			break out
		default:
		}

		go func(sliceChan chan []int, slice []int) {
			m.Lock()
			goroutinesQty += 1
			m.Unlock()

			shuffled := shuffle(slice)
			if !isSortedAsc(shuffled) {
				return
			}
			// When the slice is sorted, calls the cancel func to exit the loop.
			sliceChan <- shuffled
			cancel()
		}(sliceChan, slice)
	}
	wg.Wait()

	log.Printf("Goroutines opened: %d\n", goroutinesQty)

	return <-sliceChan, nil
}

// shuffle returns a slice of random positions.
func shuffle(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	r.Shuffle(len(slice), func(i int, j int) {
		newSlice[i], newSlice[j] = newSlice[j], newSlice[i]
	})

	return newSlice
}

// isSortedAsc checks if the slice is sorted in ascending order.
func isSortedAsc[T int](slice []T) bool {
	sliceLen := len(slice)
	if sliceLen == 0 || sliceLen == 1 {
		return true
	}
	for i := 0; i < sliceLen-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}

	return true
}
