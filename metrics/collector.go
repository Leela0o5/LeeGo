package metrics

func Collector(results <-chan Result) *Stats {
	s := NewStats()

	for res := range results {
		if res.Err != nil {
			s.RecordFailure()
			continue
		}

		s.RecordSuccess(res.Latency)
	}

	return s
}
