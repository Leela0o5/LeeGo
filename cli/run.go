package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Leela0o5/WebSocket-Load-Tester/config"
	"github.com/Leela0o5/WebSocket-Load-Tester/engine"
	"github.com/Leela0o5/WebSocket-Load-Tester/reporter"
)

func Run() error {
	cfgPath := flag.String("config", "", "path to YAML config file (default: config.yaml if present)")
	flag.Parse()

	path := *cfgPath
	if path == "" {
		if _, err := os.Stat("config.yaml"); err == nil {
			path = "config.yaml"
		}
	}

	cfg, err := config.LoadConfig(path)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	fmt.Println("WebSocket Load Testing CLI Tool")
	fmt.Printf("Target:      %s\n", cfg.URL)
	fmt.Printf("Connections: %d\n", cfg.NumWorkers)
	fmt.Printf("Duration:    %v\n", cfg.Duration)
	fmt.Printf("Rate:        %d req/s\n", cfg.RateLimit)
	fmt.Println("---------------------------------")

	stats := engine.Run(*cfg)
	reporter.PrintSummary(stats)

	return nil
}
