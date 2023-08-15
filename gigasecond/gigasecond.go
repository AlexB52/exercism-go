package gigasecond

import "time"

const GigaSecond time.Duration = 1e9 * time.Second

func AddGigasecond(t time.Time) time.Time {
	return t.Add(GigaSecond)
}
