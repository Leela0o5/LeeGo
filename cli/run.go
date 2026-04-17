package cli

import (
	"fmt"
	"os"

	"github.com/Leela0o5/WebSocket-Load-Tester/config"
	"github.com/Leela0o5/WebSocket-Load-Tester/engine"
	"github.com/Leela0o5/WebSocket-Load-Tester/reporter"
	"github.com/spf13/cobra"
)

var (
	cfgPath    string
	outputPath string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start a WebSocket load test",
	Long:  `Executes a load test against the configured WebSocket URL using parameters from the config file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Auto-discover config.yaml in the working directory when no flag is given.
		path := cfgPath
		if path == "" {
			if _, err := os.Stat("config.yaml"); err == nil {
				path = "config.yaml"
			}
		}

		cfg, err := config.LoadConfig(path)
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}

		fmt.Println("WebSocket Load Testing CLI Tool")
		fmt.Printf("Target:      %s\n", cfg.URL)
		fmt.Printf("Connections: %d\n", cfg.NumWorkers)
		fmt.Printf("Duration:    %v\n", cfg.Duration)
		fmt.Printf("Rate:        %d req/s\n", cfg.RateLimit)
		fmt.Printf("Message:     %s\n", cfg.Message)
		fmt.Println("---------------------------------")

		stats := engine.Run(*cfg)
		reporter.PrintSummary(stats)

		if outputPath != "" {
			if err := reporter.SaveJSON(stats, outputPath); err != nil {
				return fmt.Errorf("failed to save report: %w", err)
			}
			fmt.Printf("Report saved to %s\n", outputPath)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "path to YAML config file")
	runCmd.Flags().StringVarP(&outputPath, "output", "o", "", "path to save JSON report")
}
