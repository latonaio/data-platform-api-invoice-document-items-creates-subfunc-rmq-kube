package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Input_Reader"
	"data-platform-api-invoice-document-items-creates-subfunc-rmq/DPFM_API_Caller/requests"
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

func (psdc *SDC) ConvertToOrderIDByNumberSpecificationKey(sdc *api_input_reader.SDC, length int) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.BillFromParty = append(pm.BillFromParty, nil)
		pm.BillToParty = append(pm.BillToParty, nil)
	}

	data := pm
	orderIDKey := OrderIDKey{
		OrderID:                         data.OrderID,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderIDByNumberSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			InvoiceDocument:                 data.InvoiceDocument,
			OrderID:                         data.OrderID,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpecificationKey(sdc *api_input_reader.SDC) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}
	data := pm

	orderIDKey := OrderIDKey{
		OrderID:                         data.OrderID,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			InvoiceDocument:                 data.InvoiceDocument,
			OrderID:                         data.OrderID,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDByReferenceDocumentKey(sdc *api_input_reader.SDC) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(false),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}
	data := pm

	orderIDKey := OrderIDKey{
		OrderID:                         data.OrderID,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderIDByReferenceDocument(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			InvoiceDocument:                 data.InvoiceDocument,
			OrderID:                         data.OrderID,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderItemKey(sdc *api_input_reader.SDC, length int) (*OrderItemKey, error) {
	pm := &requests.OrderItemKey{
		ItemCompleteDeliveryIsDefined: getBoolPtr(true),
		ItemDeliveryStatus:            "CL",
		ItemBillingStatus:             "CL",
		ItemBillingBlockStatus:        getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.OrderID = append(pm.OrderID, nil)
	}

	data := pm

	orderItemKey := OrderItemKey{
		InvoiceDocument:               data.InvoiceDocument,
		OrderID:                       data.OrderID,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemBillingStatus:             data.ItemBillingStatus,
		ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
	}

	return &orderItemKey, nil
}

func (psdc *SDC) ConvertToOrderItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderItem, error) {
	var orderItem []OrderItem

	for i := 0; true; i++ {
		pm := &requests.OrderItem{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemBillingStatus,
			&pm.ItemBillingBlockStatus,
			&pm.OrderItem,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderItem = append(orderItem, OrderItem{
			InvoiceDocument:               data.InvoiceDocument,
			OrderID:                       data.OrderID,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemBillingStatus:             data.ItemBillingStatus,
			ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
			OrderItem:                     data.OrderItem,
		})
	}

	return &orderItem, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByNumberSpecificationKey(sdc *api_input_reader.SDC, length int) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.BillFromParty = append(pm.BillFromParty, nil)
		pm.BillToParty = append(pm.BillToParty, nil)
	}

	data := pm
	deliveryDocumentKey := DeliveryDocumentKey{
		DeliveryDocument:                data.DeliveryDocument,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByNumberSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocument{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			InvoiceDocument:                 data.InvoiceDocument,
			DeliveryDocument:                data.DeliveryDocument,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &deliveryDocument, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByRangeSpecificationKey(sdc *api_input_reader.SDC) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}
	data := pm

	deliveryDocumentKey := DeliveryDocumentKey{
		DeliveryDocument:                data.DeliveryDocument,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByRangeSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocument{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			InvoiceDocument:                 data.InvoiceDocument,
			DeliveryDocument:                data.DeliveryDocument,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &deliveryDocument, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByReferenceDocumentKey(sdc *api_input_reader.SDC) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}
	data := pm

	deliveryDocumentKey := DeliveryDocumentKey{
		DeliveryDocument:                data.DeliveryDocument,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByReferenceDocument(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocument{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			InvoiceDocument:                 data.InvoiceDocument,
			DeliveryDocument:                data.DeliveryDocument,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &deliveryDocument, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentItemKey(sdc *api_input_reader.SDC, length int) (*DeliveryDocumentItemKey, error) {
	pm := &requests.DeliveryDocumentItemKey{
		ItemCompleteDeliveryIsDefined: getBoolPtr(true),
		ItemDeliveryStatus:            "CL",
		ItemBillingStatus:             "CL",
		ItemBillingBlockStatus:        getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.DeliveryDocument = append(pm.DeliveryDocument, nil)
	}

	data := pm

	deliveryDocumentItemKey := DeliveryDocumentItemKey{
		DeliveryDocument:              data.DeliveryDocument,
		ConfirmedDeliveryDateFrom:     data.ConfirmedDeliveryDateFrom,
		ConfirmedDeliveryDateTo:       data.ConfirmedDeliveryDateTo,
		ActualGoodsIssueDateFrom:      data.ActualGoodsIssueDateFrom,
		ActualGoodsIssueDateTo:        data.ActualGoodsIssueDateTo,
		ConfirmedDeliveryDate:         data.ConfirmedDeliveryDate,
		ActualGoodsIssueDate:          data.ActualGoodsIssueDate,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemBillingStatus:             data.ItemBillingStatus,
		ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
	}

	return &deliveryDocumentItemKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocumentItem, error) {
	var deliveryDocumentItem []DeliveryDocumentItem

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocumentItem{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.ConfirmedDeliveryDate,
			&pm.ActualGoodsIssueDate,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemBillingStatus,
			&pm.ItemBillingBlockStatus,
			&pm.DeliveryDocumentItem,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocumentItem = append(deliveryDocumentItem, DeliveryDocumentItem{
			InvoiceDocument:               data.InvoiceDocument,
			DeliveryDocument:              data.DeliveryDocument,
			ConfirmedDeliveryDateFrom:     data.ConfirmedDeliveryDateFrom,
			ConfirmedDeliveryDateTo:       data.ConfirmedDeliveryDateTo,
			ActualGoodsIssueDateFrom:      data.ActualGoodsIssueDateFrom,
			ActualGoodsIssueDateTo:        data.ActualGoodsIssueDateTo,
			ConfirmedDeliveryDate:         data.ConfirmedDeliveryDate,
			ActualGoodsIssueDate:          data.ActualGoodsIssueDate,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemBillingStatus:             data.ItemBillingStatus,
			ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
			DeliveryDocumentItem:          data.DeliveryDocumentItem,
		})
	}

	return &deliveryDocumentItem, nil
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

func (psdc *SDC) ConvertToItemOrdersItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ItemOrdersItem, error) {
	var itemOrdersItem []ItemOrdersItem

	pm := &requests.ItemOrdersItem{}

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
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.ItemWeightUnit,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.ProductionPlantPartnerFunction,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.IssuingPlantPartnerFunction,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.Incoterms,
			&pm.ProductTaxClassification,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		itemOrdersItem = append(itemOrdersItem, ItemOrdersItem{
			OrderID:                        data.OrderID,
			OrderItem:                      data.OrderItem,
			OrderItemCategory:              data.OrderItemCategory,
			OrderItemText:                  data.OrderItemText,
			OrderItemTextByBuyer:           data.OrderItemTextByBuyer,
			OrderItemTextBySeller:          data.OrderItemTextBySeller,
			Product:                        data.Product,
			ProductStandardID:              data.ProductStandardID,
			ProductGroup:                   data.ProductGroup,
			BaseUnit:                       data.BaseUnit,
			PricingDate:                    data.PricingDate,
			OrderQuantityInBaseUnit:        data.OrderQuantityInBaseUnit,
			ItemWeightUnit:                 data.ItemWeightUnit,
			ItemGrossWeight:                data.ItemGrossWeight,
			ItemNetWeight:                  data.ItemNetWeight,
			NetAmount:                      data.NetAmount,
			TaxAmount:                      data.TaxAmount,
			GrossAmount:                    data.GrossAmount,
			ProductionPlantPartnerFunction: data.ProductionPlantPartnerFunction,
			ProductionPlantBusinessPartner: data.ProductionPlantBusinessPartner,
			ProductionPlant:                data.ProductionPlant,
			ProductionPlantStorageLocation: data.ProductionPlantStorageLocation,
			IssuingPlantPartnerFunction:    data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:    data.IssuingPlantBusinessPartner,
			IssuingPlant:                   data.IssuingPlant,
			IssuingPlantStorageLocation:    data.IssuingPlantStorageLocation,
			ReceivingPlantPartnerFunction:  data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:  data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                 data.ReceivingPlant,
			ReceivingPlantStorageLocation:  data.ReceivingPlantStorageLocation,
			Incoterms:                      data.Incoterms,
			ProductTaxClassification:       data.ProductTaxClassification,
			PaymentTerms:                   data.PaymentTerms,
			PaymentMethod:                  data.PaymentMethod,
			Project:                        data.Project,
			TaxCode:                        data.TaxCode,
			TaxRate:                        data.TaxRate,
			CountryOfOrigin:                data.CountryOfOrigin,
		})
	}

	return &itemOrdersItem, nil
}

func (psdc *SDC) ConvertToItemDeliveryDocumentItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]ItemDeliveryDocumentItem, error) {
	var itemDeliveryDocumentItem []ItemDeliveryDocumentItem

	pm := &requests.ItemDeliveryDocumentItem{}

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
			&pm.DeliveryDocumentItem,
			&pm.DeliveryDocumentItemCategory,
			&pm.DeliveryDocumentItemText,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.ActualGoodsIssueDate,
			&pm.ActualGoodsIssueTime,
			&pm.ActualGoodsReceiptDate,
			&pm.ActualGoodsReceiptTime,
			&pm.ProductionPlantPartnerFunction,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.IssuingPlantPartnerFunction,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderType,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.Project,
			&pm.ProductTaxClassification,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		itemDeliveryDocumentItem = append(itemDeliveryDocumentItem, ItemDeliveryDocumentItem{
			DeliveryDocument:               data.DeliveryDocument,
			DeliveryDocumentItem:           data.DeliveryDocumentItem,
			DeliveryDocumentItemCategory:   data.DeliveryDocumentItemCategory,
			DeliveryDocumentItemText:       data.DeliveryDocumentItemText,
			Product:                        data.Product,
			ProductStandardID:              data.ProductStandardID,
			ProductGroup:                   data.ProductGroup,
			BaseUnit:                       data.BaseUnit,
			ActualGoodsIssueDate:           data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:           data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:         data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:         data.ActualGoodsReceiptTime,
			ProductionPlantPartnerFunction: data.ProductionPlantPartnerFunction,
			ProductionPlantBusinessPartner: data.ProductionPlantBusinessPartner,
			ProductionPlant:                data.ProductionPlant,
			ProductionPlantStorageLocation: data.ProductionPlantStorageLocation,
			IssuingPlantPartnerFunction:    data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:    data.IssuingPlantBusinessPartner,
			IssuingPlant:                   data.IssuingPlant,
			IssuingPlantStorageLocation:    data.IssuingPlantStorageLocation,
			ReceivingPlantPartnerFunction:  data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:  data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                 data.ReceivingPlant,
			ReceivingPlantStorageLocation:  data.ReceivingPlantStorageLocation,
			ItemGrossWeight:                data.ItemGrossWeight,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemWeightUnit:                 data.ItemWeightUnit,
			OrderID:                        data.OrderID,
			OrderItem:                      data.OrderItem,
			OrderType:                      data.OrderType,
			ContractType:                   data.ContractType,
			OrderValidityStartDate:         data.OrderValidityStartDate,
			OrderValidityEndDate:           data.OrderValidityEndDate,
			PaymentTerms:                   data.PaymentTerms,
			PaymentMethod:                  data.PaymentMethod,
			InvoicePeriodStartDate:         data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:           data.InvoicePeriodEndDate,
			Project:                        data.Project,
			ProductTaxClassification:       data.ProductTaxClassification,
			TaxCode:                        data.TaxCode,
			TaxRate:                        data.TaxRate,
			CountryOfOrigin:                data.CountryOfOrigin,
		})
	}

	return &itemDeliveryDocumentItem, nil
}
