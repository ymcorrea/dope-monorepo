package player

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/dopedao/dope-monorepo/packages/api/game/dopemap"
	"github.com/dopedao/dope-monorepo/packages/api/game/events"
	"github.com/dopedao/dope-monorepo/packages/api/game/item"
	"github.com/dopedao/dope-monorepo/packages/api/game/messages"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/enttest"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/schema"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestMove(t *testing.T) {
	p := Player{}
	assert := assert.New(t)

	var x float32 = 20
	var y float32 = 10
	direction := "NORTH"

	p.Move(x, y, direction)

	assert.Equal(direction, p.Direction)
	assert.Equal(p.Position.X, x)
	assert.Equal(p.Position.Y, y)
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

	testChan := make(chan messages.BaseMessage, 1)
	defer close(testChan)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	hustlerId := "test"
	_, err := client.GameHustler.Create().SetID(hustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	p := Player{
		HustlerId: hustlerId,
		Send:      testChan,
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

	out := <-testChan

	assert.Equal(expected.Event, out.Event)
}

func TestAddItem_HandleNoHustler(t *testing.T) {
	assert := assert.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	p := Player{}

	if err := p.AddItem(context.TODO(), client, item.Item{}, true); err != nil {
		assert.ErrorContains(err, "player must have a hustler")
		return
	}

	assert.Fail("Did not throw error")
}

func TestAddItem_HandleDbError(t *testing.T) {
	assert := assert.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	p := Player{
		HustlerId: "123",
	}

	if err := p.AddItem(context.TODO(), client, item.Item{}, true); err != nil {
		assert.Error(err)
		return
	}

	assert.Fail("Did not throw error")
}

func TestAddQuest(t *testing.T) {
	assert := assert.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	testChan := make(chan messages.BaseMessage, 1)

	hustlerId := "test"
	_, err := client.GameHustler.Create().SetID(hustlerId).SetLastPosition(schema.Position{X: 10, Y: 10}).Save(context.TODO())
	if err != nil {
		assert.FailNow(err.Error())
	}

	p := Player{
		HustlerId: hustlerId,
		Send:      testChan,
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

	out := <-testChan

	assert.Equal(expected.Event, out.Event)
}
func TestAddQuest_HandleNoHustler(t *testing.T) {
	assert := assert.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	p := Player{}

	if err := p.AddQuest(context.TODO(), client, Quest{}); err != nil {
		assert.ErrorContains(err, "player must have a hustler")
		return
	}

	assert.Fail("Did not throw error")
}

func TestAddQuest_HandleDbError(t *testing.T) {
	assert := assert.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

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
	defer close(broadcastChan)

	unregisterChan := make(chan *Player, 1)
	defer close(unregisterChan)

	playerChan := make(chan messages.BaseMessage)

	pId := uuid.New()
	p := Player{
		Id:         pId,
		Unregister: unregisterChan,
		Broadcast:  broadcastChan,
		Send:       playerChan,
	}

	unregister, _ := json.Marshal(item.IdData{
		Id: p.Id.String(),
	})

	moveMsg := messages.BroadcastMessage{
		Message: messages.BaseMessage{
			Event: events.PLAYER_LEAVE,
			Data:  unregister,
		},
	}

	handlePlayerLeave(&p)

	broadcastOut := <-broadcastChan
	unregisterOut := <-unregisterChan

	assert.Equal(moveMsg, broadcastOut)
	assert.Equal(&p, unregisterOut)
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

func TestHandlePlayerUpdateMap(t *testing.T) {
	assert := assert.New(t)

	broadcastChan := make(chan messages.BroadcastMessage, 1)
	p := Player{
		Id:        uuid.New(),
		Broadcast: broadcastChan,
	}

	log := zerolog.Logger{}

	d := PlayerUpdateMapData{
		CurrentMap: "memphis",
		X:          20,
		Y:          20,
	}

	data, _ := json.Marshal(d)

	handlePlayerUpdateMap(&p, data, &log)

	var outData PlayerUpdateMapClientData
	chanOut := <-broadcastChan

	json.Unmarshal(chanOut.Message.Data, &outData)

	assert.Equal(p.Id.String(), outData.Id)
	assert.Equal(events.PLAYER_UPDATE_MAP, chanOut.Message.Event)

	// dont broadcast to player who did the move
	assert.False(chanOut.Condition(p))
}

func TestHandlePlayerUpdateMap_HandleInvalidMapdata(t *testing.T) {

}

func TestHandlePlayerChatMessage(t *testing.T) {

}

func TestHandlePlayerPickupItemEntity(t *testing.T) {

}

func TestHandlePlayerUpdateCitizenState(t *testing.T) {

}
