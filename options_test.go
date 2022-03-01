package uci

import (
	"testing"
)

func TestEngine_SetOptions(t *testing.T) {
	opts := []Options{
		{
			Hash:    100,
			Move:    10,
			MultiPV: 12,
			Ponder:  true,
			Threads: 2,
		},
	}
	engine, err := NewEngine("stockfish")
	defer engine.Close()
	if err != nil {
		t.Log(err)
	}
	for _, opt := range opts {
		if err := engine.SetOptions(opt); err != nil {
			t.Fatal(err)
		}
		_, err := engine.stdin.WriteString("uci")
		if err != nil {
			t.Fatal(err)
		}
	}
}
