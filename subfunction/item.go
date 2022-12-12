package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) ItemOrdersItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ItemOrdersItem, error) {
	var args []interface{}

	orderItem := psdc.OrderItem

	repeat := strings.Repeat("(?, ?),", len(*orderItem)-1) + "(?, ?)"
	for _, tag := range *orderItem {
		args = append(args, tag.OrderID, tag.OrderItem)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, OrderItemCategory, OrderItemText, OrderItemTextByBuyer, OrderItemTextBySeller,
		Product, ProductStandardID, ProductGroup, BaseUnit, PricingDate, OrderQuantityInBaseUnit, ItemWeightUnit,
		ItemGrossWeight, ItemNetWeight, NetAmount, TaxAmount, GrossAmount, ProductionPlantPartnerFunction,
		ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantStorageLocation, IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner, IssuingPlant, IssuingPlantStorageLocation, ReceivingPlantPartnerFunction, 
		ReceivingPlantBusinessPartner, ReceivingPlant, ReceivingPlantStorageLocation, Incoterms, ProductTaxClassification, 
		PaymentTerms, PaymentMethod, Project, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToItemOrdersItem(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ItemDeliveryDocumentItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.ItemDeliveryDocumentItem, error) {
	var args []interface{}

	deliverydocumentItem := psdc.DeliveryDocumentItem

	repeat := strings.Repeat("(?, ?),", len(*deliverydocumentItem)-1) + "(?, ?)"
	for _, tag := range *deliverydocumentItem {
		args = append(args, tag.DeliveryDocument, tag.DeliveryDocumentItem)
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, DeliveryDocumentItem, DeliveryDocumentItemCategory, DeliveryDocumentItemText,
		Product, ProductStandardID, ProductGroup, BaseUnit, ActualGoodsIssueDate, ActualGoodsIssueTime,
		ActualGoodsReceiptDate, ActualGoodsReceiptTime, ProductionPlantPartnerFunction, ProductionPlantBusinessPartner,
		ProductionPlant, ProductionPlantStorageLocation, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner,
		IssuingPlant, IssuingPlantStorageLocation, ReceivingPlantPartnerFunction, ReceivingPlantBusinessPartner,
		ReceivingPlant, ReceivingPlantStorageLocation, ItemGrossWeight, ItemNetWeight, ItemWeightUnit, OrderID,
		OrderItem, OrderType, ContractType, OrderValidityStartDate, OrderValidityEndDate, PaymentTerms, PaymentMethod,
		InvoicePeriodStartDate, InvoicePeriodEndDate, Project, ProductTaxClassification, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		WHERE (DeliveryDocument, DeliveryDocumentItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToItemDeliveryDocumentItem(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
