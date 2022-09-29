package detector

import "time"

type Pattern struct {
	Type           string
	Pair           string
	TrendDirection string
	Start          time.Time
	End            time.Time
}
