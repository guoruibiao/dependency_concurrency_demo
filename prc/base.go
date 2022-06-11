package prc

type prcConf struct {
	Name string
	Receivers []<-chan struct{}
	Notifiers []chan<- struct{}
}


type prcBase struct {
}

type handler interface {
	PreCheck() bool
	Do()
}


func (pb *prcBase) PreCheck() bool {
	return false
}

func (pb *prcBase) Do() bool {
	return false
}

