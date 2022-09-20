package enums

import "database/sql/driver"

type Period string

const (
	PeriodWeek    Period = "WEEK"
	PeriodMonth   Period = "MONTH"
	PeriodAllTime Period = "ALL_TIME"
)

func (p *Period) Scan(value any) error {
	*p = Period(value.([]byte))
	return nil
}

func (p Period) Value() (driver.Value, error) {
	return string(p), nil
}
