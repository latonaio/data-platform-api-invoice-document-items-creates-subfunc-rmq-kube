package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"encoding/json"
)

func ConvertToItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]Item, error) {
	orderID := psdc.OrderID
	calculateInvoiceDocument := psdc.CalculateInvoiceDocument
	itemOrdersItem := psdc.ItemOrdersItem
	items := make([]Item, 0, len(*itemOrdersItem))

	orderIDMap := make(map[int]api_processing_data_formatter.OrderID, len(*orderID))

	for _, v := range *orderID {
		orderIDMap[*v.OrderID] = v
	}

	for _, v := range *itemOrdersItem {
		item := Item{}
		inputItem := sdc.InvoiceDocument.InvoiceDocumentItem[0]
		inputData, err := json.Marshal(inputItem)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(inputData, &item)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &item)
		if err != nil {
			return nil, err
		}

		item.InvoiceDocument = *calculateInvoiceDocument.InvoiceDocumentLatestNumber
		items = append(items, item)
	}

	return &items, nil
}
