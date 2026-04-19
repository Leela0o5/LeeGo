package metrics

import (
	"sync"
	"time"
)

type Stats struct {
	sync.RWMutex
	TotalRequests int
	SuccessCount  int
	FailureCount  int
	Hist          *Histogram
}

func NewStats() *Stats {
	return &Stats{
		Hist: NewHistogram(),
	}
}

func (s *Stats) RecordSuccess(d time.Duration) {
	s.Lock()
	defer s.Unlock()
	s.TotalRequests++
	s.SuccessCount++
	s.Hist.Record(d)
}

func (s *Stats) RecordFailure() {
	s.Lock()
	defer s.Unlock()
	s.TotalRequests++
	s.FailureCount++
}

func (s *Stats) Average() time.Duration {
	s.RLock()
	defer s.RUnlock()
	if s.Hist != nil && s.Hist.Count() > 0 {
		return s.Hist.Mean()
	}
	return 0
}

func (s *Stats) Min() time.Duration {
	s.RLock()
	defer s.RUnlock()
	if s.Hist != nil && s.Hist.Count() > 0 {
		return s.Hist.Min()
	}
	return 0
}

func (s *Stats) Max() time.Duration {
	s.RLock()
	defer s.RUnlock()
	if s.Hist != nil && s.Hist.Count() > 0 {
		return s.Hist.Max()
	}
	return 0
}

func (s *Stats) Median() time.Duration { return s.Percentile(0.50) }
func (s *Stats) P95() time.Duration    { return s.Percentile(0.95) }
func (s *Stats) P99() time.Duration    { return s.Percentile(0.99) }

func (s *Stats) Percentile(p float64) time.Duration {
	s.RLock()
	defer s.RUnlock()
	if s.Hist != nil && s.Hist.Count() > 0 {
		return s.Hist.Percentile(p)
	}
	return 0
}
