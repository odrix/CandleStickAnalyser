package detectdaily

import "time"

type Pattern struct {
	Type  string
	Pair  string
	Start time.Time
	End   time.Time
}
