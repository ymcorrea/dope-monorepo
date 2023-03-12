package item

import (
	"testing"

	"github.com/dopedao/dope-monorepo/packages/api/game/dopemap"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewItemEntity(t *testing.T) {
	assert := assert.New(t)

	var x float32 = 20
	var y float32 = 20

	item := Item{
		Item: "coke",
	}

	itemEnt := NewItemEntity(item, x, y)

	assert.Equal(item, itemEnt.Item)
	assert.Equal(x, itemEnt.Position.X)
	assert.Equal(y, itemEnt.Position.Y)
}

func TestSerialize(t *testing.T) {
	assert := assert.New(t)

	var x float32 = 20
	var y float32 = 20

	item := Item{
		Item: "coke",
	}

	itemEnt := NewItemEntity(item, x, y)

	itemEntSer := itemEnt.Serialize()

	assert.Equal(item.Item, itemEntSer.Item)
	assert.Equal(x, itemEntSer.X)
	assert.Equal(y, itemEntSer.Y)
}

func TestGetItemEntityByUUID(t *testing.T) {
	assert := assert.New(t)

	gun := ItemEntity{
		Id: uuid.New(),
		Item: Item{
			Item: "gun",
		},
		Position: dopemap.Position{
			X: 20,
			Y: 20,
		},
	}

	coke := ItemEntity{
		Id: uuid.New(),
		Item: Item{
			Item: "coke",
		},
		Position: dopemap.Position{
			X: 30,
			Y: 30,
		},
	}

	items := make([]*ItemEntity, 2)
	items[0] = &gun
	items[1] = &coke

	res := GetItemEntityByUUID(items, gun.Id)

	assert.Equal(gun.Id, res.Id)
	assert.Equal(gun.Item, res.Item)
	assert.Equal(gun.Position, res.Position)
}

func TestGetItemEntityByUUID_ItemDoesntExist(t *testing.T) {
	assert := assert.New(t)

	res := GetItemEntityByUUID([]*ItemEntity{}, uuid.New())

	assert.Nil(res)
}
