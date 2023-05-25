package subfunction

import (
	api_input_reader "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
)

func (f *SubFunction) SetValue(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	item, err := dpfm_api_output_formatter.ConvertToItem(sdc, psdc)
	if err != nil {
		return err
	}

	component, err := dpfm_api_output_formatter.ConvertToComponent(sdc, psdc)
	if err != nil {
		return err
	}

	componentStockConfirmation, err := dpfm_api_output_formatter.ConvertToComponentStockConfirmation(sdc, psdc)
	if err != nil {
		return err
	}

	componentCosting, err := dpfm_api_output_formatter.ConvertToComponentCosting(sdc, psdc)
	if err != nil {
		return err
	}

	operations, err := dpfm_api_output_formatter.ConvertToOperations(sdc, psdc)
	if err != nil {
		return err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Item:                           item,
		ItemComponent:                  component,
		ItemComponentStockConfirmation: componentStockConfirmation,
		ItemComponentCosting:           componentCosting,
		ItemOperations:                 operations,
	}

	return err
}
