package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		res, _ := strconv.Atoi(string(key))
		return uint32(res)
	})

	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"16": "6",
		"23": "4",
		"26": "6",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	hash.Add("8")

	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yield %s, yielded %s", k, v, hash.Get(k))
		}
	}
}
