package database

import "go-redis/interface/resp"

type CmdLine = [][]byte

type Database interface {
	Exec(clienr resp.Connection, args [][]byte) resp.Reply
	Close()
	AfterClientClose(c resp.Connection)
}

type DataEntity struct {
	Data interface{}
}
