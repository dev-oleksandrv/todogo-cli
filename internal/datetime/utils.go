package datetime

import "time"

const (
	dateTimeFormat = "2006-01-02 15:04:05"
)

func GetFormattedTime(t time.Time) string {
	return t.Format(dateTimeFormat)
}

func GetFormattedTimeNow() string {
	return GetFormattedTime(time.Now())
}

func ParseFormattedTime(t string) (*time.Time, error) {
	r, err := time.Parse(dateTimeFormat, t)
	if err != nil {
		return nil, err
	}
	return &r, nil
}