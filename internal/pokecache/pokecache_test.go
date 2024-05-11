package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key: %v", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value: %v", c.val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	key := "https://example.com"
	val := []byte("testdata")
	cache.Add(key, val)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key: %v", key)
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected not to find key: %v", key)
	}
}
