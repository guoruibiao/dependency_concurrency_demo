package prc

import (
	"fmt"
	"time"
)

type ThreePrc struct {
	Name string
}

func NewThreePrc(conf *prcConf) *ThreePrc {
	inst := ThreePrc{}
	return &inst
}

func (tp *ThreePrc) PreCheck() bool {

	return true
}

func (tp *ThreePrc) Do() {
	time.Sleep(time.Second)
	fmt.Println("ThreePrc time.Sleep=1")
}

