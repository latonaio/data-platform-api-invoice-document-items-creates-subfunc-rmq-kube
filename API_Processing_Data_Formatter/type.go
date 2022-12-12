package api_processing_data_formatter

type SDC struct {
	MetaData                 *MetaData                   `json:"MetaData"`
	OrderID                  *[]OrderID                  `json:"OrderID"`
	OrderItem                *[]OrderItem                `json:"OrderItem"`
	DeliveryDocument         *[]DeliveryDocument         `json:"DeliveryDocument"`
	DeliveryDocumentItem     *[]DeliveryDocumentItem     `json:"DeliveryDocumentItem"`
	CalculateInvoiceDocument *CalculateInvoiceDocument   `json:"CalculateInvoiceDocument"`
	ItemOrdersItem           *[]ItemOrdersItem           `json:"ItemOrdersItem"`
	ItemDeliveryDocumentItem *[]ItemDeliveryDocumentItem `json:"ItemDeliveryDocumentItem"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderIDKey struct {
	OrderID                         *int   `json:"OrderID"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   []*int `json:"BillFromParty"`
	BillToParty                     []*int `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type OrderID struct {
	InvoiceDocument                 *int   `json:"InvoiceDocument"`
	OrderID                         *int   `json:"OrderID"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   *int   `json:"BillFromParty"`
	BillToParty                     *int   `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type OrdersHeaderPartner struct {
	InvoiceDocument *int   `json:"InvoiceDocument"`
	OrderID         int    `json:"OrderID"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
}

type OrderItemKey struct {
	InvoiceDocument               *int   `json:"InvoiceDocument"`
	OrderID                       []*int `json:"OrderID"`
	ItemCompleteDeliveryIsDefined *bool  `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string `json:"ItemDeliveryStatus"`
	ItemBillingStatus             string `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        *bool  `json:"ItemBillingBlockStatus"`
}

type OrderItem struct {
	InvoiceDocument               *int   `json:"InvoiceDocument"`
	OrderID                       int    `json:"OrderID"`
	ItemCompleteDeliveryIsDefined *bool  `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string `json:"ItemDeliveryStatus"`
	ItemBillingStatus             string `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        *bool  `json:"ItemBillingBlockStatus"`
	OrderItem                     int    `json:"OrderItem"`
}

type DeliveryDocumentKey struct {
	DeliveryDocument                *int   `json:"DeliveryDocument"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   []*int `json:"BillFromParty"`
	BillToParty                     []*int `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type DeliveryDocument struct {
	InvoiceDocument                 *int   `json:"InvoiceDocument"`
	DeliveryDocument                *int   `json:"DeliveryDocument"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   *int   `json:"BillFromParty"`
	BillToParty                     *int   `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type DeliveryDocumentItemKey struct {
	DeliveryDocument              []*int  `json:"DeliveryDocument"`
	ConfirmedDeliveryDateFrom     *string `json:"ConfirmedDeliveryDateFrom"`
	ConfirmedDeliveryDateTo       *string `json:"ConfirmedDeliveryDateTo"`
	ActualGoodsIssueDateFrom      *string `json:"ActualGoodsIssueDateFrom"`
	ActualGoodsIssueDateTo        *string `json:"ActualGoodsIssueDateTo"`
	ConfirmedDeliveryDate         *string `json:"ConfirmedDeliveryDate"`
	ActualGoodsIssueDate          *string `json:"ActualGoodsIssueDate"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string  `json:"ItemDeliveryStatus"`
	ItemBillingStatus             string  `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        *bool   `json:"ItemBillingBlockStatus"`
}

type DeliveryDocumentItem struct {
	InvoiceDocument               *int    `json:"InvoiceDocument"`
	DeliveryDocument              int     `json:"DeliveryDocument"`
	ConfirmedDeliveryDateFrom     *string `json:"ConfirmedDeliveryDateFrom"`
	ConfirmedDeliveryDateTo       *string `json:"ConfirmedDeliveryDateTo"`
	ActualGoodsIssueDateFrom      *string `json:"ActualGoodsIssueDateFrom"`
	ActualGoodsIssueDateTo        *string `json:"ActualGoodsIssueDateTo"`
	ConfirmedDeliveryDate         *string `json:"ConfirmedDeliveryDate"`
	ActualGoodsIssueDate          *string `json:"ActualGoodsIssueDate"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string  `json:"ItemDeliveryStatus"`
	ItemBillingStatus             string  `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        *bool   `json:"ItemBillingBlockStatus"`
	DeliveryDocumentItem          int     `json:"DeliveryDocumentItem"`
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
