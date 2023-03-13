package game

import (
	"testing"

	"github.com/dopedao/dope-monorepo/packages/api/game/dopemap"
	"github.com/dopedao/dope-monorepo/packages/api/game/player"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestAddBots(t *testing.T) {
	assert := assert.New(t)

	g := Game{}

	g.AddBots(5)

	assert.Equal(5, len(g.Players))
}

func TestPlayerByConn(t *testing.T) {
	assert := assert.New(t)

	g := Game{}
	p := player.Player{
		Id:   uuid.New(),
		Conn: &websocket.Conn{},
	}

	g.Players = append(g.Players, &p)

	res := g.PlayerByConn(p.Conn)

	assert.Equal(p.Conn, res.Conn)
	assert.Equal(p.Id, res.Id)
}

func TestUpdateBotPosition(t *testing.T) {
	assert := assert.New(t)

	var x float32 = 10
	var y float32 = 10

	bot := player.Player{
		Conn: nil,
		Position: dopemap.Position{
			X: x,
			Y: y,
		},
	}

	players := make([]*player.Player, 1)
	players[0] = &bot

	UpdateBotPosition(&players)

	assert.NotEqual(x, bot.Position.X)
	assert.NotEqual(y, bot.Position.Y)
}
