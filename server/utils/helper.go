package utils

import (
	"fmt"
	"strings"
	"time"
)

// FormatDuration 将时间间隔转换为人类可读的格式
func FormatDuration(d time.Duration) string {
	d = d.Round(time.Millisecond)

	if d < time.Second {
		return d.String()
	}

	parts := []string{}

	if h := d.Hours(); h >= 1 {
		hours := int(h)
		d -= time.Duration(hours) * time.Hour
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}

	if m := d.Minutes(); m >= 1 {
		minutes := int(m)
		d -= time.Duration(minutes) * time.Minute
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}

	if s := d.Seconds(); s >= 1 {
		seconds := int(s)
		parts = append(parts, fmt.Sprintf("%ds", seconds))
	}

	if len(parts) == 0 {
		return "0s"
	}

	return strings.Join(parts, "")
}
