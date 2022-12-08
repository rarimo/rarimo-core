package crypto

import (
	"testing"
)

func TestGetThreshold(t *testing.T) {
	cases := []int{3, 4, 5, 6, 7, 8}

	for _, n := range cases {
		t.Logf("N = %d, T = %d", n, GetThreshold(n))
	}
}
