package command

import (
	"github.com/SuchintK/GoDisKV/app/pubsub"
	"github.com/SuchintK/GoDisKV/app/resp"
	"github.com/SuchintK/GoDisKV/app/resp/client"
)

type SubscribeCommand Command

func (cmd *SubscribeCommand) Execute(con *client.Client) RESPValue {
	if len(cmd.args) != 1 {
		return resp.EncodeSimpleError(errWrongNumberOfArgs)
	}

	channel := cmd.args[0]
	count := pubsub.Global.Subscribe(con, channel)

	return resp.EncodePubSubResponse("subscribe", channel, count)
}
