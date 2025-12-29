package command

import (
	"github.com/SuchintK/GoDisKV/app/resp"
	"github.com/SuchintK/GoDisKV/app/resp/client"
)

type EchoCommand Command

func (cmd *EchoCommand) Execute(con *client.Client) RESPValue {
	if len(cmd.args) != 1 {
		return resp.EncodeSimpleError(errWrongNumberOfArgs)
	}
	return resp.EncodeBulkString(cmd.args[0])
}
