package reporter

import (
	"fmt"
	"github.com/Leela0o5/WebSocket-Load-Tester/metrics"
)

func PrintSummary(s *metrics.Stats) {
	avg := s.Average()
	p95 := s.P95()

	fmt.Println("\n--- BENCHMARK SUMMARY ---")
	fmt.Printf("Total Requests: %d\n", s.TotalRequests)
	fmt.Printf("Errors:         %d\n", s.FailureCount)
	fmt.Printf("Avg Latency:    %v\n", avg)
	fmt.Printf("P95 Latency:    %v\n", p95)
	fmt.Println("-------------------------")
}
