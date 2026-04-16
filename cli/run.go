package cli

import (
	"fmt"
	"time"
	"github.com/Leela0o5/WebSocket-Load-Tester/config"
	"github.com/Leela0o5/WebSocket-Load-Tester/engine"
	"github.com/Leela0o5/WebSocket-Load-Tester/reporter"
)

func Run() error {
	fmt.Println("WebSocket Load Testing CLI Tool")

	cfg := config.Config{
		URL:        "ws://localhost:8080/ws",
		NumWorkers: 10,
		Duration:   10 * time.Second,
	}

	fmt.Printf("Target:     %s\n", cfg.URL)
	fmt.Printf("Workers:    %d\n", cfg.NumWorkers)
	fmt.Printf("Duration:   %v\n", cfg.Duration)
	fmt.Println("---------------------------------")

	stats := engine.Run(cfg)
	reporter.PrintSummary(stats)

	return nil
}
