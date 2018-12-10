package minip

import "fmt"

type Error struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func (e *Error) IsError() bool {
	return e.Code != 0
}

func (e *Error) IsBusy() bool {
	return e.Code == -1
}

func (e *Error) Error() string {
	if e.Code == 0 {
		return ""
	}

	return fmt.Sprintf("%d:%s", e.Code, e.Message)
}
