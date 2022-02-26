package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestEngine(t *testing.T) {
	configA := Config{
		HostA: "localhost",
		PortA: "1234",
		HostB: "localhost",
		PortB: "4321",
		ID:    "A",
		Debug: true,
	}

	configB := configA
	configB.ID = "B"

	ready := make(chan GameEngine)
	go doSetup(t, configA, ready)
	go doSetup(t, configB, ready)

	var engineA *GameEngine
	var engineB *GameEngine

	waiting := 2
	for {
		engine := <-ready
		waiting = waiting - 1
		if engine.ID == "A" {
			engineA = &engine
		}
		engineB = &engine
		if waiting == 0 {
			break
		}
	}

	defer engineA.listener.Close()
	defer engineB.listener.Close()

	actions := []struct {
		engine *GameEngine
		key    tcell.Key
	}{
		{engineA, tcell.KeyEnter},
		{engineA, tcell.KeyUp},
		{engineA, tcell.KeyUp},
		{engineA, tcell.KeyEnter},
		{engineB, tcell.KeyLeft},
		{engineB, tcell.KeyEnter},
		{engineB, tcell.KeyDown},
		{engineB, tcell.KeyDown},
	}

	for _, action := range actions {
		ev := generateKeyPress(action.key)
		t.Logf("Executing move %s on engine %s", ev.Name(), action.engine.ID)
		t.Logf("Got tcell ev %v", ev)
		move := action.engine.game.RecieveEvent(ev)
		t.Logf("Got move: %v", move)
		t.Logf("Game is %v", action.engine.game)
		t.Logf("Board is %v", action.engine.game.GetBoard())
		if move.Seq > 0 {
			action.engine.game.SendMove(move)
			var rpcExitCode int
			action.engine.rpc.DoMove(move, &rpcExitCode)
		}
	}

}

func doSetup(t *testing.T, config Config, ready chan<- GameEngine) {
	t.Logf("Setting up with config: %v", config)
	running, e := setup(config)
	if e != nil {
		t.Errorf("Error setting up engine. Config: %v, Error: %v", config, e)
	}
	t.Logf("Setup succeeded for %s", config.ID)
	ready <- running
}

func generateKeyPress(k tcell.Key) *tcell.EventKey {
	return tcell.NewEventKey(k, 0, tcell.ModNone)
}
