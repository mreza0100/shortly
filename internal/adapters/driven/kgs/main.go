package kgs

import (
	"math"
)

type InitKGSOptions struct {
	SaveCounter         func(int64)
	LastModifiedCounter int64
}

func New(opt InitKGSOptions) *kgs {
	kgs := &kgs{
		counter:     opt.LastModifiedCounter,
		saveCounter: opt.SaveCounter,
		seed:        make([]byte, 0, 62),
	}
	kgs.fillSeed()

	return kgs
}

type KGS interface {
	GetKey() string
}

type kgs struct {
	counter     int64
	seed        []byte
	saveCounter func(int64)
}

func (kgs *kgs) updateCounter() {
	kgs.counter++

	// TODO: a system to determine when to save the counter based on the number of keys generated and requests/secent
	if kgs.counter%10 == 0 {
		kgs.saveCounter(kgs.counter)
	}
}

func (kgs *kgs) fillSeed() {
	for _, range_ := range [][2]byte{
		{'0', '9'}, // numbers
		{'a', 'z'}, // lowercase
		{'A', 'Z'}, // uppercase
	} {
		from, to := range_[0], range_[1]

		for i := from; i <= to; i++ {
			kgs.seed = append(kgs.seed, i)
		}
	}
}

func (kgs *kgs) GetKey() string {
	key := make([]byte, 0, 10)

	for c := float64(kgs.counter); c != 0; {
		key = append(key, kgs.seed[int(c)%62])
		c = math.Floor(c / 62)
	}

	kgs.updateCounter()
	return string(key)
}
