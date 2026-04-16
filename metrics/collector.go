package metrics
import (
	"time"
)

func Collector(results <-chan Result) Stats {
	s := Stats{
		Latencies: make([]time.Duration, 0),
	}

	for res := range results {
		s.TotalRequests++

		if res.Err != nil {
			s.FailureCount++
			continue
		}

		s.SuccessCount++
		s.Latencies = append(s.Latencies, res.Latency)
	}

	return s
}