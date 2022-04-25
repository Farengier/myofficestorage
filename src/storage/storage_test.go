package storage

import (
	"testing"
)

func storages() map[string]Storage {
	return map[string]Storage{
		"mem": MemoryStorage(),
	}
}

func runStorageTests(t *testing.T, tf func(s Storage, t *testing.T)) {
	for k, s := range storages() {
		t.Run(k, func(t *testing.T) {
			tf(s, t)
		})
	}
}

func TestEmpty(t *testing.T) {
	tf := func(s Storage, t *testing.T) {
		_, err := s.Get("some_key")
		if err == nil {
			t.Fatal("Get on newly created memory storage should not found values")
		}
	}
	runStorageTests(t, tf)
}

func FuzzStoring(f *testing.F) {
	var d []byte
	for i := 0; i <= 255; i++ {
		d = append(d, byte(i))
		f.Add([]byte{byte(i)})
	}
	f.Add(d)
	f.Fuzz(func(t *testing.T, value []byte) {
		tf := func(s Storage, t *testing.T) {
			err := s.Store("some_key", value)
			if err != nil {
				t.Fatal("Error while storing value")
			}

			val, err := s.Get("some_key")
			if err != nil {
				t.Fatal("Error getting previously stored value")
			}
			if len(value) != len(val) {
				t.Fatalf("Values have different sizes %d initial and %d was stored", len(value), len(val))
			}
			for i, vi := range value {
				if vs := val[i]; vi != vs {
					t.Fatalf("Data error: was %d, got %d", vi, vs)
				}
			}
		}
		runStorageTests(t, tf)
	})
}

func TestDeletion(t *testing.T) {
	tf := func(s Storage, t *testing.T) {
		value := []byte("some value")

		err := s.Delete("some_key")
		if err != nil {
			t.Fatal("Deleting should not return error even if key does not exist")
		}

		err = s.Store("some_key", value)
		if err != nil {
			t.Fatal("Error while storing value")
		}

		_, err = s.Get("some_key")
		if err != nil {
			t.Fatal("Error getting previously stored value")
		}

		err = s.Delete("some_key")
		if err != nil {
			t.Fatal("Error deleting previously stored value")
		}

		_, err = s.Get("some_key")
		if err == nil {
			t.Fatal("Value was deleted, should return error")
		}
	}
	runStorageTests(t, tf)
}
