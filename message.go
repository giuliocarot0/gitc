package gitc

type MessageType int

const (
	MSG0 MessageType = iota
	MSG1
	MSG2
	MSG3
	MSG4
	MSG5
	MSG6
	MSG7
	MSG8
	MSG9
	MSG10
	MSG11
	MSG12
	MSG13
	MSG14
)

type Message struct {
	From    string      // Sender task name
	To      string      // Receiver task name
	Type    MessageType // The type of expected payload
	Payload interface{} // Actual message content
}
