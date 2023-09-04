package bitmap

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestBitmap(t *testing.T) {
	bm := New(10)
	bm.Set(9)
	bm.Set(10)
	bm.Set(1)
	assert.Equal(t, bm.Exits(9), true)
	assert.Equal(t, bm.Exits(10), true)
	assert.Equal(t, bm.Exits(1), true)
	bm.Clean(9)
	assert.Equal(t, bm.Exits(9), false)

}
