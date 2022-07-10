// KGS stands for Key Generator Service
package kgs

import (
	"math"
	"sync"

	"github.com/mreza0100/shortly/internal/ports"
)

// kgs dependency
type KGSDep struct {
	SaveCounter      func(int64)
	LastSavedCounter int64
}

// Get New instance of KGS
func New(dep *KGSDep) ports.KGS {
	// Counter <= 0 not allowed
	if dep.LastSavedCounter <= 0 {
		dep.LastSavedCounter = 1
	}
	if dep.LastSavedCounter != 1 {
		dep.LastSavedCounter += 10
	}

	kgs := &kgs{
		counter:     dep.LastSavedCounter,
		saveCounter: dep.SaveCounter,
		mu:          new(sync.Mutex),
		seed:        make([]byte, 0, 62),
	}
	kgs.fillSeed()

	return kgs
}

// kgs implementation
type kgs struct {
	saveCounter func(int64)
	counter     int64
	seed        []byte
	mu          *sync.Mutex
}

// Update Counter - must be called after generating a new key
// this function is NOT thread-safe. concurrency is handled by the caller
// TODO: add thread-safety and handle concurrency here
func (kgs *kgs) updateCounter() {
	kgs.counter++

	// TODO: a system to determine when to save the counter based on the number of keys generated and requests per second
	// save when the counters first digit is 0
	if kgs.counter%10 == 0 {
		kgs.saveCounter(kgs.counter)
	}
}

func (kgs *kgs) fillSeed() {
	// seed is a byte array of 62 characters (0-9a-zA-Z)
	for _, r := range [][2]byte{
		{'0', '9'}, // numbers
		{'a', 'z'}, // lowercase
		{'A', 'Z'}, // uppercase
	} {
		from, to := r[0], r[1]

		for i := from; i <= to; i++ {
			kgs.seed = append(kgs.seed, i)
		}
	}
}

// converting desired number to base62
func (kgs *kgs) convertToBase62(c int64) string {
	key := make([]byte, 0, 10)
	for c := float64(kgs.counter); c != 0; {
		key = append(key, kgs.seed[int(c)%62])
		c = math.Floor(c / 62)
	}
	return string(key)
}

// GenerateKey generates a new key
func (kgs *kgs) GetKey() string {
	kgs.mu.Lock()

	key := kgs.convertToBase62(kgs.counter)

	kgs.updateCounter()
	// not using defer for performance reasons
	// reference: https://medium.com/i0exception/runtime-overhead-of-using-defer-in-go-7140d5c40e32
	kgs.mu.Unlock()
	return string(key)
}
