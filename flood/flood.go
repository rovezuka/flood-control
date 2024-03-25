package flood

import (
	"context"
	"sync"
	"time"
)

type floodControl struct {
	callCounts map[int64]int64
	lock       sync.Mutex // Для безопасного доступа к map из нескольких goroutine
	N          int64      // Период времени в секундах
	K          int64      // Максимальное количество вызовов в период времени N
}

func NewFloodControl(N, K int64) *floodControl {
	return &floodControl{
		callCounts: make(map[int64]int64),
		N:          N,
		K:          K,
	}
}

func (fc *floodControl) Check(ctx context.Context, userID int64) (bool, error) {
	fc.lock.Lock()
	defer fc.lock.Unlock()

	currentTime := time.Now().Unix()

	// Проверяем, были ли вызовы метода Check для этого пользователя
	if lastCallTime, ok := fc.callCounts[userID]; ok {
		// Если последний вызов был в течение последних N секунд
		if currentTime-lastCallTime <= fc.N {
			fc.callCounts[userID]++
			// Проверяем, не превышено ли максимальное количество вызовов K
			if fc.callCounts[userID] > fc.K {
				return false, nil
			}
		} else {
			// Если прошло больше N секунд с последнего вызова, сбрасываем счетчик
			fc.callCounts[userID] = 1
		}
	} else {
		// Если это первый вызов метода Check для данного пользователя
		fc.callCounts[userID] = 1
	}

	return true, nil
}

func (fc *floodControl) SetConfig(N, K int64) {
	fc.lock.Lock()
	defer fc.lock.Unlock()

	fc.N = N
	fc.K = K
}
