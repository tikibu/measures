package measures

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMeasureStore_AlreadyExistValue(t *testing.T) {
	storageMap := map[string]float64{
		"existent111": 10.0,
	}
	ms := CreateMeasureStore(NewMemIncrementer(storageMap), time.Second)
	a, err := ms.Get("existent111")
	assert.NoError(t, err)
	assert.Equal(t, 10.0, a)
}
