package player

import (
	"encoding/json"
	"testing"

	"github.com/dopedao/dope-monorepo/packages/api/game/events"
	"github.com/dopedao/dope-monorepo/packages/api/game/messages"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestHandePlayerCommand_InvalidCommand(t *testing.T) {
	assert := assert.New(t)

	broadcast := make(chan messages.BroadcastMessage, 1)

	p := Player{
		Broadcast: broadcast,
	}

	cmd := ChatCommandData{
		Name: uuid.NewString(),
	}

	cmdJson, _ := json.Marshal(cmd)

	handlePlayerCommand(&p, cmdJson, &zerolog.Logger{})

	out := <-broadcast

	var toast messages.Toast
	json.Unmarshal(out.Message.Data, &toast)

	assert.Equal(events.PLAYER_CHAT_COMMAND_RESULT, out.Message.Event)
	assert.Equal(messages.ERROR, toast.Status)
	assert.True(out.Condition(&p))
}

func TestHandePlayerCommand_InvalidJson(t *testing.T) {
	assert := assert.New(t)

	broadcast := make(chan messages.BroadcastMessage, 1)
	pChan := make(chan messages.BaseMessage, 1)

	p := Player{
		Broadcast: broadcast,
		Send:      pChan,
	}

	handlePlayerCommand(&p, json.RawMessage{}, &zerolog.Logger{})

	out := <-pChan

	var err messages.ErrorMessageData
	json.Unmarshal(out.Data, &err)

	assert.Equal(events.ERROR, out.Event)
	assert.NotEmpty(err.Message)
	assert.Equal(500, err.Code)
}

func TestHandleSetcolor(t *testing.T) {
	assert := assert.New(t)

	broadcast := make(chan messages.BroadcastMessage, 1)

	p := Player{
		Broadcast: broadcast,
	}

	args := make([]string, 1)
	args[0] = allowedColors[0]

	handleSetcolor(args, &p, &zerolog.Logger{})

	out := <-broadcast

	var toast messages.Toast
	json.Unmarshal(out.Message.Data, &toast)

	assert.Equal(events.PLAYER_CHAT_COMMAND_RESULT, out.Message.Event)
	assert.Equal(messages.SUCCESS, toast.Status)
	assert.Equal(allowedColors[0], p.Chatcolor)
	assert.True(out.Condition(&p))
}

func TestHandleSetcolor_InvalidColor(t *testing.T) {
	assert := assert.New(t)

	broadcast := make(chan messages.BroadcastMessage, 1)

	p := Player{
		Broadcast: broadcast,
	}

	args := make([]string, 1)
	args[0] = "fake"

	handleSetcolor(args, &p, &zerolog.Logger{})

	out := <-broadcast

	var toast messages.Toast
	json.Unmarshal(out.Message.Data, &toast)

	assert.Equal(messages.ERROR, toast.Status)
	assert.NotEqual(args[0], p.Chatcolor)
}
