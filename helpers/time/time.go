package time

import "time"

var (
	StandardFormat     = "2006-01-02 15:04:05"
	NotificationFormat = "02/01/2006 - 15:04PM"
)

func GetNow() string {
	return time.Now().Format(StandardFormat)
}

func TimeNotificationFormat(t time.Time) string {
	return t.Format(NotificationFormat)
}

func StdTimeFormat(t time.Time) string {
	return t.Format(StandardFormat)
}

func Parse(layout, value string) (time.Time, error) {
	tm, err := time.Parse(layout, value)
	if err != nil {
		return tm, err
	}
	return tm, nil
}
