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
	var y float32 = 20

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

	assert.Equal(x, bot.LastPosition.X)
	assert.Equal(y, bot.LastPosition.Y)
}

func TestPlayerByHustlerID(t *testing.T) {
	assert := assert.New(t)

	players := make([]*player.Player, 2)
	players[0] = &player.Player{HustlerId: uuid.NewString()}
	players[1] = &player.Player{HustlerId: uuid.NewString()}

	g := Game{
		Players: players,
	}

	out := g.PlayerByHustlerID(players[1].HustlerId)

	assert.Equal(out, players[1])
}

func TestPlayerByHustlerID_NotFound(t *testing.T) {
	assert := assert.New(t)

	players := make([]*player.Player, 1)
	players[0] = &player.Player{HustlerId: uuid.NewString()}

	g := Game{
		Players: players,
	}

	out := g.PlayerByHustlerID(uuid.NewString())

	assert.Nil(out)
}

func TestPlayerByUUID(t *testing.T) {
	assert := assert.New(t)

	players := make([]*player.Player, 2)
	players[0] = &player.Player{Id: uuid.New()}
	players[1] = &player.Player{Id: uuid.New()}

	g := Game{
		Players: players,
	}

	out := g.PlayerByUUID(players[1].Id)

	assert.Equal(out, players[1])
}

func TestPlayerByUUID_NotFound(t *testing.T) {
	assert := assert.New(t)

	players := make([]*player.Player, 2)
	players[0] = &player.Player{Id: uuid.New()}
	players[1] = &player.Player{Id: uuid.New()}

	g := Game{
		Players: players,
	}

	out := g.PlayerByUUID(uuid.New())

	assert.Nil(out)
}
