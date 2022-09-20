package enums

import "database/sql/driver"

type ProductType string

const (
	ProductTypePack            ProductType = "PACK"
	ProductTypeAudio           ProductType = "AUDIO"
	ProductTypeVideo           ProductType = "VIDEO"
	ProductTypePhoto           ProductType = "PHOTO"
	ProductTypeCustomAudio     ProductType = "CUSTOM_AUDIO"
	ProductTypeCustomVideo     ProductType = "CUSTOM_VIDEO"
	ProductTypeCustomPhoto     ProductType = "CUSTOM_PHOTO" // Plaquinha
	ProductTypeRentAGirlfriend ProductType = "RENT_A_GIRLFRIEND"
	ProductTypeSexyVoiceCall   ProductType = "SEXY_VOICE_CALL" // GF/Gozofone
	ProductTypeSexyVideoCall   ProductType = "SEXY_VIDEO_CALL"
)

func (p *ProductType) Scan(value any) error {
	*p = ProductType(value.([]byte))
	return nil
}

func (p ProductType) Value() (driver.Value, error) {
	return string(p), nil
}
