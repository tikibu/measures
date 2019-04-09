package measures

import (
	"sync"
	"time"
)

type MeasureStore struct {
	incrementer Incrementer
	measures    map[string]*FloatMeasure
	mu          sync.Mutex
	synctime    time.Duration
}

func CreateMeasureStore(inc Incrementer, synctime time.Duration) *MeasureStore {
	ms := &MeasureStore{incrementer: inc, measures: map[string]*FloatMeasure{}, synctime: synctime}
	return ms
}

func (ms *MeasureStore) Get(key string) (v float64, err error) {
	ms.mu.Lock()
	measure, ok := ms.measures[key]
	if !ok {
		fm, err := NewFloatMesure(key, ms.incrementer, ms.synctime)
		if err != nil {
			return 0.0, err
		}
		ms.measures[key], err = fm, err
		ms.mu.Unlock()
		return 0.0, nil
	} else {
		ms.mu.Unlock()
		return measure.Get(), nil
	}
}

func (ms *MeasureStore) Inc(key string, v float64) (err error) {
	ms.mu.Lock()
	measure, ok := ms.measures[key]
	if !ok {
		fm, err := NewFloatMesure(key, ms.incrementer, ms.synctime)
		if err != nil {
			return err
		}
		ms.measures[key], err = fm, err
		ms.mu.Unlock()
		fm.Inc(v)
		return nil
	} else {
		ms.mu.Unlock()
		measure.Inc(v)
		return nil
	}
}
