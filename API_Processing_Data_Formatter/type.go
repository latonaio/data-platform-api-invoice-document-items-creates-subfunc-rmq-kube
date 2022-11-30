package api_processing_data_formatter

type SDC struct {
	MetaData                 *MetaData                 `json:"MetaData"`
	OrderID                  *[]OrderID                `json:"OrderID"`
	DeliveryDocument         *[]DeliveryDocument       `json:"DeliveryDocument"`
	CalculateInvoiceDocument *CalculateInvoiceDocument `json:"CalculateInvoiceDocument"`
	OrderItem                *OrderItem                `json:"OrderItem"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderIDKey struct {
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type OrderID struct {
	InvoiceDocument                 *int   `json:"InvoiceDocument"`
	OrderID                         *int   `json:"OrderID`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type DeliveryDocumentKey struct {
	CompleteDeliveryIsDefined *bool  `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus     string `json:"OverallDeliveryStatus"`
}

type DeliveryDocument struct {
	InvoiceDocument           *int   `json:"InvoiceDocument"`
	DeliveryDocument          *int   `json:"DeliveryDocument`
	CompleteDeliveryIsDefined *bool  `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus     string `json:"OverallDeliveryStatus"`
}

type CalculateInvoiceDocumentKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateInvoiceDocumentQueryGets struct {
	ServiceLabel                string `json:"service_label"`
	FieldNameWithNumberRange    string `json:"FieldNameWithNumberRange"`
	InvoiceDocumentLatestNumber *int   `json:"InvoiceDocumentLatestNumber"`
}

type CalculateInvoiceDocument struct {
	InvoiceDocumentLatestNumber *int `json:"InvoiceDocumentLatestNumber"`
	InvoiceDocument             *int `json:"InvoiceDocument"`
}

type OrderItemKey struct {
	OrderID                   *int  `json:"OrderID`
	ItemInvoiceDocumentStatus *bool `json:"ItemInvoiceDocumentStatus"`
}

type OrderItem struct {
	ItemInvoiceDocumentStatus     *bool    `json:"ItemInvoiceDocumentStatus"`
	InvoiceDocumentItem           *int     `json:"InvoiceDocumentItem`
	DocumentItemCategory          *string  `json:"DocumentItemCategory"`
	CreationDate                  *string  `json:"CreationDate"`
	CreationTime                  *string  `json:"CreationTime"`
	Division                      *string  `json:"Division"`
	Customer                      *int     `json:"Customer"`
	Supplier                      *int     `json:"Supplier"`
	DeliverToParty                *int     `json:"DeliverToParty"`
	DeliverFromParty              *int     `json:"DeliverFromParty"`
	Product                       *string  `json:"Product"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	Batch                         *string  `json:"Batch"`
	ProductGroup                  *string  `json:"ProductGroup"`
	IssuingPlant                  *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation   *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlant                *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation *string  `json:"ReceivingPlantStorageLocation"`
	InvoiceDocumentItemText       *string  `json:"InvoiceDocumentItemText"`
	ServicesRenderedDate          *string  `json:"ServicesRenderedDate"`
	InvoiceQuantity               *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit           *string  `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit     *float32 `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                      *string  `json:"BaseUnit`
	ActualGoodsIssueDate          *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime          *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate        *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime        *string  `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight               *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                 *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                *string  `json:"ItemWeightUnit"`
	ItemVolume                    *float32 `json:"ItemVolume"`
	ItemVolumeUnit                *string  `json:"ItemVolumeUnit"`
	NetAmount                     *float32 `json:"NetAmount"`
	GoodsIssueOrReceiptSlipNumber *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency           *string  `json:"TransactionCurrency"`
	BusinessPartnerCurrency       *string  `json:"BusinessPartnerCurrency"`
	GrossAmount                   *float32 `json:"GrossAmount"`
	PricingDate                   *string  `json:"PricingDate"`
	TaxAmount                     *float32 `json:"TaxAmount"`
	ProductTaxClassification1     *string  `json:"ProductTaxClassification1"`
	BusinessArea                  *string  `json:"BusinessArea"`
	ProfitCenter                  *string  `json:"ProfitCenter"`
	Project                       *string  `json:"Project"`
	OrderID                       *int     `json:"OrderID"`
	OrderItem                     *int     `json:"OrderItem`
	OrderType                     *string  `json:"OrderType"`
	ContractType                  *string  `json:"ContractType"`
	OrderVaridityStartDate        *string  `json:"OrderVaridityStartDate"`
	OrderValidityEndDate          *string  `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate      *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate        *string  `json:"InvoiceScheduleEndDate"`
	DeliveryDocument              *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem          *int     `json:"DeliveryDocumentItem"`
	OriginDocument                *int     `json:"OriginDocument"`
	OriginDocumentItem            *int     `json:"OriginDocumentItem"`
	ReferenceDocument             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem         *int     `json:"ReferenceDocumentItem"`
	ReferenceDocumentType         *string  `json:"ReferenceDocumentType"`
	ExternalReferenceDocument     *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem *string  `json:"ExternalReferenceDocumentItem"`
	OrderSalesOrganization        *string  `json:"OrderSalesOrganization"`
	OrderPurchaseOrganization     *string  `json:"OrderPurchaseOrganization"`
	OrderDistributionChannel      *string  `json:"OrderDistributionChannel"`
	TaxCode                       *string  `json:"TaxCode"`
	TaxRate                       *string  `json:"TaxRate"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
}
