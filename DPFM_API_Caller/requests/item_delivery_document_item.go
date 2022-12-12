package requests

type ItemDeliveryDocumentItem struct {
	DeliveryDocument               *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem           *int     `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory   *string  `json:"DeliveryDocumentItemCategory"`
	DeliveryDocumentItemText       *string  `json:"DeliveryDocumentItemText"`
	Product                        *string  `json:"Product"`
	ProductStandardID              *string  `json:"ProductStandardID"`
	ProductGroup                   *string  `json:"ProductGroup"`
	BaseUnit                       *string  `json:"BaseUnit"`
	ActualGoodsIssueDate           *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime           *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate         *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime         *string  `json:"ActualGoodsReceiptTime"`
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
	ItemGrossWeight                *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                  *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                 *string  `json:"ItemWeightUnit"`
	OrderID                        *int     `json:"OrderID"`
	OrderItem                      *int     `json:"OrderItem"`
	OrderType                      *string  `json:"OrderType"`
	ContractType                   *string  `json:"ContractType"`
	OrderValidityStartDate         *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate           *string  `json:"OrderValidityEndDate"`
	PaymentTerms                   *string  `json:"PaymentTerms"`
	PaymentMethod                  *string  `json:"PaymentMethod"`
	InvoicePeriodStartDate         *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate           *string  `json:"InvoicePeriodEndDate"`
	Project                        *string  `json:"Project"`
	ProductTaxClassification       *string  `json:"ProductTaxClassification"`
	TaxCode                        *string  `json:"TaxCode"`
	TaxRate                        *float32 `json:"TaxRate"`
	CountryOfOrigin                *string  `json:"CountryOfOrigin"`
}
