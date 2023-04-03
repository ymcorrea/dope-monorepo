package game

import (
	"context"
	"encoding/json"
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

func newPlayerMock(broadcastChan *chan messages.BroadcastMessage, isBot bool) *player.Player {
	p := player.Player{
		Id:        uuid.New(),
		Send:      make(chan messages.BaseMessage, 1),
		Broadcast: *broadcastChan,
		Position: dopemap.Position{
			X: 10,
			Y: 10,
		},
		LastPosition: dopemap.Position{
			X: 5,
			Y: 5,
		},
		Direction:  "NE",
		CurrentMap: "memphis",
	}

	if !isBot {
		p.HustlerId = uuid.NewString()
		p.Conn = &websocket.Conn{}
	}

	return &p
}

func newDbClientMock(t *testing.T) *ent.Client {
	return enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
}

func newItemMock() *item.ItemEntity {
	return &item.ItemEntity{
		Id: uuid.New(),
	}
}

func newGameMock(playerCount int, isBot bool, itemCount int) *Game {
	g := Game{
		Broadcast: make(chan messages.BroadcastMessage, 1),
	}

	for i := 0; i < playerCount; i++ {
		g.Players = append(g.Players, newPlayerMock(&g.Broadcast, isBot))
	}

	for i := 0; i < itemCount; i++ {
		g.ItemEntities = append(g.ItemEntities, newItemMock())
	}

	return &g
}

func TestAddBots(t *testing.T) {
	assert := assert.New(t)

	g := Game{}

	g.AddBots(5)

	assert.Equal(5, len(g.Players))
}

func TestPlayerByConn(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(2, false, 0)
	firstPlayer := g.Players[0]

	res := g.PlayerByConn(firstPlayer.Conn)

	assert.Equal(firstPlayer.Conn, res.Conn)
	assert.Equal(firstPlayer.Id, res.Id)
}

func TestUpdateBotPosition(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(1, true, 0)
	currPos := g.Players[0].Position

	g.UpdateBotPosition()
	lastPos := g.Players[0].LastPosition

	assert.Equal(currPos.X, lastPos.X)
	assert.Equal(currPos.Y, lastPos.Y)
}

func TestPlayerByHustlerID(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(2, false, 0)
	secPlayer := g.Players[1]

	out := g.PlayerByHustlerID(secPlayer.HustlerId)

	assert.Equal(out, secPlayer)
}

func TestPlayerByHustlerID_NotFound(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(2, false, 0)

	out := g.PlayerByHustlerID(uuid.NewString())

	assert.Nil(out)
}

func TestPlayerByUUID(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(2, true, 0)
	firstPlayer := g.Players[1]

	out := g.PlayerByUUID(firstPlayer.Id)

	assert.Equal(out, firstPlayer)
}

func TestPlayerByUUID_NotFound(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(2, true, 0)

	out := g.PlayerByUUID(uuid.New())

	assert.Nil(out)
}

func TestDispatchPlayerJoin(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(1, true, 0)
	newP := *newPlayerMock(&g.Broadcast, true)

	g.DispatchPlayerJoin(context.TODO(), &newP)

	out1 := <-g.Broadcast
	var joinData player.PlayerData
	json.Unmarshal(out1.Message.Data, &joinData)

	assert.Equal(events.PLAYER_JOIN, out1.Message.Event)
	assert.Equal(newP.Id.String(), joinData.Id)

	// dont send msg to p2
	assert.True(!out1.Condition(&newP))
}

func TestDispatchPlayerLeave(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(2, false, 0)
	p1 := g.Players[0]

	g.DispatchPlayerLeave(context.TODO(), p1)

	gOut := <-g.Broadcast
	var p1LeaveData item.IdData
	json.Unmarshal(gOut.Message.Data, &p1LeaveData)

	assert.Equal(events.PLAYER_LEAVE, gOut.Message.Event)
	assert.Equal(p1.Id.String(), p1LeaveData.Id)
}

func TestGenerateItemEntitiesData(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(0, true, 2)

	itemEntData := g.GenerateItemEntitiesData()

	assert.Equal(g.ItemEntities[0].Id.String(), itemEntData[0].Id)
	assert.Equal(g.ItemEntities[1].Id.String(), itemEntData[1].Id)
}

func TestGeneratePlayersData(t *testing.T) {
	assert := assert.New(t)

	g := newGameMock(2, false, 0)

	playerData := g.GeneratePlayersData()

	assert.Equal(g.Players[0].Id.String(), playerData[0].Id)
	assert.Equal(g.Players[1].Id.String(), playerData[1].Id)
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

	g := newGameMock(2, false, 0)

	assert.Panics(func() {
		g.unregisterPlayer(g.Players[0], context.TODO(), &ent.Client{}, &zerolog.Logger{})
	})
}

func TestUnregisterPlayer_SaveLastPos(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)

	g := newGameMock(1, false, 0)
	p := g.Players[0]

	_, err := client.GameHustler.Create().SetID(p.HustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	g.unregisterPlayer(p, context.TODO(), client, &zerolog.Logger{})

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

	g := newGameMock(2, false, 0)
	p1 := g.Players[0]
	p2 := g.Players[1]

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

func TestGenerateHandshakeData(t *testing.T) {
	assert := assert.New(t)

	p := player.Player{
		HustlerId: uuid.NewString(),
		Id:        uuid.New(),
	}

	client := newDbClientMock(t)

	_, err := client.GameHustler.
		Create().
		SetLastPosition(schema.Position{X: 20, Y: 20, CurrentMap: "dopecity"}).
		SetID(p.HustlerId).
		Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	_, err = client.GameHustlerRelation.
		Create().
		SetCitizen("tommy").
		SetText(0).
		SetConversation("robbery").
		SetHustlerID(p.HustlerId).
		SetID(uuid.NewString()).
		Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	g := Game{
		ItemEntities: []*item.ItemEntity{
			{
				Id: uuid.New(),
			},
		},
		Players: []*player.Player{
			{
				Id: uuid.New(),
			},
			&p,
		},
	}

	handshake := g.GenerateHandshakeData(context.TODO(), client, &p)
	var relations map[string]player.Relation
	err = json.Unmarshal(handshake.Relations, &relations)
	if err != nil {
		assert.FailNow(err.Error())
	}

	assert.Equal(p.Id.String(), handshake.Id)
	assert.Len(handshake.Players, 1)
	assert.Equal(g.ItemEntities[0].Id.String(), handshake.ItemEntities[0].Id)
	assert.Equal(g.Players[0].Id.String(), handshake.Players[0].Id)
	assert.Equal("robbery", relations["tommy"].Conversation)
}

func TestHandlePlayerJoin_BotOrNoHustler(t *testing.T) {
	assert := assert.New(t)

	g := Game{
		Register:  make(chan *player.Player, 1),
		Broadcast: make(chan messages.BroadcastMessage, 1),
	}

	g.HandlePlayerJoin(context.TODO(), nil, nil, nil)

	gOut := <-g.Register
	<-g.Broadcast

	assert.Empty(gOut.HustlerId)
	assert.Equal("Hustler", gOut.Name)
	assert.NotEmpty(gOut.Id)
}

func TestCheckWhitelist(t *testing.T) {
	assert := assert.New(t)

	assert.True(isWhitelisted("2", ""))
	assert.False(isWhitelisted("NotANum", ""))
}
