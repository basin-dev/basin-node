package adapters

import (
	"bytes"
	"testing"
)

func FuzzLocalAdapter(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
		f.Add(seed)
	}
	_, err := StartDB()
	if err != nil {
		f.Errorf("Failed to start LevelDB: %s", err.Error())
	}

	// Check that you get the same thing out that you put in
	f.Fuzz(func(t *testing.T, expect []byte) {
		err := LocalAdapter.Write("test", expect)
		if err != nil {
			t.Errorf("Failed to write to LocalAdapter: %s", err.Error())
		}
		val, err := LocalAdapter.Read("test")
		if err != nil {
			t.Errorf("Failed to read from LocalAdapter: %s", err.Error())
		}
		if bytes.Compare(val, expect) != 0 {
			t.Error("Wrong value returned")
		}
	})
}

func BenchmarkLocalAdapter(b *testing.B) {
	StartDB()
	b.ResetTimer()
	b.Run("Write/Read benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LocalAdapter.Write("test", []byte("test"))
			LocalAdapter.Read("test")
		}
	})
}
