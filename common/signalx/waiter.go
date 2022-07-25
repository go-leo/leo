package signalx

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type SignalHook = map[os.Signal]func()

type SignalWaiter struct {
	signals        []os.Signal
	signalC        chan os.Signal
	incomingSignal os.Signal
	hooks          []func(os.Signal)
	waitTimeout    time.Duration
	stopC          chan any
	locker         sync.RWMutex
}

func NewSignalWaiter(signals []os.Signal, waitTimeout time.Duration) *SignalWaiter {
	w := &SignalWaiter{
		signals:        signals,
		signalC:        make(chan os.Signal),
		incomingSignal: nil,
		hooks:          make([]func(os.Signal), 0),
		waitTimeout:    waitTimeout,
		stopC:          make(chan any, 1),
		locker:         sync.RWMutex{},
	}
	signal.Notify(w.signalC, w.signals...)
	return w
}

func (w *SignalWaiter) AddHook(f func(os.Signal)) *SignalWaiter {
	w.locker.Lock()
	defer w.locker.Unlock()
	if f == nil {
		return w
	}
	w.hooks = append(w.hooks, f)
	return w
}

func (w *SignalWaiter) KillSelf(signum syscall.Signal) error {
	return syscall.Kill(syscall.Getpid(), signum)
}

func (w *SignalWaiter) WaitSignals() *SignalWaiter {
	w.incomingSignal = <-w.signalC
	return w
}

func (w *SignalWaiter) WaitHooksAsyncInvoked() *SignalWaiter {
	go func(sig os.Signal) {
		w.locker.RLock()
		defer w.locker.RUnlock()
		defer close(w.stopC)
		for i := range w.hooks {
			f := w.hooks[len(w.hooks)-1-i]
			f(sig)
		}
	}(w.incomingSignal)
	return w
}

func (w *SignalWaiter) WaitUntilTimeout() *SignalWaiter {
	select {
	case <-w.stopC:
		return w
	case w.incomingSignal = <-w.signalC:
		return w
	case <-time.After(w.waitTimeout):
		return w
	}
}

func (w *SignalWaiter) Signal() os.Signal {
	return w.incomingSignal
}

func (w *SignalWaiter) Err() error {
	w.locker.RLock()
	defer w.locker.RUnlock()
	if w.incomingSignal == nil {
		return nil
	}
	return &SignalError{Signal: w.incomingSignal}
}
