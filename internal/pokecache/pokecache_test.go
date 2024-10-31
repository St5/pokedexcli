package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nill")
	}
}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("{some test}"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("{some test2}"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("{some test3}"),
		},
	}

	for _, cs := range cases {

		cache.Add(cs.inputKey, cs.inputVal)
		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Error("Some error happend")
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("Return not equal: %v vs %v", string(cs.inputVal), string(actual))
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("{some data}"))

	time.Sleep(interval + time.Millisecond)

	_ , ok := cache.Get(keyOne)
	if ok {
		t.Error("The value still present")
	}
	
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("{some data}"))

	time.Sleep(interval/2)

	_ , ok := cache.Get(keyOne)
	if !ok {
		t.Error("The value was disapeared")
	}
	
}

