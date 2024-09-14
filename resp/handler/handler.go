package handler

import (
	"context"
	"go-redis/database"
	databaseface "go-redis/interface/database"
	"go-redis/lib/logger"
	"go-redis/lib/sync/atomic"
	"go-redis/resp/connection"
	"go-redis/resp/parser"
	"go-redis/resp/reply"
	"io"
	"net"
	"strings"
	"sync"
)

var (
	unknowErrReplyBytes = []byte("-Err unknown\r\n")
)

func MakeHandler() *RespHandler {
	var db databaseface.Database
	db = database.NewEchoDatebase()
	return &RespHandler{
		db: db,
	}

}

type RespHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
	db         databaseface.Database
}

func (r *RespHandler) closeClient(client *connection.Connection) {
	_ = client.Close()
	r.db.AfterClientClose(client)
	r.activeConn.Delete(client)
}

func (r *RespHandler) Handle(ctx context.Context, conn net.Conn) {
	if r.closing.Get() {
		_ = conn.Close()
	}
	client := connection.NewConn(conn)
	r.activeConn.Store(client, struct{}{})
	ch := parser.ParseStream(conn)
	for payload := range ch {
		if payload.Err != nil {
			if payload.Err == io.EOF || payload.Err == io.ErrUnexpectedEOF || strings.Contains(payload.Err.Error(), "use of closed network connection") {
				r.closeClient(client)
				logger.Info("connection closed" + client.RemoteAddr().String())
				return
			}
			//协议错误
			errReply := reply.MakeErrReply(payload.Err.Error())
			err := client.Write(errReply.ToBytes())
			if err != nil {
				r.closeClient(client)
				logger.Info("connectong closed" + client.RemoteAddr().String())
				return
			}
			continue
		}
		if payload.Data == nil {
			continue
		}
		reply, ok := payload.Data.(*reply.MultiBulkReply)
		if !ok {
			logger.Error("require multi bulk reply")
			continue
		}
		result := r.db.Exec(client, reply.Args)
		if result != nil {
			_ = client.Write(result.ToBytes())
		} else {
			_ = client.Write(unknowErrReplyBytes)
		}
	}
}

func (r *RespHandler) Close() error {
	logger.Info("handler shuting down")
	r.closing.Set(true)
	r.activeConn.Range(
		func(key interface{}, value interface{}) bool {
			client := key.(*connection.Connection)
			_ = client.Close()
			return true
		})
	r.db.Close()
	return nil
}
