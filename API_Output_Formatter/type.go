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
	Item *[]Item `json:"Item"`
}

type Header struct {
	InvoiceDocument                   int      `json:"InvoiceDocument"`
	CreationDate                      *string  `json:"CreationDate"`
	CreationTime                      *string  `json:"CreationTime"`
	LastChangeDate                    *string  `json:"LastChangeDate"`
	LastChangeTime                    *string  `json:"LastChangeTime"`
	SupplyChainRelationshipID         *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID  *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID  *int     `json:"SupplyChainRelationshipPaymentID"`
	BillToParty                       *int     `json:"BillToParty"`
	BillFromParty                     *int     `json:"BillFromParty"`
	BillToCountry                     *string  `json:"BillToCountry"`
	BillFromCountry                   *string  `json:"BillFromCountry"`
	Payer                             *int     `json:"Payer"`
	Payee                             *int     `json:"Payee"`
	InvoiceDocumentDate               *string  `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               *string  `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              *string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string  `json:"AccountingPostingDate"`
	IsExportImport                    *bool    `json:"IsExportImport"`
	HeaderBillingIsConfirmed          *bool    `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus           *string  `json:"HeaderBillingConfStatus"`
	TotalNetAmount                    *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32 `json:"TotalGrossAmount"`
	TransactionCurrency               *string  `json:"TransactionCurrency"`
	Incoterms                         *string  `json:"Incoterms"`
	PaymentTerms                      *string  `json:"PaymentTerms"`
	DueCalculationBaseDate            *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string  `json:"PaymentDueDate"`
	NetPaymentDays                    *int     `json:"NetPaymentDays"`
	PaymentMethod                     *string  `json:"PaymentMethod"`
	ExternalReferenceDocument         *string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText                *string  `json:"DocumentHeaderText"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
	InvoiceDocumentIsCancelled        *bool    `json:"InvoiceDocumentIsCancelled"`
	CancelledInvoiceDocument          *int     `json:"CancelledInvoiceDocument"`
}

type Partner struct {
	InvoiceDocument         int     `json:"InvoiceDocument"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type HeaderDoc struct {
	InvoiceDocument          int     `json:"InvoiceDocument"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Address struct {
	InvoiceDocument int     `json:"InvoiceDocument"`
	AddressID       int     `json:"AddressID"`
	PostalCode      *string `json:"PostalCode"`
	LocalRegion     *string `json:"LocalRegion"`
	Country         *string `json:"Country"`
	District        *string `json:"District"`
	StreetName      *string `json:"StreetName"`
	CityName        *string `json:"CityName"`
	Building        *string `json:"Building"`
	Floor           *int    `json:"Floor"`
	Room            *int    `json:"Room"`
}

type Item struct {
	InvoiceDocumentItem                     *int     `json:"InvoiceDocumentItem"`
	InvoiceDocumentItemCategory             *string  `json:"InvoiceDocumentItemCategory"`
	SupplyChainRelationshipID               *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	InvoiceDocumentItemText                 *string  `json:"InvoiceDocumentItemText"`
	InvoiceDocumentItemTextByBuyer          *string  `json:"InvoiceDocumentItemTextByBuyer"`
	InvoiceDocumentItemTextBySeller         *string  `json:"InvoiceDocumentItemTextBySeller"`
	Product                                 *string  `json:"Product"`
	ProductGroup                            *string  `json:"ProductGroup"`
	ProductStandardID                       *string  `json:"ProductStandardID"`
	CreationDate                            *string  `json:"CreationDate"`
	CreationTime                            *string  `json:"CreationTime"`
	LastChangeDate                          *string  `json:"LastChangeDate"`
	LastChangeTime                          *string  `json:"LastChangeTime"`
	ItemBillingIsConfirmed                  *bool    `json:"ItemBillingIsConfirmed"`
	Buyer                                   *int     `json:"Buyer"`
	Seller                                  *int     `json:"Seller"`
	DeliverToParty                          *int     `json:"DeliverToParty"`
	DeliverFromParty                        *int     `json:"DeliverFromParty"`
	DeliverToPlant                          *string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation           *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation         *string  `json:"DeliverFromPlantStorageLocation"`
	ProductionPlantBusinessPartner          *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                         *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation          *string  `json:"ProductionPlantStorageLocation"`
	ServicesRenderedDate                    *string  `json:"ServicesRenderedDate"`
	InvoiceQuantity                         *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit                     *string  `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit               *float32 `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                                *string  `json:"BaseUnit"`
	ActualGoodsIssueDate                    *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                    *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                  *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                  *string  `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                         *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                           *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                          *string  `json:"ItemWeightUnit"`
	NetAmount                               *float32 `json:"NetAmount"`
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	GoodsIssueOrReceiptSlipNumber           *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency                     *string  `json:"TransactionCurrency"`
	PricingDate                             *string  `json:"PricingDate"`
	TransactionTaxClassification            *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                *string  `json:"DefinedTaxClassification"`
	Project                                 *string  `json:"Project"`
	OrderID                                 *int     `json:"OrderID"`
	OrderItem                               *int     `json:"OrderItem"`
	OrderType                               *string  `json:"OrderType"`
	ContractType                            *string  `json:"ContractType"`
	OrderVaridityStartDate                  *string  `json:"OrderVaridityStartDate"`
	OrderValidityEndDate                    *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate                  *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string  `json:"InvoicePeriodEndDate"`
	DeliveryDocument                        *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem                    *int     `json:"DeliveryDocumentItem"`
	OriginDocument                          *int     `json:"OriginDocument"`
	OriginDocumentItem                      *int     `json:"OriginDocumentItem"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	ExternalReferenceDocument               *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem           *string  `json:"ExternalReferenceDocumentItem"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	CountryOfOrigin                         *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 *string  `json:"CountryOfOriginLanguage"`
	ItemPaymentRequisitionIsCreated         *bool    `json:"ItemPaymentRequisitionIsCreated"`
	ItemPaymentBlockStatus                  *bool    `json:"ItemPaymentBlockStatus"`
}

type ItemPricingElement struct {
	InvoiceDocument            *int     `json:"InvoiceDocument"`
	InvoiceDocumentItem        *int     `json:"InvoiceDocumentItem"`
	PricingProcedureCounter    *int     `json:"PricingProcedureCounter"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
