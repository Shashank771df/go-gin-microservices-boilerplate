package date

import (
	"app/core/utils"
	"errors"
	"time"
)

type DateInformation struct {
	Original  time.Time
	Date      string
	DateTime  string
	TimeStamp int64
	Error     error
}

type TimeLapseInformation struct {
	Label     string
	Duration  int64
	DurationF float64
	Error     error
}

type UtilDate struct {
}

func (obj UtilDate) CurrentTime() DateInformation {
	now := obj.Now()

	return DateInformation{
		Error:     nil,
		Original:  now,
		Date:      obj.Format(now, "2006-01-02"),
		DateTime:  obj.FormatMilliseconds(now),
		TimeStamp: now.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)),
	}
}

func (obj UtilDate) CurrentTimeUTC() DateInformation {
	now := obj.Now().UTC()

	return DateInformation{
		Error:     nil,
		Original:  now,
		Date:      obj.Format(now, "2006-01-02"),
		DateTime:  obj.FormatMilliseconds(now),
		TimeStamp: now.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)),
	}
}

func (obj UtilDate) Now() time.Time {
	return time.Now()
}

func (obj UtilDate) ExecTimeLapse(start time.Time, label string) TimeLapseInformation {
	if start.IsZero() {
		return TimeLapseInformation{Error: errors.New("invalid date")}
	}

	duration := time.Since(start)
	var DurationVal int64 = 0
	var DurationValF float64 = 0

	switch label {
	case "ms":
		DurationVal = duration.Milliseconds()
		DurationValF = float64(DurationVal)
	case "us":
		DurationVal = duration.Microseconds()
		DurationValF = float64(DurationVal)
	case "ns":
		DurationVal = duration.Nanoseconds()
		DurationValF = float64(DurationVal)
	case "s":
		DurationValF = duration.Seconds()
		DurationVal = int64(DurationValF)
	case "h":
		DurationValF = duration.Hours()
		DurationVal = int64(DurationValF)
	default:
		return TimeLapseInformation{Error: errors.New("unknow period label")}
	}

	var util utils.Util
	DurationValF = util.Round(DurationValF, 4)

	return TimeLapseInformation{
		Error:     nil,
		Label:     label,
		Duration:  DurationVal,
		DurationF: DurationValF,
	}
}

func (obj UtilDate) TimeLapseMS(start time.Time, end time.Time) TimeLapseInformation {
	if start.IsZero() {
		return TimeLapseInformation{Error: errors.New("invalid date")}
	}

	duration := end.Sub(start)

	return TimeLapseInformation{
		Error:    nil,
		Label:    "ms",
		Duration: duration.Milliseconds(),
	}
}

func (obj UtilDate) ExecTimeLapseMS(start time.Time) TimeLapseInformation {
	if start.IsZero() {
		return TimeLapseInformation{Error: errors.New("invalid date")}
	}

	duration := time.Since(start)

	return TimeLapseInformation{
		Error:    nil,
		Label:    "ms",
		Duration: duration.Milliseconds(),
	}
}

func (obj UtilDate) AddMinutes(date time.Time, value int) DateInformation {
	return obj.add(date, time.Duration(value)*time.Minute)
}

func (obj UtilDate) AddHours(date time.Time, value int) DateInformation {
	return obj.add(date, time.Duration(value)*time.Hour)
}

func (obj UtilDate) AddDays(date time.Time, value int) DateInformation {
	return obj.addDate(date, value, 0, 0)
}

func (obj UtilDate) AddMonths(date time.Time, value int) DateInformation {
	return obj.addDate(date, 0, value, 0)
}

func (obj UtilDate) AddYears(date time.Time, value int) DateInformation {
	return obj.addDate(date, 0, 0, value)
}

func (obj UtilDate) Week(date time.Time) (int, error) {
	if date.IsZero() {
		return 0, errors.New("invalid date")
	}

	_, week := date.ISOWeek()

	return week, nil
}

func (obj UtilDate) add(date time.Time, period time.Duration) DateInformation {
	finaldate := date

	if date.IsZero() {
		return DateInformation{Error: errors.New("invalid date")}
	}

	finaldate = date.Add(period)

	return DateInformation{
		Error:     nil,
		Original:  date,
		Date:      obj.Format(finaldate, "2006-01-02"),
		DateTime:  obj.FormatMilliseconds(finaldate),
		TimeStamp: finaldate.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)),
	}
}

func (obj UtilDate) addDate(date time.Time, years int, months int, days int) DateInformation {
	finaldate := date

	if date.IsZero() {
		return DateInformation{Error: errors.New("invalid date")}
	}

	finaldate = date.AddDate(years, months, days)

	return DateInformation{
		Error:     nil,
		Original:  date,
		Date:      obj.Format(finaldate, "2006-01-02"),
		DateTime:  obj.FormatMilliseconds(finaldate),
		TimeStamp: finaldate.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)),
	}
}

func (obj UtilDate) FormatMilliseconds(date time.Time) string {
	return obj.Format(date, "2006-01-02 15:04:05.000")
}

// 2006-01-02 15:04:05.000 = yyyy-MM-dd HH:mm:ss.zzz(250)
// 2006-01-02 03:04:05.999 = yyyy-MM-dd hh:mm:ss.zzz(25)
func (obj *UtilDate) Format(date time.Time, format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05.000"
	}

	var str string

	if !date.IsZero() {
		str = date.Format(format)
	} else {
		str = time.Now().Format(format)
	}

	return str
}

func (obj *UtilDate) DateFromString(d string) (time.Time, error) {
	var format string = "2006-01-02"
	var err error
	var parsedDate time.Time

	parsedDate, err = time.Parse(format, d)

	return parsedDate, err
}

func (obj *UtilDate) TimeFromString(t string, withoutSeconds bool) (time.Time, error) {
	var format string = "15:04:05"
	var err error
	var parsedDate time.Time

	if withoutSeconds {
		format = "15:04"
	}

	parsedDate, err = time.Parse(format, t)

	return parsedDate, err
}

func (obj *UtilDate) Difference(startDate time.Time, endDate time.Time) time.Duration {
	return endDate.Sub(startDate)
}

func (obj UtilDate) TimeUnixNano() int64 {
	return time.Now().UnixNano()
}
func (obj *UtilDate) DifferenceTimestamp(startTimestamp int64, endTimestamp int64) int64 {
	return endTimestamp - startTimestamp
}
