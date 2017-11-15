package TimeStrings

type TimeStringAggregator interface {
	PrintSum(times []string) string
}

type TimeString struct {}

func (t *TimeString) PrintSum(times []string) string {
	return "0:0:0"
}
