package ordenador

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{4, 1, 3, 2}
	sliceExpected := []int{1, 2, 3, 4}

	output := SliceOrdenador(nums)

	assert.Equal(t, sliceExpected, output)
}
