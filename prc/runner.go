package prc

import (
	"fmt"
	"sync"
)

func Process() {
	wg := &sync.WaitGroup{}
	wg.Add(len(PrcConfs))
	for _, conf := range PrcConfs {
		tmpConf := conf
		go run(wg, &tmpConf)
	}

	wg.Wait()
}

func createHandler(conf *prcConf) handler {
	switch conf.Name {
	case "OnePrc":
		return NewOnePrc(conf)
	case "TwoPrc":
		return NewTwoPrc(conf)
	case "ThreePrc":
		return NewThreePrc(conf)
	}
	return nil
}

func run(wg *sync.WaitGroup, conf *prcConf) {
	defer func() {
		if conf.Notifiers != nil && len(conf.Notifiers) > 0 {
			for _, ch := range conf.Notifiers {
				close(ch)
			}
		}
		wg.Done()
	}()

	receivers := conf.Receivers
	if receivers != nil && len(receivers) > 0 {
		for _, ch := range receivers {
			<-ch
		}
	}

	handler := createHandler(conf)
	if handler == nil {
		fmt.Println("HERE???")
		return
	}


	if !handler.PreCheck() {
		fmt.Println(conf.Name + " failed")
		return
	}

	handler.Do()
	fmt.Println(conf.Name + " success")
}