package kgs

import "math"

const updateInterval = 100

type InitKGSOptions struct {
	SaveCounter         func(int)
	LastModifiedCounter int
}

func NewKGS(opt InitKGSOptions) *kgs {
	kgs := &kgs{
		counter:            opt.LastModifiedCounter,
		lastUpdatedCounter: opt.LastModifiedCounter,
		saveCounter:        opt.SaveCounter,
		seed:               make([]byte, 0, 62),
	}
	kgs.fillSeed()

	return kgs
}

type KGS interface {
	GetKey() string
}

type kgs struct {
	counter            int
	lastUpdatedCounter int
	seed               []byte
	saveCounter        func(int)
}

func (kgs *kgs) updateCounter() {
	kgs.counter++

	if kgs.counter >= kgs.lastUpdatedCounter+updateInterval {
		kgs.lastUpdatedCounter = kgs.counter
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

	for c := kgs.counter; c != 0; {
		key = append(key, kgs.seed[int(c)%62])
		c = int(math.Floor(float64(c) / 62))
	}

	kgs.updateCounter()
	return string(key)
}
