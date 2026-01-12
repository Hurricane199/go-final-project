package api

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Hurricane199/go-final-project/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return errors.New("Неправильный формат даты, должен быть ГГГГММДД")
	}

	if task.Repeat != "" {
		_, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if !t.After(now) {
		task.Date = now.Format(dateFormat)
	}

	return nil
}

func isRepeatValid(repeat string) bool {
	if repeat == "" {
		return true
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return false
	}

	switch parts[0] {
	case "y":
		return len(parts) == 1

	case "d":
		if len(parts) != 2 {
			return false
		}

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return false
		}

		return n >= 1 && n <= 400
	}

	return false
}
