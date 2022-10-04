package detector

type Trend int16

const (
	Continuation Trend = 1
	Bearish            = 2
	Bullish            = 3
)

func (t Trend) Icon() string {
	switch t {
	case Bullish:
		return "📈"
	case Bearish:
		return "📉"
	default:
		return ""
	}
}

func (t Trend) Label() string {
	switch t {
	case Bullish:
		return "bullish"
	case Bearish:
		return "bearish"
	default:
		return ""
	}
}
