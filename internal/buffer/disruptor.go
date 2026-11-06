package buffer

import (
	"context"
	"runtime"
	"sync/atomic"
)

type BetEvent struct {
	ID        uint64
	UserID    uint64
	Amount    int64
	Odds      float64
	IsSuccess bool
}

type DisruptorBuffer struct {
	capacity uint64
	mask     uint64
	ring     []BetEvent
	writeSeq uint64
	readSeq  uint64
}

func NewDisruptorBuffer(sizeLog2 uint32) *DisruptorBuffer {
	capacity := uint64(1 << sizeLog2)
	return &DisruptorBuffer{
		capacity: capacity,
		mask:     capacity - 1,
		ring:     make([]BetEvent, capacity),
	}
}

func (db *DisruptorBuffer) Publish(id, userID uint64, amount int64, odds float64) bool {
	nextWrite := atomic.AddUint64(&db.writeSeq, 1)

	for nextWrite-atomic.LoadUint64(&db.readSeq) > db.capacity {
		runtime.Gosched()
	}

	idx := nextWrite & db.mask
	db.ring[idx] = BetEvent{
		ID:     id,
		UserID: userID,
		Amount: amount,
		Odds:   odds,
	}
	return true
}

func (db *DisruptorBuffer) StartConsumer(ctx context.Context, processFn func(event *BetEvent)) {
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		localReadSeq := uint64(0)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				currentWrite := atomic.LoadUint64(&db.writeSeq)
				if localReadSeq < currentWrite {
					localReadSeq++
					idx := localReadSeq & db.mask

					processFn(&db.ring[idx])

					atomic.StoreUint64(&db.readSeq, localReadSeq)
				} else {
					runtime.Gosched()
				}
			}
		}
	}()
}
