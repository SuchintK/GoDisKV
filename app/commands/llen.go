package command

import (
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/resp/client"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

type LLenCommand Command

func (cmd *LLenCommand) Execute(con *client.Client) RESPValue {
	if len(cmd.args) != 1 {
		return resp.EncodeSimpleError(errWrongNumberOfArgs)
	}

	key := cmd.args[0]
	val, exists := store.Get(key)

	// Key doesn't exist
	if !exists {
		return resp.EncodeInteger(0)
	}

	// Check if it's not a list
	if val.ListData == nil {
		return resp.EncodeSimpleError(errWrongType)
	}

	return resp.EncodeInteger(int64(len(val.ListData)))
}
