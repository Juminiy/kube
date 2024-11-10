package safe_go

import (
	"runtime/debug"
	"sync"
	"time"
)

func Go(
	wg *sync.WaitGroup,
	ignoreRecover bool,
	doRecoverStack func(recoverTime time.Time, debugStack []byte),
	handler func(),
	catchHandler func(v ...any),
) {
	if wg != nil {
		wg.Add(1)
	}
	go func() {
		defer func() {
			if rcv := recover(); rcv != nil {
				if !ignoreRecover {
					doRecoverStack(time.Now(), debug.Stack())
				}
				if catchHandler != nil {
					if wg != nil {
						wg.Add(1)
					}
					go func() {
						defer func() {
							if rcv2 := recover(); rcv2 != nil {
								if !ignoreRecover {
									doRecoverStack(time.Now(), debug.Stack())
								}
							}
							if wg != nil {
								wg.Done()
							}
						}()
						catchHandler(rcv)
					}()
				}
			}
			if wg != nil {
				wg.Done()
			}
		}()
		handler()
	}()
}
