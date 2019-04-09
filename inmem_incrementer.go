package measures

import "sync"

type MemIncrementer struct {
	t  map[string]float64
	mu sync.Mutex
}

func NewMemIncrementer(startwith map[string]float64) (fi *MemIncrementer) {
	fi = &MemIncrementer{}
	if startwith != nil {
		fi.t = startwith
	} else {
		fi.t = map[string]float64{}
	}
	return
}

func (fi *MemIncrementer) IncAndGet(key string, by float64) (res float64, err error) {
	fi.mu.Lock()
	defer fi.mu.Unlock()
	fi.t[key] += by
	return fi.t[key], nil
}

func (fi *MemIncrementer) Get(key string) (res float64, err error) {
	fi.mu.Lock()
	defer fi.mu.Unlock()
	res, ok := fi.t[key]
	if !ok {
		res = 0.0
	}
	return
}
