package statsp

type Cleaner struct {
	values map[string]float64
}

func NewCleaner() Cleaner {
	return Cleaner{map[string]float64{}}
}

func (c Cleaner) Clean(m Metric) Metric {
	switch m.Type {
	case Guage:
		if m.Relative {
			m.Value = c.values[m.Name] + m.Value
			c.values[m.Name] = m.Value
			m.Relative = false
		}
	default:
		m.Relative = false
	}
	return m
}

func (c Cleaner) CleanMetrics(metrics []Metric) []Metric {
	r := make([]Metric, len(metrics))
	for i, m := range metrics {
		r[i] = c.Clean(m)
	}
	return r
}
