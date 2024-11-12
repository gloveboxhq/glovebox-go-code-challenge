package mockemail

import (
	"encoding/json"

	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email"
)

type SendLog struct {
	tos     []string
	ccs     []string
	tplID   email.TplID
	message json.RawMessage
}

func (sl *SendLog) ExtractTos() []string {
	return sl.tos
}

func (sl *SendLog) ExtractCCs() []string {
	return sl.ccs
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
