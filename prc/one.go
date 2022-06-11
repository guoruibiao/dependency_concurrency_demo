package prc

import (
	"fmt"
	"time"
)

type OnePrc struct {
	Name string
}

func NewOnePrc(conf *prcConf) *OnePrc {
	inst := OnePrc{}
	return &inst
}

func (op *OnePrc) PreCheck() bool {
	return true
}

func (op *OnePrc) Do() {
	time.Sleep(time.Second*2)
	fmt.Println("OnePrc time.Sleep=2")
}

