package kgs

import (
	"math"
	"sync"

	"github.com/mreza0100/shortly/internal/ports"
)

type InitKGSOptions struct {
	SaveCounter      func(int64)
	LastSavedCounter int64
}

func New(opt *InitKGSOptions) ports.KGS {
	kgs := &kgs{
		counter:     opt.LastSavedCounter,
		saveCounter: opt.SaveCounter,
		mu:          new(sync.Mutex),
		seed:        make([]byte, 0, 62),
	}
	kgs.fillSeed()

	return kgs
}

type kgs struct {
	saveCounter func(int64)
	counter     int64
	seed        []byte
	mu          *sync.Mutex
}

func (kgs *kgs) updateCounter() {
	kgs.counter++

	// TODO: a system to determine when to save the counter based on the number of keys generated and requests/secent
	if kgs.counter%10 == 0 {
		go kgs.saveCounter(kgs.counter)
	}
}

func (kgs *kgs) fillSeed() {
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

func (kgs *kgs) convertToBase62(c int64) string {
	key := make([]byte, 0, 10)
	for c := float64(kgs.counter); c != 0; {
		key = append(key, kgs.seed[int(c)%62])
		c = math.Floor(c / 62)
	}
	return string(key)
}

func (kgs *kgs) GetKey() string {
	kgs.mu.Lock()

	key := kgs.convertToBase62(kgs.counter)

	kgs.updateCounter()
	// not using defer for performance reasons
	// reference: https://medium.com/i0exception/runtime-overhead-of-using-defer-in-go-7140d5c40e32
	kgs.mu.Unlock()
	return string(key)
}
