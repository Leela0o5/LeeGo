package metrics

import (
	"slices"
	"time"
)


type Stats struct {
	TotalRequests int
	SuccessCount  int
	FailureCount  int
	Latencies     []time.Duration
}

func (s *Stats) Average() time.Duration {
	if len(s.Latencies) == 0 {
		return 0
	}

	var total time.Duration
	for _, l := range s.Latencies {
		total += l
	}

	return total / time.Duration(len(s.Latencies))
}

func (s *Stats) P95() time.Duration {
	return s.Percentile(0.95)
}

func (s *Stats) Median() time.Duration {
	return s.Percentile(0.50)
}

func (s *Stats) Min() time.Duration {
	if len(s.Latencies) == 0 {
		return 0
	}
	s.sortLatencies()
	return s.Latencies[0]
}

func (s *Stats) Max() time.Duration {
	if len(s.Latencies) == 0 {
		return 0
	}
	s.sortLatencies()
	return s.Latencies[len(s.Latencies)-1]
}

func (s *Stats) Percentile(p float64) time.Duration {
	if len(s.Latencies) == 0 {
		return 0
	}

	s.sortLatencies()

	idx := int(float64(len(s.Latencies)) * p)
	if idx >= len(s.Latencies) {
		idx = len(s.Latencies) - 1
	}

	return s.Latencies[idx]
}

func (s *Stats) sortLatencies() {
	slices.Sort(s.Latencies)
}
