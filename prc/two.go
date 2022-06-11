package prc

import (
	"fmt"
	"time"
)

type TwoPrc struct {
	Name string
}

func NewTwoPrc(conf *prcConf) *TwoPrc {
	inst := TwoPrc{}
	return &inst
}

func (tp *TwoPrc) PreCheck() bool {

	return true
}

func (tp *TwoPrc) Do() {
	time.Sleep(time.Second)
	fmt.Println("TwoPrc time.Sleep=1")
}

