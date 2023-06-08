package signalx

import (
	"os"
	"os/signal"
)

func AsyncWatch(sig ...os.Signal) <-chan os.Signal {
	signalC := make(chan os.Signal)
	go func() {
		signalC <- SyncWatch(sig...)
		close(signalC)
	}()
	return signalC
}

func SyncWatch(sig ...os.Signal) os.Signal {
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, sig...)
	incomingSignal := <-signalC
	signal.Stop(signalC)
	close(signalC)
	return incomingSignal
}
