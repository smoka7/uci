package uci

import (
	"fmt"
	"reflect"
)

// Options, for initializing the chess engine
type Options struct {
	EvalFile          string `uci:"evalFile"`          // default nn-13406b1dcbe0.nnue stockfish engine
	SyzygyPath        string `uci:"syzygyPath"`        // default <empty>
	Hash              uint   `uci:"hash"`              // hash size in MB 1<<33554432
	Move              uint   `uci:"move overhead"`     // default 10 min 0 max 5000
	MultiPV           uint   `uci:"multipv"`           // default 1 min 1 max 500
	Skill             uint   `uci:"skill level"`       // default 20 min 0 max 20
	Slow              uint   `uci:"slow mover"`        // default 100 min 10 max 1000
	SyzygyProbeDepth  uint   `uci:"syzygyProbeDepth"`  // spin default 1 min 1 max 100
	SyzygyProbeLimit  uint   `uci:"syzygyProbeLimit"`  // spin default 7 min 0 max 7
	Threads           uint   `uci:"threads"`           // spin default 1 min 1 max 512
	UCI_Elo           uint   `uci:"UCI_Elo"`           // spin default 1350 min 1350 max 2850
	Nodestime         uint   `uci:"nodestime"`         // spin default 0 min 0 max 10000
	Ponder            bool   `uci:"ponder"`            // whether the engine should ponder
	Syzygy50MoveRule  bool   `uci:"syzygy50MoveRule"`  // default true
	UCI_AnalyseMode   bool   `uci:"UCI_AnalyseMode"`   // check default false
	UCI_Chess960      bool   `uci:"UCI_Chess960"`      // check default false
	UCI_LimitStrength bool   `uci:"UCI_LimitStrength"` // check default false
	UCI_ShowWDL       bool   `uci:"UCI_ShowWDL"`       // check default false
	Use_NNUE          bool   `uci:"use NNUE"`          // check default true
}

func NewOptions() Options {
	return Options{
		EvalFile:          "nn-13406b1dcbe0.nnue",
		SyzygyPath:        "",
		Hash:              16,
		Move:              10,
		MultiPV:           1,
		Skill:             20,
		Slow:              100,
		SyzygyProbeDepth:  1,
		SyzygyProbeLimit:  7,
		Threads:           1,
		UCI_Elo:           1350,
		Nodestime:         0,
		Ponder:            false,
		Syzygy50MoveRule:  false,
		UCI_AnalyseMode:   false,
		UCI_Chess960:      false,
		UCI_LimitStrength: false,
		UCI_ShowWDL:       false,
		Use_NNUE:          true,
	}
}

// SetOptions sends setoption commands to the Engine
// for the values set in the Options record passed in
func (eng *Engine) SetOptions(opt Options) error {
	var err error
	optType := reflect.TypeOf(opt)
	optValues := reflect.ValueOf(opt)
	for i, field := range reflect.VisibleFields(optType) {
		name := field.Tag.Get("uci")
		value := optValues.Field(i)

		err = eng.SendOption(name, value)
		if err != nil {
			return err
		}
	}
	return err
}

// SendOption sends setoption command to the Engine
func (eng *Engine) SendOption(name string, value interface{}) error {
	_, err := eng.stdin.WriteString(fmt.Sprintf("setoption name %s value %v\n", name, value))
	if err != nil {
		return err
	}
	err = eng.stdin.Flush()
	return err
}
