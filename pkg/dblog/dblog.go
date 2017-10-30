package dblog

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	sqlRegexp = regexp.MustCompile(`(\$\d+)|\?`)
)

// New initializes a new GORM logger instance.
func New(logger log.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Logger provides the implementation for a GORM logger.
type Logger struct {
	logger log.Logger
}

// Print provides the implementation for a GORM logger.
func (l *Logger) Print(values ...interface{}) {
	if len(values) == 1 {
		level.Error(l.logger).Log(
			"msg", fmt.Sprint(values...),
		)

		return
	}

	var (
		lev    = values[0]
		source = values[1]
	)

	if lev != "sql" {
		level.Error(l.logger).Log(
			"msg", fmt.Sprint(values[2:]...),
			"level", lev,
			"source", source,
		)

		return
	}

	var (
		formattedValues []interface{}
		duration        = values[2]
	)

	for _, value := range values[4].([]interface{}) {
		indirectValue := reflect.Indirect(reflect.ValueOf(value))

		if indirectValue.IsValid() {
			value = indirectValue.Interface()

			if t, ok := value.(time.Time); ok {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format(time.RFC3339)))
			} else if b, ok := value.([]byte); ok {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", string(b)))
			} else if r, ok := value.(driver.Valuer); ok {
				if value, err := r.Value(); err == nil && value != nil {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			} else {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
			}
		} else {
			formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
		}
	}

	level.Debug(l.logger).Log(
		"msg", fmt.Sprintf(sqlRegexp.ReplaceAllString(values[3].(string), "%v"), formattedValues...),
		"duration", duration,
		"level", lev,
		"source", source,
	)
}
