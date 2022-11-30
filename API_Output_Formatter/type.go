package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	Message             Message  `json:"message"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	Item []Item `json:"Header"`
}

type Header struct {
	InvoiceDocument            int      `json:"InvoiceDocument"`
	CreationDate               *string  `json:"CreationDate"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	BillToParty                *int     `json:"BillToParty"`
	BillFromParty              *int     `json:"BillFromParty"`
	BillToCountry              *string  `json:"BillToCountry"`
	BillFromCountry            *string  `json:"BillFromCountry"`
	InvoiceDocumentDate        *string  `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime        *string  `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate     *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate       *string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate      *string  `json:"AccountingPostingDate"`
	InvoiceDocumentIsCancelled *bool    `json:"InvoiceDocumentIsCancelled"`
	CancelledInvoiceDocument   *int     `json:"CancelledInvoiceDocument"`
	IsExportImportDelivery     *bool    `json:"IsExportImportDelivery"`
	HeaderBillingIsConfirmed   *bool    `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus    *string  `json:"HeaderBillingConfStatus"`
	TotalNetAmount             *float32 `json:"TotalNetAmount"`
	TotalTaxAmount             *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount           *float32 `json:"TotalGrossAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	Incoterms                  *string  `json:"Incoterms"`
	PaymentTerms               *string  `json:"PaymentTerms"`
	DueCalculationBaseDate     *string  `json:"DueCalculationBaseDate"`
	NetPaymentDays             *int     `json:"NetPaymentDays"`
	PaymentMethod              *string  `json:"PaymentMethod"`
	HeaderPaymentBlockStatus   *bool    `json:"HeaderPaymentBlockStatus"`
	ExternalReferenceDocument  *string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText         *string  `json:"DocumentHeaderText"`
}

type Item struct {
	InvoiceDocument                int      `json:"InvoiceDocument"`
	InvoiceDocumentItem            int      `json:"InvoiceDocumentItem"`
	DocumentItemCategory           *string  `json:"DocumentItemCategory"`
	InvoiceDocumentItemText        *string  `json:"InvoiceDocumentItemText"`
	CreationDate                   *string  `json:"CreationDate"`
	CreationTime                   *string  `json:"CreationTime"`
	ItemBillingIsConfirmed         *bool    `json:"ItemBillingIsConfirmed"`
	ItemBillingConfStatus          *string  `json:"ItemBillingConfStatus"`
	Buyer                          *int     `json:"Buyer"`
	Seller                         *int     `json:"Seller"`
	DeliverToParty                 *int     `json:"DeliverToParty"`
	DeliverFromParty               *int     `json:"DeliverFromParty"`
	ProductStandardID              *string  `json:"ProductStandardID"`
	Batch                          *string  `json:"Batch"`
	Product                        *string  `json:"Product"`
	ProductGroup                   *string  `json:"ProductGroup"`
	ProductionPlantPartnerFunction *string  `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner *string  `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation *string  `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction    *string  `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner    *string  `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                   *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation    *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction  *string  `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner  *string  `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                 *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation  *string  `json:"ReceivingPlantStorageLocation"`
	ServicesRenderedDate           *string  `json:"ServicesRenderedDate"`
	InvoiceQuantity                *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit            *string  `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit      *float32 `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                       *string  `json:"BaseUnit"`
	ActualGoodsIssueDate           *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime           *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate         *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime         *string  `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                  *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                 *string  `json:"ItemWeightUnit"`
	NetAmount                      *float32 `json:"NetAmount"`
	TaxAmount                      *float32 `json:"TaxAmount"`
	GrossAmount                    *float32 `json:"GrossAmount"`
	ItemPaymentBlockStatus         *bool    `json:"ItemPaymentBlockStatus"`
	GoodsIssueOrReceiptSlipNumber  *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency            *string  `json:"TransactionCurrency"`
	PricingDate                    *string  `json:"PricingDate"`
	ProductTaxClassification       *string  `json:"ProductTaxClassification"`
	Project                        *string  `json:"Project"`
	OrderID                        *int     `json:"OrderID"`
	OrderItem                      *int     `json:"OrderItem"`
	OrderType                      *string  `json:"OrderType"`
	ContractType                   *string  `json:"ContractType"`
	OrderVaridityStartDate         *string  `json:"OrderVaridityStartDate"`
	OrderVaridityEndDate           *string  `json:"OrderVaridityEndDate"`
	InvoiceScheduleStartDate       *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         *string  `json:"InvoiceScheduleEndDate"`
	DeliveryDocument               *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem           *int     `json:"DeliveryDocumentItem"`
	OriginDocument                 *int     `json:"OriginDocument"`
	OriginDocumentItem             *int     `json:"OriginDocumentItem"`
	ReferenceDocument              *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem          *int     `json:"ReferenceDocumentItem"`
	ReferenceDocumentType          *string  `json:"ReferenceDocumentType"`
	ExternalReferenceDocument      *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem  *string  `json:"ExternalReferenceDocumentItem"`
	TaxCode                        *string  `json:"TaxCode"`
	TaxRate                        *float32 `json:"TaxRate"`
	CountryOfOrigin                *string  `json:"CountryOfOrigin"`
}
