package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"fmt"
)

func (f *SubFunction) SetValue(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) (*dpfm_api_output_formatter.SDC, error) {
	var outItem *[]dpfm_api_output_formatter.Item
	var err error

	outItem, err = dpfm_api_output_formatter.ConvertToItem(sdc, psdc)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Item: *outItem,
	}

	return osdc, nil
}
