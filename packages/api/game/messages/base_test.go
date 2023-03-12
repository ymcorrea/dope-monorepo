package messages

import (
	"encoding/json"
	"testing"

	"github.com/dopedao/dope-monorepo/packages/api/game/events"
	"github.com/stretchr/testify/assert"
)

func TestGenerateError(t *testing.T) {
	assert := assert.New(t)
	code := 500
	reason := "fail"

	msg := GenerateErrorMessage(code, reason)

	var data ErrorMessageData
	json.Unmarshal(msg.Data, &data)

	assert.Equal(events.ERROR, msg.Event)
	assert.Equal(code, data.Code)
	assert.Equal(reason, data.Message)
}
