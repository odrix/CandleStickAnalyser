package detector

import "time"

type Pattern struct {
	Type           string
	Pair           string
	TrendDirection Trend
	Start          time.Time
	End            time.Time
}
