package game

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dopedao/dope-monorepo/packages/api/game/dopemap"
	"github.com/dopedao/dope-monorepo/packages/api/game/events"
	"github.com/dopedao/dope-monorepo/packages/api/game/item"
	"github.com/dopedao/dope-monorepo/packages/api/game/messages"
	"github.com/dopedao/dope-monorepo/packages/api/game/player"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/enttest"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/schema"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
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
	}

	item2 := item.ItemEntity{
		Id: uuid.New(),
	}

	g := Game{
		ItemEntities: []*item.ItemEntity{&item1, &item2},
	}

	itemEntData := g.GenerateItemEntitiesData()

	assert.Equal(item1.Id.String(), itemEntData[0].Id)
	assert.Equal(item2.Id.String(), itemEntData[1].Id)
}

func TestGeneratePlayersData(t *testing.T) {
	assert := assert.New(t)

	p1 := player.Player{
		Id: uuid.New(),
	}

	p2 := player.Player{
		Id: uuid.New(),
	}

	g := Game{
		Players: []*player.Player{&p1, &p2},
	}

	playerData := g.GeneratePlayersData()

	assert.Equal(p1.Id.String(), playerData[0].Id)
	assert.Equal(p2.Id.String(), playerData[1].Id)
}

/*
func TestRegisterPlayer(t *testing.T) {
	assert := assert.New(t)

	p := player.Player{
		Conn: &websocket.Conn{},
		Id:   uuid.New(),
		Name: "DJ Screw",
		Send: make(chan messages.BaseMessage, 1),
	}

	g := Game{}

	g.registerPlayer(&p, context.TODO(), &ent.Client{})

	pOut := <-p.Send

	var handshakeData HandshakeData
	json.Unmarshal(pOut.Data, &handshakeData)

	assert.Equal(events.PLAYER_HANDSHAKE, pOut.Event)
}
*/

func TestUnregisterPlayer_RemovesPlayer(t *testing.T) {
	assert := assert.New(t)

	p := player.Player{}

	g := NewGame()
	g.Players = append(g.Players, &p)

	g.unregisterPlayer(&p, context.TODO(), &ent.Client{}, &zerolog.Logger{})

	assert.Len(g.Players, 0)
}

func TestUnregisterPlayer_PanicIfCantGetHustler(t *testing.T) {
	assert := assert.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	p := player.Player{
		HustlerId: uuid.NewString(),
		Position: dopemap.Position{
			X: 10,
			Y: 10,
		},
		CurrentMap: "memphis",
	}

	g := NewGame()
	g.Players = append(g.Players, &p)

	assert.Panics(func() {
		g.unregisterPlayer(&p, context.TODO(), &ent.Client{}, &zerolog.Logger{})
	})
}

func TestUnregisterPlayer_SaveLastPos(t *testing.T) {
	assert := assert.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	p := player.Player{
		HustlerId: uuid.NewString(),
		Position: dopemap.Position{
			X: 20,
			Y: 20,
		},
		CurrentMap: "memphis",
	}

	_, err := client.GameHustler.Create().SetID(p.HustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	g := NewGame()
	g.Players = append(g.Players, &p)

	g.unregisterPlayer(&p, context.TODO(), client, &zerolog.Logger{})

	dbHustler, err := client.GameHustler.Get(context.TODO(), p.HustlerId)
	if err != nil {
		assert.Fail(err.Error())
	}

	assert.Equal(p.Position.X, dbHustler.LastPosition.X)
	assert.Equal(p.Position.Y, dbHustler.LastPosition.Y)
	assert.Equal(p.CurrentMap, dbHustler.LastPosition.CurrentMap)
}

func TestBroadcast_TestCondition(t *testing.T) {
	assert := assert.New(t)

	g := NewGame()
	p := player.Player{
		Send: make(chan messages.BaseMessage, 1),
	}

	g.Players = append(g.Players, &p)

	testMsg := messages.BroadcastMessage{
		Message: messages.BaseMessage{
			Event: events.Event("test"),
		},
	}

	//Dont send to bot
	g.broadcast(testMsg)
	assert.Len(p.Send, 0)

	//Send to self
	p.Conn = &websocket.Conn{}
	g.broadcast(testMsg)
	assert.Len(p.Send, 1)

	// clear chan
	<-p.Send

	//Send if not self
	testMsg.Condition = func(i interface{}) bool {
		ptr := i.(*player.Player)

		return i == &ptr
	}

	g.broadcast(testMsg)
	assert.Len(p.Send, 0)
}

func TestNewGame(t *testing.T) {
	assert := assert.New(t)

	g := NewGame()

	assert.NotNil(g.Ticker)
	assert.NotNil(g.SpawnPosition)
	assert.NotNil(g.Register)
	assert.NotNil(g.Unregister)
	assert.NotNil(g.Broadcast)
}

func TestGetNearPlayersMoveData(t *testing.T) {
	assert := assert.New(t)

	g := NewGame()

	p1 := player.Player{
		Id:         uuid.New(),
		CurrentMap: "test",
		Position: dopemap.Position{
			X: 10,
			Y: 10,
		},
	}

	p2 := player.Player{
		Id:         uuid.New(),
		CurrentMap: "test",
		Position: dopemap.Position{
			X: 11,
			Y: 11,
		},
		LastPosition: dopemap.Position{
			X: 6,
			Y: 6,
		},
		Direction: "e",
	}

	g.Players = []*player.Player{&p1, &p2}

	moveData := *g.getNearPlayersMoveData(&p1)

	assert.Len(moveData, 1)
	assert.Equal(p2.Id.String(), moveData[0].Id)
	assert.Equal(p2.Position.X, moveData[0].X)
	assert.Equal(p2.Position.Y, moveData[0].Y)
	assert.Equal(p2.Direction, moveData[0].Direction)
}

func TestTick(t *testing.T) {
	assert := assert.New(t)

	g := NewGame()

	p1 := player.Player{
		Conn: &websocket.Conn{},
		Send: make(chan messages.BaseMessage, 1),
	}

	p2 := player.Player{
		Conn: &websocket.Conn{},
		Send: make(chan messages.BaseMessage, 1),
	}

	g.Players = []*player.Player{&p1, &p2}
	g.tick(context.TODO(), time.Time{})

	p1Out := <-p1.Send
	p2Out := <-p2.Send

	assert.Equal(events.TICK, p1Out.Event)
	assert.Equal(events.TICK, p2Out.Event)
}

/*
func TestHandlePlayerJoin(t *testing.T) {
	assert := assert.New(t)

	g := NewGame()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	g.HandlePlayerJoin(context.TODO(), &websocket.Conn{}, client, &ent.GameHustler{})

	gout := <-g.Register

	assert.NotNil(gout)
}
*/

func WssFactory() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, _ := upgrader.Upgrade(w, r, nil)

		// Echo messages back to the client
		for {
			var msg messages.BaseMessage
			conn.ReadJSON(&msg)

			conn.WriteJSON(msg)
		}
	}))
}
