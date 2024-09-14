package database

import (
	"go-redis/interface/resp"
	"go-redis/resp/reply"
)

type EchoDatebase struct {
}

func NewEchoDatebase() *EchoDatebase {
	return &EchoDatebase{}
}
func (e EchoDatebase) Exec(clienr resp.Connection, args [][]byte) resp.Reply {
	return reply.MakeMultiBulkReply(args)
}

func (e EchoDatebase) Close() {
	//TODO implement me
	panic("implement me")
}

func (e EchoDatebase) AfterClientClose(c resp.Connection) {
	//TODO implement me
	panic("implement me")
}
