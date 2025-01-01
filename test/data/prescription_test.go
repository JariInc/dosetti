package data_test

import (
	"testing"
	"time"

	"github.com/jariinc/dosetti/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestHourlyOccurrancesBetweenDates(t *testing.T) {
	p := data.Prescription{
		Interval:     2,
		IntervalUnit: data.IntervalHourly,
		StartAt:      time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	from := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC)

	occurrances := p.OccurrancesBetweenDates(from, to)

	assert.Len(t, occurrances, 12, "Expected 12 occurrances")
	assert.Equal(t, []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, occurrances)
}

func TestDailyOccurrancesBetweenDates(t *testing.T) {
	p := data.Prescription{
		Interval:     3,
		IntervalUnit: data.IntervalDaily,
		StartAt:      time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	from := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	occurrances := p.OccurrancesBetweenDates(from, to)

	assert.Len(t, occurrances, 10, "Expected 1o occurrances")
	assert.Equal(t, []int{51, 52, 53, 54, 55, 56, 57, 58, 59, 60}, occurrances)
}

func TestWeeklyOccurrancesBetweenDates(t *testing.T) {
	p := data.Prescription{
		Interval:     1,
		IntervalUnit: data.IntervalWeekly,
		StartAt:      time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	from := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC)

	occurrances := p.OccurrancesBetweenDates(from, to)

	assert.Len(t, occurrances, 13, "Expected 13 occurrances")
	assert.Equal(t, []int{22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34}, occurrances)
}
