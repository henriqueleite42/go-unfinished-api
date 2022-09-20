package enums

import "database/sql/driver"

type SaleStatus string

const (
	SaleStatusPending   SaleStatus = "PENDING"
	SaleStatusPaid      SaleStatus = "PAID"
	SaleStatusDelivered SaleStatus = "DELIVERED"
	SaleStatusCanceled  SaleStatus = "CANCELED"
	SaleStatusRefunded  SaleStatus = "REFUNDED"
)

func (p *SaleStatus) Scan(value any) error {
	*p = SaleStatus(value.([]byte))
	return nil
}

func (p SaleStatus) Value() (driver.Value, error) {
	return string(p), nil
}
