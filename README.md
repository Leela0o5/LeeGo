# WebSocket Load Tester

This is Version 1 of my project. It is a fast tool to test how many messages a WebSocket server can handle. It uses Go's parallel workers to send many messages at once.

## Performance
- 60,000+ requests per second.
- Very low latency tracking.

## How to run
1. Start the test server in one terminal:
   go run cmd/test_server.go
2. Start the load tester in another terminal:
   go run main.go

## Project Roadmap

### Version 1 — Core Load Engine (Done)
**What it does:** Simulates many WebSocket clients at once to measure server performance under heavy load.

**Features:**
- Parallel workers using Go routines
- Each worker keeps its own connection open
- Fast loop for sending and receiving messages
- Results gathered through channels
- Measures average and p95 latency
- Tracks errors and shuts down safely
- Very high speed (60k+ requests/sec)

### Version 2 — Control & Accuracy (Next)
**What it does:** Adds control over the speed of the test and makes it more accurate.

**Features:**
- Rate limiting (control requests per second)
- Config file support (YAML)
- Settings for connections, time, and speed
- Better math for percentiles
- Support for custom messages

### Version 3 — CLI & Developer Experience
**What it does:** Turns the tool into a professional CLI utility.

**Features:**
- Built with Cobra
- New commands to run tests or report on past results
- Export results as JSON
- Cleaner terminal output

### Version 4 — Visualization & UI
**What it does:** Shows test results in a visual dashboard.

**Features:**
- React-based dashboard
- Charts for latency and speed over time
- Error rate tracking
- Real-time updates (optional)

## Tests
Run this command to test the code:
go test ./tests/...
