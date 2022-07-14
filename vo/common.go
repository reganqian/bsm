package vo

import (
	. "common/static"
)

type BaseReply struct {
	ResCode string `json:"resCode"`
	ResDesc string `json:"resDesc"`
}

func (t *BaseReply) Success() {
	t.ResCode = SUCCESS
	t.ResDesc = DEFAULT_SUCCESS_DESC
}

func (t *BaseReply) Failed(resCode, resDesc string) {
	t.ResCode = resCode
	t.ResDesc = resDesc
}

type BoolReply struct {
	BaseReply
	Data bool `json:"data"`
}

type IntReply struct {
	BaseReply
	Data int `json:"data"`
}

type StringReply struct {
	BaseReply
	Data string `json:"data"`
}

type Float64Reply struct {
	BaseReply
	Data float64 `json:"data"`
}
