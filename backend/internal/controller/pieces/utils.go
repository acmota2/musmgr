package pieces

import (
	"time"
)

func formatPieceDate(s string) (time.Time, error) {
	return time.Parse("2006", s)
}
