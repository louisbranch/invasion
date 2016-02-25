package message

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/luizbranco/invasion/internal/client"
)

type Code int

const (
	Join Code = iota
	Leave
	Disconnect
	ChatMessage
	SetName

	ErrorFormat
	ErrorInvalidCode
)

type Message struct {
	Code
	Content []byte
	Client  client.Client
}

func (m *Message) UnmarshalText(text []byte) error {
	s := bytes.SplitN(text, []byte(" "), 2)
	switch len(s) {
	case 2:
		m.Content = s[1]
		fallthrough
	case 1:
		i, err := strconv.Atoi(string(s[0]))
		if err != nil {
			return fmt.Errorf("failed to unmarshal code '%s': %s", text, err)
		}
		m.Code = Code(i)
	default:
		return fmt.Errorf("failed to unmarshal content '%s'", text)
	}
	return nil
}
