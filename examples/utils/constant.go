package utils

type Interval string

type IntervalPath string

const (
	//Interval1s  Interval = "1s"  // 1 second
	Interval1m  Interval = "1m"  // 1 minute
	Interval3m  Interval = "3m"  // 3 minutes
	Interval5m  Interval = "5m"  // 5 minutes
	Interval15m Interval = "15m" // 15 minutes
	Interval30m Interval = "30m" // 30 minutes
	Interval1h  Interval = "1h"  // 1 hour
	Interval2h  Interval = "2h"  // 2 hours
	Interval4h  Interval = "4h"  // 4 hours
	Interval6h  Interval = "6h"  // 6 hours
	Interval8h  Interval = "8h"  // 8 hours
	Interval12h Interval = "12h" // 12 hours
	Interval1d  Interval = "1d"  // 1 day
	Interval3d  Interval = "3d"  // 3 days
	Interval1w  Interval = "1w"  // 1 week
	Interval1M  Interval = "1M"  // 1 month
)

const (
	IntervalPathMinute IntervalPath = "minute" // 1 minute
	IntervalPathHour   IntervalPath = "hour"   // 1 hour
	IntervalPathDay    IntervalPath = "day"    // 1 day
	IntervalPathWeek   IntervalPath = "week"   // 1 week
	IntervalPathMonth  IntervalPath = "month"  // 1 month
)

var IntervalPathMap = map[Interval]IntervalPath{
	Interval1m:  IntervalPathMinute,
	Interval3m:  IntervalPathMinute,
	Interval5m:  IntervalPathMinute,
	Interval15m: IntervalPathMinute,
	Interval30m: IntervalPathMinute,
	Interval1h:  IntervalPathHour,
	Interval2h:  IntervalPathHour,
	Interval4h:  IntervalPathHour,
	Interval6h:  IntervalPathHour,
	Interval8h:  IntervalPathHour,
	Interval12h: IntervalPathHour,
	Interval1d:  IntervalPathDay,
	Interval3d:  IntervalPathDay,
	Interval1w:  IntervalPathWeek,
	Interval1M:  IntervalPathMonth,
}

var IntervalList = []Interval{
	Interval1m,
	Interval3m,
	Interval5m,
	Interval15m,
	Interval30m,
	Interval1h,
	Interval2h,
	Interval4h,
	Interval6h,
	Interval8h,
	Interval12h,
	Interval1d,
	Interval3d,
	Interval1w,
	Interval1M,
}
