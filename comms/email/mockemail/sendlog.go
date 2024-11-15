package mockemail

import (
	"encoding/json"

	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email"
)

type SendLog struct {
	to      string
	cc      string
	tplID   email.TplID
	message json.RawMessage
}

func (sl *SendLog) ExtractTo() string {
	return sl.to
}

func (sl *SendLog) ExtractCc() string {
	return sl.cc
}

func (sl *SendLog) ExtractTplID() email.TplID {
	return sl.tplID
}

func (sl *SendLog) ExtractMessage() json.RawMessage {
	return sl.message
}

type SendLogs []SendLog

func (s SendLogs) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// Last will return the last created log
func (s SendLogs) Last() *SendLog {
	return &s[len(s)-1]
}
