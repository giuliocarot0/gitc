package gitc

type MessageType int

const (
	MSG0  MessageType = iota
	MSG1  MessageType = iota
	MSG2  MessageType = iota
	MSG3  MessageType = iota
	MSG4  MessageType = iota
	MSG5  MessageType = iota
	MSG6  MessageType = iota
	MSG7  MessageType = iota
	MSG8  MessageType = iota
	MSG9  MessageType = iota
	MSG10 MessageType = iota
	MSG11 MessageType = iota
	MSG12 MessageType = iota
	MSG13 MessageType = iota
	MSG14 MessageType = iota
)

type Message struct {
	From    string      // Sender task name
	To      string      // Receiver task name
	Type    MessageType // The type of expected payload
	Payload interface{} // Actual message content
}
