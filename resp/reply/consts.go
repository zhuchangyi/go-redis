package reply

type PongReply struct {
}

var pongbytes = []byte("+PONG\r\n")

func (r PongReply) ToBytes() []byte {
	return pongbytes
}

func MakePongReply() *PongReply {
	return &PongReply{}
}

type OkReply struct{}

var okBytes = []byte("+OK\r\n")

// ToBytes marshal redis.Reply
func (r *OkReply) ToBytes() []byte {
	return okBytes
}

var theOkReply = new(OkReply)

// MakeOkReply returns a ok reply
func MakeOkReply() *OkReply {
	return theOkReply
}

var nullBulkyBytes = []byte("$-1\r\n")

type NullBulkReply struct{}

func (n NullBulkReply) ToBytes() []byte {
	return nullBulkyBytes
}
func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

type EmptyMultiBulkReply struct{}

var emptyMultiBulkReply = []byte("*0\r\n")

func (e EmptyMultiBulkReply) ToBytes() []byte {
	//TODO implement me
	return emptyMultiBulkReply
}

type NoReply struct{}

var noreply = []byte("")

func (NoReply) ToBytes() []byte {

	return noreply
}
