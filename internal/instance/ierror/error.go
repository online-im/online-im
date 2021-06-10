package ierror

import (
	"github.com/online-im/online-im/internal/instance/conn_holder"
	"github.com/online-im/online-im/pkg/constant"
)

type Error struct {
	err  error
	Code constant.CoreErrorCode
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) SendToConn(fromID string) {
	conn_holder.GetWSConnHolder().SendMessageWithErr(fromID, e.Code, e.err.Error())
}

func NewError(code constant.CoreErrorCode, err error) error {
	return &Error{
		err:  err,
		Code: code,
	}
}

func SendError(err error, fromID string) {
	if ierr, ok := err.(*Error); ok {
		ierr.SendToConn(fromID)
		return
	}
	conn_holder.GetWSConnHolder().SendMessageWithErr(fromID, constant.CoreErrorCode_Unknown, err.Error())
}
