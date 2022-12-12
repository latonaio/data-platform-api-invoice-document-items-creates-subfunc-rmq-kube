package requests

type ItemOrdersItem struct {
	OrderID                        *int     `json:"OrderID"`
	OrderItem                      *int     `json:"OrderItem"`
	OrderItemCategory              *string  `json:"OrderItemCategory"`
	OrderItemText                  *string  `json:"OrderItemText"`
	OrderItemTextByBuyer           *string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller          *string  `json:"OrderItemTextBySeller"`
	Product                        *string  `json:"Product"`
	ProductStandardID              *string  `json:"ProductStandardID"`
	ProductGroup                   *string  `json:"ProductGroup"`
	BaseUnit                       *string  `json:"BaseUnit"`
	PricingDate                    *string  `json:"PricingDate"`
	OrderQuantityInBaseUnit        *float32 `json:"OrderQuantityInBaseUnit"`
	ItemWeightUnit                 *string  `json:"ItemWeightUnit"`
	ItemGrossWeight                *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                  *float32 `json:"ItemNetWeight"`
	NetAmount                      *float32 `json:"NetAmount"`
	TaxAmount                      *float32 `json:"TaxAmount"`
	GrossAmount                    *float32 `json:"GrossAmount"`
	ProductionPlantPartnerFunction *string  `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation *string  `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction    *string  `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner    *int     `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                   *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation    *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction  *string  `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner  *int     `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                 *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation  *string  `json:"ReceivingPlantStorageLocation"`
	Incoterms                      *string  `json:"Incoterms"`
	ProductTaxClassification       *string  `json:"ProductTaxClassification"`
	PaymentTerms                   *string  `json:"PaymentTerms"`
	PaymentMethod                  *string  `json:"PaymentMethod"`
	Project                        *string  `json:"Project"`
	TaxCode                        *string  `json:"TaxCode"`
	TaxRate                        *float32 `json:"TaxRate"`
	CountryOfOrigin                *string  `json:"CountryOfOrigin"`
}
