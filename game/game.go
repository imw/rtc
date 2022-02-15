package game

import (
	"errors"
	"fmt"
	"net/rpc"

	"0x539.lol/rtc/board"
	"0x539.lol/rtc/util"
)

//Game encapsulates data about a game session
type Game struct {
	board  *board.Board
	peerID string
	client *rpc.Client
}

//New returns a new Game struct and GameRPC client
func New(b *board.Board) (*Game, *RPC) {
	g := &Game{
		board: b,
	}
	gr := &RPC{
		game: g,
	}
	return g, gr
}

//SetClient updates the rpc client for a game
func (g *Game) SetClient(c *rpc.Client) {
	g.client = c
}

//RPC is an RPC wrapper for the Game structure
type RPC struct {
	game *Game
}

//SendMove sends a move to a remote peer
func (g *Game) SendMove(m board.Move) error {
	var exit int
	err := g.client.Call("GameRPC.DoMove", m, &exit)
	if exit != 0 || err != nil {
		util.Write(fmt.Sprintf("Error sending move %v: %s", m, err))
		return errors.New("RPC call failed")
	}
	return err
}

//Peered returns whether a game session has an active peer
func (g *Game) Peered() bool {
	return g.peerID != ""
}

//DoMove executes a move against a local game board
func (gr *RPC) DoMove(m board.Move, exit *int) error {
	util.Write(fmt.Sprintf("Local move: %v", m))
	loc := gr.game.board.Position(m.Loc)
	tgt := gr.game.board.Position(m.Tgt)
	util.Write(fmt.Sprintf("moving %v to %v", loc, tgt))
	p := loc.Occupant()
	util.Write(fmt.Sprintf("address of p: %v", &p))
	tgt.SetOccupant(&p)
	if loc != tgt {
		loc.SetOccupant(nil)
	}
	p.Move()
	util.Write(fmt.Sprintf("after move: %v to %v", loc, tgt))
	// if p side is this (local) side, set cursor loc to p loc
	if gr.game.board.Side() == p.Side() {
		gr.game.board.SetLoc(tgt)
	}
	return nil
}

//Register updates the peerID for a game session
func (gr *RPC) Register(clientID string, exit *int) error {
	gr.game.peerID = clientID
	return nil
}
