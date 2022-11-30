package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	"data-platform-api-invoice-document-headers-creates-subfunc-rmq/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) (*MetaData, error) {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}
	data := pm

	metaData := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &metaData, nil
}

func (psdc *SDC) ConvertToOrderIDKey(sdc *api_input_reader.SDC) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(false),
		OverallDeliveryStatus:           "CL",
	}
	data := pm

	orderIDKey := OrderIDKey{
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		OverallDeliveryStatus:           data.OverallDeliveryStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderID(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID
	pm := &requests.OrderID{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			OrderID:                         data.OrderID,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			OverallDeliveryStatus:           data.OverallDeliveryStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentKey(sdc *api_input_reader.SDC) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		CompleteDeliveryIsDefined: getBoolPtr(false),
		OverallDeliveryStatus:     "CL",
	}
	data := pm

	deliveryDocumentKey := DeliveryDocumentKey{
		CompleteDeliveryIsDefined: data.CompleteDeliveryIsDefined,
		OverallDeliveryStatus:     data.OverallDeliveryStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocument(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument
	pm := &requests.DeliveryDocument{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.CompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			DeliveryDocument:          data.DeliveryDocument,
			CompleteDeliveryIsDefined: data.CompleteDeliveryIsDefined,
			OverallDeliveryStatus:     data.OverallDeliveryStatus,
		})
	}

	return &deliveryDocument, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func (psdc *SDC) ConvertToCalculateInvoiceDocumentKey() (*CalculateInvoiceDocumentKey, error) {
	pm := &requests.CalculateInvoiceDocumentKey{
		ServiceLabel:             "",
		FieldNameWithNumberRange: "InvoiceDocument",
	}
	data := pm

	calculateInvoiceDocumentKey := CalculateInvoiceDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateInvoiceDocumentKey, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocumentQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*CalculateInvoiceDocumentQueryGets, error) {
	pm := &requests.CalculateInvoiceDocumentQueryGets{
		ServiceLabel:                "",
		FieldNameWithNumberRange:    "",
		InvoiceDocumentLatestNumber: nil,
	}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.InvoiceDocumentLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	calculateInvoiceDocumentQueryGets := CalculateInvoiceDocumentQueryGets{
		ServiceLabel:                data.ServiceLabel,
		FieldNameWithNumberRange:    data.FieldNameWithNumberRange,
		InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
	}

	return &calculateInvoiceDocumentQueryGets, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocument(
	invoiceDocumentLatestNumber *int,
) (*CalculateInvoiceDocument, error) {
	pm := &requests.CalculateInvoiceDocument{
		InvoiceDocumentLatestNumber: nil,
		InvoiceDocument:             nil,
	}

	pm.InvoiceDocumentLatestNumber = invoiceDocumentLatestNumber
	data := pm

	calculateInvoiceDocument := CalculateInvoiceDocument{
		InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
		InvoiceDocument:             data.InvoiceDocument,
	}

	return &calculateInvoiceDocument, nil
}

func (psdc *SDC) ConvertToOrderItemKey(sdc *api_input_reader.SDC, orderID *[]OrderID) (*[]OrderItemKey, error) {

	var orderItemKey []OrderItemKey
	var pm []requests.OrderItemKey
	for _, v := range *orderID {
		pm = append(pm, requests.OrderItemKey{
			OrderID:                   v.OrderID,
			ItemInvoiceDocumentStatus: getBoolPtr(false),
		})
	}

	data := pm

	for _, v := range data {
		orderItemKey = append(orderItemKey, OrderItemKey{
			OrderID:                   v.OrderID,
			ItemInvoiceDocumentStatus: v.ItemInvoiceDocumentStatus,
		})
	}

	return &orderItemKey, nil
}

func (psdc *SDC) ConvertToOrderItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderItem, error) {
	var orderItem []OrderItem

	pm := &requests.OrderItem{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ItemInvoiceDocumentStatus,
			&pm.InvoiceDocumentItem,
			&pm.DocumentItemCategory,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.Division,
			&pm.Customer,
			&pm.Supplier,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.Batch,
			&pm.ProductGroup,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.InvoiceDocumentItemText,
			&pm.ServicesRenderedDate,
			&pm.InvoiceQuantity,
			&pm.InvoiceQuantityUnit,
			&pm.InvoiceQuantityInBaseUnit,
			&pm.BaseUnit,
			&pm.ActualGoodsIssueDate,
			&pm.ActualGoodsIssueTime,
			&pm.ActualGoodsReceiptDate,
			&pm.ActualGoodsReceiptTime,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.ItemVolume,
			&pm.ItemVolumeUnit,
			&pm.NetAmount,
			&pm.GoodsIssueOrReceiptSlipNumber,
			&pm.TransactionCurrency,
			&pm.BusinessPartnerCurrency,
			&pm.GrossAmount,
			&pm.PricingDate,
			&pm.TaxAmount,
			&pm.ProductTaxClassification1,
			&pm.BusinessArea,
			&pm.ProfitCenter,
			&pm.Project,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderType,
			&pm.ContractType,
			&pm.OrderVaridityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.OriginDocument,
			&pm.OriginDocumentItem,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ReferenceDocumentType,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.OrderSalesOrganization,
			&pm.OrderPurchaseOrganization,
			&pm.OrderDistributionChannel,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		orderItem = append(orderItem, OrderItem{
			ItemInvoiceDocumentStatus:     data.ItemInvoiceDocumentStatus,
			InvoiceDocumentItem:           data.InvoiceDocumentItem,
			DocumentItemCategory:          data.DocumentItemCategory,
			CreationDate:                  data.CreationDate,
			CreationTime:                  data.CreationTime,
			Division:                      data.Division,
			Customer:                      data.Customer,
			Supplier:                      data.Supplier,
			DeliverToParty:                data.DeliverToParty,
			DeliverFromParty:              data.DeliverFromParty,
			Product:                       data.Product,
			ProductStandardID:             data.ProductStandardID,
			Batch:                         data.Batch,
			ProductGroup:                  data.ProductGroup,
			IssuingPlant:                  data.IssuingPlant,
			IssuingPlantStorageLocation:   data.IssuingPlantStorageLocation,
			ReceivingPlant:                data.ReceivingPlant,
			ReceivingPlantStorageLocation: data.ReceivingPlantStorageLocation,
			InvoiceDocumentItemText:       data.InvoiceDocumentItemText,
			ServicesRenderedDate:          data.ServicesRenderedDate,
			InvoiceQuantity:               data.InvoiceQuantity,
			InvoiceQuantityUnit:           data.InvoiceQuantityUnit,
			InvoiceQuantityInBaseUnit:     data.InvoiceQuantityInBaseUnit,
			BaseUnit:                      data.BaseUnit,
			ActualGoodsIssueDate:          data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:          data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:        data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:        data.ActualGoodsReceiptTime,
			ItemGrossWeight:               data.ItemGrossWeight,
			ItemNetWeight:                 data.ItemNetWeight,
			ItemWeightUnit:                data.ItemWeightUnit,
			ItemVolume:                    data.ItemVolume,
			ItemVolumeUnit:                data.ItemVolumeUnit,
			NetAmount:                     data.NetAmount,
			GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
			TransactionCurrency:           data.TransactionCurrency,
			BusinessPartnerCurrency:       data.BusinessPartnerCurrency,
			GrossAmount:                   data.GrossAmount,
			PricingDate:                   data.PricingDate,
			TaxAmount:                     data.TaxAmount,
			ProductTaxClassification1:     data.ProductTaxClassification1,
			BusinessArea:                  data.BusinessArea,
			ProfitCenter:                  data.ProfitCenter,
			Project:                       data.Project,
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			OrderType:                     data.OrderType,
			ContractType:                  data.ContractType,
			OrderVaridityStartDate:        data.OrderVaridityStartDate,
			OrderValidityEndDate:          data.OrderValidityEndDate,
			InvoiceScheduleStartDate:      data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:        data.InvoiceScheduleEndDate,
			DeliveryDocument:              data.DeliveryDocument,
			DeliveryDocumentItem:          data.DeliveryDocumentItem,
			OriginDocument:                data.OriginDocument,
			OriginDocumentItem:            data.OriginDocumentItem,
			ReferenceDocument:             data.ReferenceDocument,
			ReferenceDocumentItem:         data.ReferenceDocumentItem,
			ReferenceDocumentType:         data.ReferenceDocumentType,
			ExternalReferenceDocument:     data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem: data.ExternalReferenceDocumentItem,
			OrderSalesOrganization:        data.OrderSalesOrganization,
			OrderPurchaseOrganization:     data.OrderPurchaseOrganization,
			OrderDistributionChannel:      data.OrderDistributionChannel,
			TaxCode:                       data.TaxCode,
			TaxRate:                       data.TaxRate,
			CountryOfOrigin:               data.CountryOfOrigin,
		})
	}

	return &orderItem, nil
}
