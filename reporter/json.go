package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Leela0o5/WebSocket-Load-Tester/metrics"
)

type JSONReport struct {
	Total      int           `json:"total"`
	Errors     int           `json:"errors"`
	AvgLatency time.Duration `json:"avg_latency"`
	P95        time.Duration `json:"p95"`
}

func SaveJSON(s *metrics.Stats, path string) error {
	report := JSONReport{
		Total:      s.TotalRequests,
		Errors:     s.FailureCount,
		AvgLatency: s.Average(),
		P95:        s.P95(),
	}

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("serialize report: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write report file: %w", err)
	}

	return nil
}

func LoadJSON(path string) (*JSONReport, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var report JSONReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, err
	}

	return &report, nil
}
