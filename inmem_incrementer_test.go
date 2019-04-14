package measures

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemIncrementer_Get(t *testing.T) {
	ri := NewMemIncrementer(nil)
	a, err := ri.Get("nonexistent111")
	assert.NoError(t, err)
	assert.Equal(t, 0.0, a)
}

func TestMemIncrementer_IncrAndGet(t *testing.T) {
	ri := NewMemIncrementer(nil)
	tkey := "tkey_" + time.Now().String()
	a, err := ri.Get(tkey)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, a)

	a, err = ri.IncAndGet(tkey, 1.1)
	assert.NoError(t, err)
	assert.Equal(t, 1.1, a)

	a, err = ri.Get(tkey)
	assert.NoError(t, err)
	assert.Equal(t, 1.1, a)

}
