package player

import (
	"context"
	"testing"

	"entgo.io/ent/entc/integration/config/ent/enttest"
	"github.com/dopedao/dope-monorepo/packages/api/game/item"
	"github.com/dopedao/dope-monorepo/packages/api/game/messages"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent"
	"github.com/stretchr/testify/assert"
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

	baseChan := make(chan messages.BaseMessage)
	defer close(baseChan)

	p := Player{
		Send: baseChan,
	}

	item := item.Item{
		Item: "gun",
	}

	go func() {
		p.AddItem(context.TODO(), , item, true)
	}()

	var got messages.BaseMessage
	select {
	case got = <-p.Send:
		break
	}

}

func TestAddQuest(t *testing.T) {

}

func TestReadPump(t *testing.T) {

}

func TestWritePump(t *testing.T) {

}

func TestHandlePlayerLeave(t *testing.T) {

}

func TestHandlePlayerUpdateMap(t *testing.T) {

}

func TestHandlePlayerChatMessage(t *testing.T) {

}

func TestHandlePlayerPickupItemEntity(t *testing.T) {

}

func TestHandlePlayerUpdateCitizenState(t *testing.T) {

}
