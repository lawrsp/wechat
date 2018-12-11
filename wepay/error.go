package wepay

import "fmt"

type Error struct {
	Code    string `xml:"return_code"` //SUCCESS/FAIL
	Message string `xml:"return_msg"`
}

func (e *Error) IsError() bool {
	return e.Code != "SUCCESS"
}

func (e *Error) Error() string {
	if e.Code == "SUCCESS" {
		return ""
	}

	return fmt.Sprintf("%s:%s", e.Code, e.Message)
}
