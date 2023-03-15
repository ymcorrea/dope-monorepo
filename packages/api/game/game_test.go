package game

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/dopedao/dope-monorepo/packages/api/game/dopemap"
	"github.com/dopedao/dope-monorepo/packages/api/game/events"
	"github.com/dopedao/dope-monorepo/packages/api/game/item"
	"github.com/dopedao/dope-monorepo/packages/api/game/messages"
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

func TestDispatchPlayerJoin(t *testing.T) {
	assert := assert.New(t)

	gChan := make(chan messages.BroadcastMessage, 1)

	p1 := player.Player{
		Id: uuid.New(),
	}

	p2 := player.Player{
		Id: uuid.New(),
	}

	players := make([]*player.Player, 2)
	players[0] = &p1
	players[1] = &p2

	g := Game{
		Players:   players,
		Broadcast: gChan,
	}

	g.DispatchPlayerJoin(context.TODO(), &p2)

	out1 := <-gChan
	var joinData player.PlayerData
	json.Unmarshal(out1.Message.Data, &joinData)

	assert.Equal(events.PLAYER_JOIN, out1.Message.Event)
	assert.Equal(p2.Id.String(), joinData.Id)

	// dont send msg to p2
	assert.True(!out1.Condition(&p2))
}

func TestDispatchPlayerLeave(t *testing.T) {
	assert := assert.New(t)

	g := Game{
		Broadcast: make(chan messages.BroadcastMessage, 1),
	}

	p1 := player.Player{
		Conn:      &websocket.Conn{},
		HustlerId: uuid.NewString(),
		Id:        uuid.New(),
		Broadcast: g.Broadcast,
		Send:      make(chan messages.BaseMessage, 1),
	}

	players := make([]*player.Player, 1)
	players[0] = &p1

	g.Players = players

	g.DispatchPlayerLeave(context.TODO(), players[0])

	gOut := <-g.Broadcast
	var p1LeaveData item.IdData
	json.Unmarshal(gOut.Message.Data, &p1LeaveData)

	assert.Equal(events.PLAYER_LEAVE, gOut.Message.Event)
	assert.Equal(p1.Id.String(), p1LeaveData.Id)
}

func TestGenerateItemEntitiesData(t *testing.T) {
	assert := assert.New(t)

	item1 := item.ItemEntity{
		Id: uuid.New(),
		Item: item.Item{
			Item: "Powder",
		},
		Position: dopemap.Position{
			X: 10,
			Y: 15,
		},
	}

	g := Game{
		ItemEntities: []*item.ItemEntity{&item1},
	}

	itemEntData := g.GenerateItemEntitiesData()

	assert.Equal(item1.Id.String(), itemEntData[0].Id)
	assert.Equal(item1.Item.Item, itemEntData[0].Item)
	assert.Equal(item1.Position.X, itemEntData[0].X)
	assert.Equal(item1.Position.Y, itemEntData[0].Y)
}
