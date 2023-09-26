package factorial_bogosort

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var (
	wg            sync.WaitGroup
	m             sync.Mutex
	ctx           context.Context
	cancel        context.CancelFunc
	randomOrders  []int64
	goroutinesQty int64
)

// Sort sorts the array.
// Each possible order opens a new goroutine that randomly shuffles the array and verifies if it's ordered.
func Sort(slice []int) ([]int, error) {
	sliceLen := len(slice)

	if sliceLen > 10 {
		return nil, errors.New("weeew why so much numbers")
	}

	maxOrdering, err := recursiveFactorial(sliceLen)
	if err != nil {
		return nil, err
	}

	log.Printf("Max possible ordering: %d\n", maxOrdering)

	sliceChan := make(chan []int, sliceLen)

	ctx, cancel = context.WithCancel(context.Background())

	wg.Add(1)
	for i := int64(0); i < maxOrdering; i++ {
		go shuffle(sliceChan, slice[:], i, ctx)
	}
	wg.Wait()

	log.Printf("Goroutines opened: %d\n", goroutinesQty)

	slice = <-sliceChan

	return slice, nil
}

// recursiveFactorial calculates the number factorial.
func recursiveFactorial(number int) (int64, error) {
	switch {
	case number < 1:
		return 0, errors.New("number can only be 0, 1 or greater")
	case number == 0 || number == 1:
		return 1, nil
	default:
		result, err := recursiveFactorial(number - 1) // rand numbers generated are >= 1
		if err != nil {
			return 0, err
		}

		return (int64(number) * result), nil
	}
}

// shuffle get a slice of random positions and verify if the slice is sorted.
func shuffle(sliceChan chan []int, slice []int, it int64, ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	m.Lock()
	goroutinesQty += 1
	m.Unlock()

	sliceLen := len(slice)

	newSlice := make([]int, sliceLen)
	positions := getRandomPositions(slice)

	for i := 0; i < sliceLen; i++ {
		newSlice[i] = slice[int(positions[i]-1)]
	}

	if sort.IntsAreSorted(newSlice) {
		sliceChan <- newSlice
		cancel()
		wg.Done()
	}

	return
}

// getRandomPositions returns a slice of random positions.
// The returned positions are >= 1 (1, 2, 3...).
func getRandomPositions(slice []int) []int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sliceLen := len(slice)

	positions := make([]int64, sliceLen)
	for i := 0; i < sliceLen; i++ {
		randNumber := int64(r.Intn(sliceLen)) + 1
		if intInSlice(randNumber, positions) {
			i--
			continue
		}
		positions[i] = randNumber
	}

	positionInteger := uniqueId(positions)

	m.Lock()
	if intInSlice(positionInteger, randomOrders) {
		m.Unlock()
		return getRandomPositions(slice)
	} else {
		randomOrders = append(randomOrders, positionInteger)
		m.Unlock()
	}

	return positions
}

// uniqueId generates a unique identifier for each sequence.
// Example: [4, 1, 2] -> 412
func uniqueId(slice []int64) int64 {
	var res int64

	op := int64(1)
	for i := len(slice) - 1; i >= 0; i-- {
		res += int64(slice[i]) * op
		op *= 10
	}

	return res
}

// intInSlice verifies if the integer is in the given slice.
func intInSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
