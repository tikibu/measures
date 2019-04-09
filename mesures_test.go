package measures

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

type FaultyInc struct {
	t    map[string]float64
	pret float64
}

func NewFaultyInc(p float64, startwith map[string]float64) (fi *FaultyInc) {
	fi = &FaultyInc{}
	fi.pret = p
	if startwith != nil {
		fi.t = startwith
	} else {
		fi.t = map[string]float64{}
	}
	return
}

func (fi *FaultyInc) IncAndGet(key string, by float64) (res float64, err error) {
	if rand.Float64() < fi.pret {
		return 0.0, errors.New("err")
	} else {
		fi.t[key] += by
		return fi.t[key], nil
	}
}

func (fi *FaultyInc) Get(key string) (res float64, err error) {
	res, ok := fi.t[key]
	if !ok {
		res = 0.0
	}
	return
}

func TestNewFloatMesure(t *testing.T) {
	rand.Seed(42)
	fi := NewFaultyInc(0.1, map[string]float64{"t": 2.})
	m, err := NewFloatMesure("t", fi, time.Millisecond*5)
	assert.NoError(t, err)
	for i := 1; i < 1000; i++ {
		m.Inc(1.0)
		r := m.Get()
		assert.Equal(t, float64(i)+2., r)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Millisecond * 200)
	r, _ := fi.Get("t")
	assert.Equal(t, float64(1001), r)

}
