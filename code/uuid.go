package code

import (
	"github.com/google/uuid"
	"math"
	"sync/atomic"
	"time"
)

var defaultUuidSeq uint64

func UUID() string {
	return uuid.New().String()
}

func epoch() time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05.000", "2019-03-15 12:00:00.000")
	return t
}

// @warn it's a simple sequence generator. Please use your own redis sequence generator.
func incrDefaultUuidSeq() uint64 {
	atomic.CompareAndSwapUint64(&defaultUuidSeq, math.MaxUint64, 0)
	return atomic.AddUint64(&defaultUuidSeq, 1)
}

// Uint64Short 支持1秒钟并发1024
func Uint64ShortUUID() uint64 {
	now := time.Now().Unix() - epoch().Unix()
	seq := incrDefaultUuidSeq()
	return ToUint64ShortUUID(now, seq)
}

// ShortUuidToTime 反转short uuid
func ShortUuidToTime(id uint64) (ts int64, ets int64, seq uint64) {
	if ets < 0 {
		ets = 0
	}
	seq = 1023 & id
	ets = int64(id >> 10) // 1024 = 10 bit
	ts = ets + epoch().Unix()
	return
}

func ToUint64ShortUUID(ets int64, seq uint64) uint64 {
	id := uint64(ets) << (64 - 54)
	// 10 bit : 0~1024
	id |= seq % 1024
	return id
}
