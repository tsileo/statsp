package statsp

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

var nl = []byte{0x0A}

func Parse(b []byte) (metrics *[]Metric, e error) {
	split := bytes.Split(bytes.Trim(b, "\n"), nl)
	r := make([]Metric, len(split))
	re := regexp.MustCompile("^([^:]+):(([\\-\\+]?)[0-9\\..]+)\\|(c|ms|g|h|s)(?:\\|@([0-9\\.]+))?")
	m := Metric{}

	for i, line := range split {
		res := re.FindStringSubmatch(string(line))
		if res == nil {
			return nil, fmt.Errorf("Invalid statsd packet")
		}

		m.Relative = (res[3] != "")

		if t, ok := metric_type_lookup[res[4]]; ok {
			m.Type = t
		} else {
			return nil, fmt.Errorf("Unknown statsd type")
		}
		m.Name = res[1]
		m.Value, e = strconv.ParseFloat(res[2], 64)
		if e != nil {
			return nil, e
		}
		if res[5] != "" {
			if m.Type != Counter {
				return nil, fmt.Errorf("Invalid statsd packet")
			}
			m.SampleRate, e = strconv.ParseFloat(res[5], 64)
			if e != nil {
				return nil, e
			}
		}
		r[i] = m
	}

	return &r, nil

}
