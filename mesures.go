package measures

import (
	"sync"
	"time"
)

type FloatMeasure struct {
	Key         string
	incrementer Incrementer

	mu sync.Mutex

	common    float64
	delta     float64
	syncDelta float64
}

func NewFloatMesure(key string, inc Incrementer, syncperiod time.Duration) (fm *FloatMeasure, err error) {
	fm = &FloatMeasure{Key: key, incrementer: inc}
	fm.common, err = inc.Get(key)
	if err != nil {
		return
	}
	go func() {
		for {
			fm.mu.Lock()
			fm.syncDelta = fm.delta
			fm.delta = 0.0
			fm.mu.Unlock()

			new_common, err := fm.incrementer.IncAndGet(fm.Key, fm.syncDelta)

			fm.mu.Lock()
			if err == nil {
				fm.common = new_common
				fm.syncDelta = 0.0 // this one just in case
			} else {
				fm.delta += fm.syncDelta
				fm.syncDelta = 0.0 //this one just in case
			}
			fm.mu.Unlock()
			time.Sleep(syncperiod)
		}
	}()

	return
}

func (m *FloatMeasure) Inc(v float64) {
	m.mu.Lock()
	m.delta += v
	m.mu.Unlock()
}

func (m *FloatMeasure) Get() (ret float64) {
	m.mu.Lock()
	ret = m.common + m.delta + m.syncDelta
	m.mu.Unlock()
	return
}
