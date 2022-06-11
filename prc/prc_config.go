package prc

type prcs []prcConf

var PrcConfs prcs

func init() {
	PrcConfs = make([]prcConf, 0)

	// 假设 tow 依赖 one 的执行完毕
	oneChan := make(chan struct{})
	threeChan := make(chan struct{})

	PrcConfs = append(PrcConfs, prcConf{
		Name: "TwoPrc",
		Receivers: []<-chan struct{}{
			oneChan, threeChan,
		},
	})

	PrcConfs = append(PrcConfs, prcConf{
		Name: "OnePrc",
		Notifiers: []chan<- struct{}{
			oneChan,
		},
	})

	PrcConfs = append(PrcConfs, prcConf{
		Name: "ThreePrc",
		Notifiers: []chan<- struct{}{
			threeChan,
		},
	})
}

