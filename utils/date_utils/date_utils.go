package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	// JSTの現在時刻
	now := time.Now().UTC().In(jst)

	return now
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
