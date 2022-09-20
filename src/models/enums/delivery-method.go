package enums

import "database/sql/driver"

type DeliveryMethod string

const (
	DeliveryMethodDiscord     DeliveryMethod = "DISCORD"
	DeliveryMethodGoogleDrive DeliveryMethod = "GOOGLE_DRIVE"
)

func (p *DeliveryMethod) Scan(value any) error {
	*p = DeliveryMethod(value.([]byte))
	return nil
}

func (p DeliveryMethod) Value() (driver.Value, error) {
	return string(p), nil
}
