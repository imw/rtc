package game

import (
	"errors"
	"fmt"
	"net/rpc"

	"0x539.lol/rtc/board"
	"0x539.lol/rtc/util"
)

type Game struct {
	board  *board.Board
	peerID string
	client *rpc.Client
}

func New(b *board.Board) (*Game, *GameRPC) {
	g := &Game{
		board: b,
	}
	gr := &GameRPC{
		game: g,
	}
	return g, gr
}

func (g *Game) SetClient(c *rpc.Client) {
	g.client = c
}

type GameRPC struct {
	game *Game
}

func (g *Game) SendMove(m board.Move) error {
	var exit int
	err := g.client.Call("GameRPC.DoMove", m, &exit)
	if exit != 0 || err != nil {
		util.Write(fmt.Sprintf("Error sending move %v: %s", m, err))
		return errors.New("RPC call failed")
	}
	return err
}

func (g *Game) Peered() bool {
	return g.peerID != ""
}

func (gr *GameRPC) DoMove(m board.Move, exit *int) error {
	util.Write(fmt.Sprintf("Local move: %v", m))
	loc := gr.game.board.Position(m.Loc)
	tgt := gr.game.board.Position(m.Tgt)
	util.Write(fmt.Sprintf("moving %v to %v", loc, tgt))
	p := loc.Occupant()
	util.Write(fmt.Sprintf("address of p: %v", &p))
	tgt.SetOccupant(&p)
	loc.SetOccupant(nil)
	p.Move()
	util.Write(fmt.Sprintf("after move: %v to %v", loc, tgt))
	// if p side is this (local) side, set cursor loc to p loc
	if gr.game.board.Side() == p.Side() {
		gr.game.board.SetLoc(tgt)
	}
	return nil
}

func (gr *GameRPC) Register(clientID string, exit *int) error {
	gr.game.peerID = clientID
	return nil
}
