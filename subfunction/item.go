package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) OrderItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderItem, error) {
	var args []interface{}

	orderID := psdc.OrderID
	dataKey, err := psdc.ConvertToOrderItemKey(sdc, orderID)
	if err != nil {
		return nil, err
	}

	repeat := strings.Repeat("(?, ?),", len(*dataKey)-1) + "(?, ?)"
	for _, tag := range *dataKey {
		args = append(args, tag.OrderID, tag.ItemInvoiceDocumentStatus)
	}

	rows, err := f.db.Query(
		`SELECT ItemInvoiceDocumentStatus, InvoiceDocumentItem, DocumentItemCategory, CreationDate, CreationTime, Division,
		Customer, Supplier, DeliverToParty, DeliverFromParty, Product, ProductStandardID, Batch, ProductGroup,
		IssuingPlant, IssuingPlantStorageLocation, ReceivingPlant, ReceivingPlantStorageLocation, InvoiceDocumentItemText,
		ServicesRenderedDate, InvoiceQuantity, InvoiceQuantityUnit, InvoiceQuantityInBaseUnit, BaseUnit,ActualGoodsIssueDate,
		ActualGoodsIssueTime, ActualGoodsReceiptDate, ActualGoodsReceiptTime, ItemGrossWeight, ItemNetWeight, ItemWeightUnit, 
		ItemVolume, ItemVolumeUnit, NetAmount, GoodsIssueOrReceiptSlipNumber, TransactionCurrency, BusinessPartnerCurrency,
		GrossAmount, PricingDate, TaxAmount, ProductTaxClassification1, BusinessArea, ProfitCenter, Project, OrderID,OrderItem,
		OrderType,ContractType, OrderVaridityStartDate, OrderValidityEndDate, InvoiceScheduleStartDate ,InvoiceScheduleEndDate,
		DeliveryDocument, DeliveryDocumentItem, OriginDocument, OriginDocumentItem, ReferenceDocument, ReferenceDocumentItem,
		ReferenceDocumentType, ExternalReferenceDocument, ExternalReferenceDocumentItem, OrderSalesOrganization,
		OrderPurchaseOrganization,OrderDistributionChannel, TaxCode, TaxRate, CountryOfOrigin
		WHERE (OrderID, ItemInvoiceDocumentStatus) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderItem(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
