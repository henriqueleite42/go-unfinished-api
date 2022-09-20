package enums

import "database/sql/driver"

type DiscountType string

const (
	DiscountTypeRawValue   DiscountType = "RAW_VALUE"
	DiscountTypePercentage DiscountType = "PERCENTAGE"
)

func (p *DiscountType) Scan(value any) error {
	*p = DiscountType(value.([]byte))
	return nil
}

func (p DiscountType) Value() (driver.Value, error) {
	return string(p), nil
}
