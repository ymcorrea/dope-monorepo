package player

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dopedao/dope-monorepo/packages/api/game/dopemap"
	"github.com/dopedao/dope-monorepo/packages/api/game/events"
	"github.com/dopedao/dope-monorepo/packages/api/game/item"
	"github.com/dopedao/dope-monorepo/packages/api/game/messages"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/enttest"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/hustler"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/schema"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func NewPlayerMock(broadcastChan *chan messages.BroadcastMessage, isBot bool) *Player {
	p := Player{
		Id:   uuid.New(),
		Send: make(chan messages.BaseMessage, 1),
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
		Chatcolor:  "blue",
	}

	if broadcastChan != nil {
		p.Broadcast = *broadcastChan
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

func TestMove(t *testing.T) {
	p := Player{}
	assert := assert.New(t)

	var x float32 = 20
	var y float32 = 10
	direction := "NORTH"

	p.Move(x, y, direction)

	assert.Equal(direction, p.Direction)
	assert.Equal(x, p.Position.X)
	assert.Equal(y, p.Position.Y)
}

func TestRemoveItemEntity(t *testing.T) {
	assert := assert.New(t)

	var items []*item.ItemEntity

	first := item.ItemEntity{
		Item: item.Item{
			Item: "first",
		},
	}

	second := item.ItemEntity{
		Item: item.Item{
			Item: "second",
		},
	}

	items = append(items, &first)
	items = append(items, &second)

	assert.True(RemoveItemEntity(&items, &second))
	assert.Len(items, 1)
	assert.Equal(items[0], &first)
}

func TestAddItem(t *testing.T) {
	assert := assert.New(t)

	p := NewPlayerMock(nil, false)

	client := newDbClientMock(t)

	_, err := client.GameHustler.Create().SetID(p.HustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	item := item.Item{
		Item: "gun",
	}

	expected := messages.BaseMessage{
		Event: events.PLAYER_ADD_ITEM,
	}

	if err := p.AddItem(context.TODO(), client, item, true); err != nil {
		assert.FailNow(err.Error())
	}

	out := <-p.Send

	assert.Equal(expected.Event, out.Event)
}

func TestAddItem_HandleNoHustler(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)

	p := Player{}

	if err := p.AddItem(context.TODO(), client, item.Item{}, true); err != nil {
		assert.ErrorContains(err, "player must have a hustler")
		return
	}

	assert.Fail("Did not throw error")
}

func TestAddItem_HandleDbError(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)

	p := NewPlayerMock(nil, false)

	if err := p.AddItem(context.TODO(), client, item.Item{}, true); err != nil {
		assert.Error(err)
		return
	}

	assert.Fail("Did not throw error")
}

func TestAddQuest(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)
	p := NewPlayerMock(nil, false)

	_, err := client.GameHustler.Create().SetID(p.HustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	quest := Quest{
		Quest:     "test_quest",
		Completed: false,
	}

	expected := messages.BaseMessage{
		Event: events.PLAYER_ADD_QUEST,
	}

	if err := p.AddQuest(context.TODO(), client, quest); err != nil {
		assert.FailNow(err.Error())
	}

	out := <-p.Send

	assert.Equal(expected.Event, out.Event)
}
func TestAddQuest_HandleNoHustler(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)

	p := Player{}

	if err := p.AddQuest(context.TODO(), client, Quest{}); err != nil {
		assert.Error(err)
		return
	}

	assert.Fail("Did not throw error")
}

func TestAddQuest_HandleDbError(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)

	p := Player{}

	if err := p.AddItem(context.TODO(), client, item.Item{}, true); err != nil {
		assert.Error(err)
		return
	}

	assert.Fail("Did not throw error")
}

func TestHandlePlayerLeave(t *testing.T) {
	assert := assert.New(t)

	broadcastChan := make(chan messages.BroadcastMessage, 1)
	unregisterChan := make(chan *Player, 1)

	p := NewPlayerMock(&broadcastChan, false)
	p.Unregister = unregisterChan

	unregister, _ := json.Marshal(item.IdData{
		Id: p.Id.String(),
	})

	moveMsg := messages.BroadcastMessage{
		Message: messages.BaseMessage{
			Event: events.PLAYER_LEAVE,
			Data:  unregister,
		},
	}

	handlePlayerLeave(p)

	broadcastOut := <-broadcastChan
	unregisterOut := <-unregisterChan

	assert.Equal(moveMsg, broadcastOut)
	assert.Equal(p, unregisterOut)
}

func TestHandlePlayerMove(t *testing.T) {
	assert := assert.New(t)

	previousPos := dopemap.Position{
		X: 10,
		Y: 10,
	}
	p := Player{
		Position: previousPos,
	}

	moveData := PlayerMoveData{
		X: 20,
		Y: 20,
	}
	moveJson, _ := json.Marshal(moveData)
	handlePlayerMove(&p, moveJson)

	assert.Equal(p.LastPosition.X, previousPos.X)
	assert.Equal(p.LastPosition.Y, previousPos.Y)

	assert.Equal(p.Position.X, moveData.X)
	assert.Equal(p.Position.Y, moveData.Y)
}

func TestHandlePlayerMove_InvalidJson(t *testing.T) {
	assert := assert.New(t)

	p := NewPlayerMock(nil, false)

	handlePlayerMove(p, json.RawMessage{})
	out := <-p.Send

	var err messages.ErrorMessageData
	json.Unmarshal(out.Data, &err)

	assert.Equal(500, err.Code)
	assert.NotEmpty(err.Message)
}

func TestHandlePlayerUpdateMap_InvalidJson(t *testing.T) {
	assert := assert.New(t)

	p := NewPlayerMock(nil, false)
	handlePlayerUpdateMap(p, json.RawMessage{}, &zerolog.Logger{})

	out := <-p.Send

	var err messages.ErrorMessageData
	json.Unmarshal(out.Data, &err)

	assert.Equal(500, err.Code)
	assert.NotEmpty(err.Message)
}

func TestHandlePlayerUpdateMap(t *testing.T) {
	assert := assert.New(t)

	gChannel := make(chan messages.BroadcastMessage, 1)
	p := NewPlayerMock(&gChannel, false)

	d := PlayerUpdateMapData{
		CurrentMap: "memphis",
		X:          20,
		Y:          20,
	}

	data, _ := json.Marshal(d)

	handlePlayerUpdateMap(p, data, &zerolog.Logger{})

	var outData PlayerUpdateMapClientData
	chanOut := <-gChannel

	json.Unmarshal(chanOut.Message.Data, &outData)

	assert.Equal(p.Id.String(), outData.Id)
	assert.Equal(events.PLAYER_UPDATE_MAP, chanOut.Message.Event)

	// dont broadcast to player who did the move
	assert.False(chanOut.Condition(p))
}

func TestHandlePlayerChatMessage(t *testing.T) {
	assert := assert.New(t)

	broadcastChan := make(chan messages.BroadcastMessage, 1)
	p := NewPlayerMock(&broadcastChan, false)

	chatMessage := ChatMessageData{
		Message: "wassup playa",
	}

	chatMessageJson, _ := json.Marshal(chatMessage)
	handlePlayerChatMessage(p, chatMessageJson, &zerolog.Logger{})

	out := <-broadcastChan

	var chatMessageClient ChatMessageClientData
	json.Unmarshal(out.Message.Data, &chatMessageClient)

	assert.Equal(events.PLAYER_CHAT_MESSAGE, out.Message.Event)
	assert.Equal(p.Id.String(), chatMessageClient.Author)
	assert.Equal(chatMessage.Message, chatMessageClient.Message)
	assert.Equal(p.Chatcolor, chatMessageClient.Color)
}

func TestHandlePlayerChatMessage_HandleEmptyMsg(t *testing.T) {
	assert := assert.New(t)

	broadcastChan := make(chan messages.BroadcastMessage)
	p := NewPlayerMock(&broadcastChan, false)
	chatMessage := ChatMessageData{}

	chatMessageJson, _ := json.Marshal(chatMessage)
	handlePlayerChatMessage(p, chatMessageJson, &zerolog.Logger{})

	var out messages.BroadcastMessage
	// dont block when nothing comes in
	select {
	case out = <-broadcastChan:
	default:
	}

	assert.Equal(messages.BroadcastMessage{}, out)
}

func TestHandlePlayerPickupItemEntity(t *testing.T) {
	assert := assert.New(t)

	broadcastChan := make(chan messages.BroadcastMessage, 1)
	p := NewPlayerMock(&broadcastChan, false)

	client := newDbClientMock(t)

	hustlerId := "test"
	log := zerolog.Logger{}
	_, err := client.GameHustler.Create().SetID(hustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	itemId := uuid.New()
	i := item.IdData{
		Id: itemId.String(),
	}

	var gameItems []*item.ItemEntity
	gameItems = append(gameItems, &item.ItemEntity{
		Id: itemId,
	})

	p.GameItems = gameItems

	itemJson, _ := json.Marshal(i)
	handlePlayerPickupItemEntity(p, itemJson, context.TODO(), &log, client)

	pOutBroadcast := <-broadcastChan

	var pOutBroadcastParsed item.IdData
	err = json.Unmarshal(pOutBroadcast.Message.Data, &pOutBroadcastParsed)
	if err != nil {
		assert.FailNow(err.Error())
	}

	assert.Equal(events.PLAYER_PICKUP_ITEMENTITY, pOutBroadcast.Message.Event)
	assert.Equal(i.Id, pOutBroadcastParsed.Id)
}

func TestHandlePlayerPickupItemEntity_NoHustler(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)
	p := NewPlayerMock(nil, true)

	handlePlayerPickupItemEntity(p, json.RawMessage{}, context.TODO(), &zerolog.Logger{}, client)

	out := <-p.Send

	var err messages.ErrorMessageData
	json.Unmarshal(out.Data, &err)

	assert.Equal(500, err.Code)
	assert.Contains(err.Message, "must have a hustler")
}

func TestHandlePlayerPickupItemEntity_InvalidItemId(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)
	p := NewPlayerMock(nil, false)

	i := item.IdData{
		Id: "fake",
	}

	itemJson, _ := json.Marshal(i)

	handlePlayerPickupItemEntity(p, itemJson, context.TODO(), &zerolog.Logger{}, client)

	out := <-p.Send

	var err messages.ErrorMessageData
	json.Unmarshal(out.Data, &err)

	assert.Equal(500, err.Code)
	assert.Contains(err.Message, "could not parse item")
}

func TestHandlePlayerPickupItemEntity_ItemNotInGame(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)
	p := NewPlayerMock(nil, false)

	i := item.IdData{
		Id: uuid.NewString(),
	}

	itemJson, _ := json.Marshal(i)

	handlePlayerPickupItemEntity(p, itemJson, context.TODO(), &zerolog.Logger{}, client)

	out := <-p.Send

	var err messages.ErrorMessageData
	json.Unmarshal(out.Data, &err)

	assert.Equal(500, err.Code)
	assert.Contains(err.Message, "could not find item")
}

func TestHandlePlayerUpdateCitizenState(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)
	p := NewPlayerMock(nil, false)

	_, dbErr := client.GameHustler.Create().SetID(p.HustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if dbErr != nil {
		assert.FailNow(dbErr.Error())
	}

	data := CitizenUpdateStateData{
		Citizen:      "tommy",
		Conversation: "wassup foo",
		Text:         1,
	}

	dataJson, _ := json.Marshal(data)
	handlePlayerUpdateCitizenState(p, dataJson, context.TODO(), &zerolog.Logger{}, client)

	res, _ := client.GameHustlerRelation.Get(context.TODO(), fmt.Sprintf("%s:%s", p.HustlerId, data.Citizen))

	assert.Equal(data.Conversation, res.Conversation)
	assert.Equal(data.Citizen, res.Citizen)
	assert.Equal(data.Text, res.Text)
}

func TestHandlePlayerUpdateCitizenState_NoHustlerId(t *testing.T) {
	assert := assert.New(t)
	client := newDbClientMock(t)

	p := NewPlayerMock(nil, true)

	handlePlayerUpdateCitizenState(p, json.RawMessage{}, context.TODO(), &zerolog.Logger{}, client)

	pChanOut := <-p.Send

	var errMsg messages.ErrorMessageData
	json.Unmarshal(pChanOut.Data, &errMsg)

	assert.Equal(500, errMsg.Code)
	assert.Contains(errMsg.Message, "must have a hustler")
}

func TestHandlePlayerUpdateCitizenState_InvalidJson(t *testing.T) {
	assert := assert.New(t)

	client := newDbClientMock(t)

	p := NewPlayerMock(nil, false)

	handlePlayerUpdateCitizenState(p, json.RawMessage{}, context.TODO(), &zerolog.Logger{}, client)

	out := <-p.Send
	var err messages.ErrorMessageData
	json.Unmarshal(out.Data, &err)

	assert.Equal(500, err.Code)
	assert.Contains(err.Message, "could not unmarshal")
}

func TestNewPlayer(t *testing.T) {
	assert := assert.New(t)

	name := "Paul"
	id := uuid.NewString()
	currentMap := "DopeCity"
	var x float32 = 20
	var y float32 = 30
	conn := &websocket.Conn{}
	broadcast := make(chan messages.BroadcastMessage)
	unregister := make(chan *Player)
	gameItems := make([]*item.ItemEntity, 0)

	p := NewPlayer(
		conn,
		broadcast,
		unregister,
		id,
		name,
		currentMap,
		x,
		y,
		gameItems,
	)

	assert.Equal(conn, p.Conn)
	assert.Equal(name, p.Name)
	assert.Equal(broadcast, p.Broadcast)
	assert.Equal(unregister, p.Unregister)
	assert.Equal(gameItems, p.GameItems)
	assert.Equal(x, p.Position.X)
	assert.Equal(y, p.Position.Y)
	assert.Equal(x, p.LastPosition.X)
	assert.Equal(y, p.LastPosition.Y)
	assert.Equal(id, p.HustlerId)
	assert.Equal("white", p.Chatcolor)
}

func TestReadPump_LeaveOnErr(t *testing.T) {
	assert := assert.New(t)
	client := newDbClientMock(t)

	p := NewPlayerMock(nil, false)
	p.Unregister = make(chan *Player, 1)

	go p.ReadPump(context.TODO(), client)

	out := <-p.Unregister

	assert.Equal(p, out)
	assert.Equal(p.Id, out.Id)
}

func TestSerialize(t *testing.T) {
	assert := assert.New(t)

	p := NewPlayerMock(nil, false)
	serialized := p.Serialize()
	sId, _ := uuid.Parse(serialized.Id)

	assert.Equal(p.Id, sId)
	assert.Equal(p.HustlerId, serialized.HustlerId)
	assert.Equal(p.Name, serialized.Name)
	assert.Equal(p.CurrentMap, serialized.CurrentMap)
	assert.Equal(p.Position.X, serialized.X)
	assert.Equal(p.Position.Y, serialized.Y)
}

func TestPopulateFromDB(t *testing.T) {
	assert := assert.New(t)
	client := newDbClientMock(t)

	p := NewPlayerMock(nil, false)

	hustlerName := "crunchy black"
	_, err := client.Hustler.
		Create().
		SetName(hustlerName).
		SetType(hustler.TypeREGULAR).
		SetAge(2).
		SetID(p.Id.String()).
		Save(context.TODO())
	if err != nil {
		assert.FailNow("hustler: ", err.Error())
	}

	item := item.Item{
		Item: "gun",
	}
	itemId := uuid.NewString()
	_, err = client.GameHustlerItem.
		Create().
		SetItem(item.Item).
		SetID(itemId).
		Save(context.TODO())
	if err != nil {
		assert.FailNow("item: ", err.Error())
	}

	quest := Quest{
		Quest:     "rob and steal",
		Completed: true,
	}
	questId := uuid.NewString()
	_, err = client.GameHustlerQuest.
		Create().
		SetQuest(quest.Quest).
		SetCompleted(quest.Completed).
		SetID(questId).
		Save(context.TODO())
	if err != nil {
		assert.FailNow("quest: ", err.Error())
	}

	gameHustler, err := client.GameHustler.
		Create().
		SetID(p.Id.String()).
		AddItemIDs(itemId).
		AddQuestIDs(questId).
		SetLastPosition(schema.Position{CurrentMap: "dopecity", X: p.LastPosition.X, Y: p.LastPosition.Y}).
		Save(context.TODO())
	if err != nil {
		assert.FailNow("gameHustler: ", err.Error())
	}

	if err := p.PopulateFromDB(context.TODO(), client, gameHustler, &zerolog.Logger{}); err != nil {
		assert.FailNow(err.Error())
	}

	assert.Equal(hustlerName, p.Name)
	assert.Equal(p.Items[0], item)
	assert.Equal(p.Quests[0], quest)
}
