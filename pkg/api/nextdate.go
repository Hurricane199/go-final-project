package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", nil
	}

	start, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", nil
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", nil
	}

	date := start

	switch parts[0] {

	case "d":
		if len(parts) != 2 {
			return "", nil
		}

		n, err := strconv.Atoi(parts[1])
		if err != nil || n < 1 || n > 400 {
			return "", nil
		}

		date = date.AddDate(0, 0, n)

		for !date.After(now) {
			date = date.AddDate(0, 0, n)
		}

	case "y":
		if len(parts) != 1 {
			return "", nil
		}

		y, m, d := date.Date()
		date = time.Date(y+1, m, d, 0, 0, 0, 0, date.Location())

		// 29 февраля
		if m == time.February && d == 29 && date.Month() != time.February {
			date = time.Date(y+1, time.March, 1, 0, 0, 0, 0, date.Location())
		}

		for !date.After(now) {
			y, m, d := date.Date()
			next := time.Date(
				y+1, m, d,
				0, 0, 0, 0,
				date.Location(),
			)

			// 29 февраля
			if m == time.February && d == 29 && next.Month() != time.February {
				next = time.Date(
					y+1, time.March, 1,
					0, 0, 0, 0,
					date.Location(),
				)
			}

			date = next
		}

	default:
		return "", nil
	}

	return date.Format(dateFormat), nil
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")

	var now time.Time
	var err error

	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	next, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, next)
}
