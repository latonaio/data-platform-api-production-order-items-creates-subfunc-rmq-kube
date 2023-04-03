package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Item, error) {
	var err error

	items := make([]*Item, 0, len(sdc.Header.Item))
	for i := range psdc.ProductionOrderItem {
		item := &Item{}
		inputItem := sdc.Header.Item[0]

		// 入力ファイル
		item, err = jsonTypeConversion(item, inputItem)
		if err != nil {
			return nil, err
		}

		item.ProductionOrder = sdc.Header.ProductionOrder
		item.ProductionOrderItem = psdc.ProductionOrderItem[i].ProductionOrderItemNumber
		item.CreationDate = &psdc.CreationDateItem.CreationDate
		item.LastChangeDate = &psdc.LastChangeDateItem.LastChangeDate
		item.ItemIsReleased = getBoolPtr(true)
		// item.ItemIsPartiallyConfirmed =  //2-1-4
		// item.ItemIsConfirmed =  //2-1-4
		item.ItemIsLocked = getBoolPtr(false)
		item.ItemIsMarkedForDeletion = getBoolPtr(false)
		item.ProductionOrderHasGeneratedOperations = getBoolPtr(true)
		item.ProductAvailabilityIsNotChecked = psdc.ProductAvailabilityIsNotChecked.ProductAvailabilityIsNotChecked
		// item.InternalBillOfOperations = psdc.InternalBillOfOperations[i].InternalBillOfOperations
		item.MRPController = psdc.PlannedOrderItem[i].MRPController
		item.ProductionVersion = getStringPtr("0001")
		item.PlannedOrder = &psdc.PlannedOrderItem[i].PlannedOrder
		item.OrderID = psdc.PlannedOrderItem[i].OrderID
		item.OrderItem = psdc.PlannedOrderItem[i].OrderItem
		item.TotalQuantity = &psdc.TotalQuantity[i].TotalQuantity
		item.PlannedScrapQuantity = &psdc.PlannedScrapQuantityItem[i].PlannedScrapQuantity
		item.ConfirmedYieldQuantity = &psdc.TotalQuantity[i].TotalQuantity

		items = append(items, item)
	}

	return items, nil
}

func ConvertToComponent(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Component, error) {
	var err error

	length := 0
	for _, v := range sdc.Header.Item {
		length += len(v.Component)
	}

	components := make([]*Component, 0, length)
	for _, item := range sdc.Header.Item {
		inputComponent := item.Component[0]

		for _, v := range item.Component {
			component := &Component{}

			_ = v

			// 入力ファイル
			component, err = jsonTypeConversion(component, inputComponent)
			if err != nil {
				return nil, err
			}

			component.ProductionOrder = sdc.Header.ProductionOrder
			// component.ProductionOrderItem = psdc.ProductionOrderItem[i].ProductionOrderItem
			// component.ProductionOrderSequence =  //1-1-3
			// component.ProductionOrderOperation =  //1-1-3
			// component.OrderInternalBillOfOperations =  //1-1-3
			// component.Reservation =  //1-1-3
			// component.ReservationItem =  //1-1-3
			// component.ComponentProduct =  //1-1-3
			// component.ComponentProductRequirementDate =  //1-1-3
			// component.ComponentProductRequirementTime =  //1-1-3
			component.ComponentProductIsMarkedForBackflush = getBoolPtr(false)
			// component.PlannedOrder =  //1-1-3
			// component.OrderID =  //1-1-3
			// component.OrderItem =  //1-1-3
			// component.BillOfMaterial =  //1-1-3
			// component.BOMItem =  //1-1-3
			// component.BOMItemDescription =  //1-1-3
			// component.Batch =  //2-3-2
			component.ProductCompIsAlternativeItem = getBoolPtr(false)
			component.CostingPolicy = getStringPtr("S")
			// component.ComponentScrapInPercent =  //1-1-3
			// component.OperationScrapInPercent =  //1-1-3
			// component.BaseUnit =  //1-1-3
			// component.RequiredQuantity =  //1-1-3
			// component.WithdrawnQuantity =  //1-1-3
			// component.ConfirmedAvailableQuantity =  //2-1
			// component.ProductCompOriginalQuantity =  //1-1-3
			component.IsMarkedForDeletion = getBoolPtr(false)

			components = append(components, component)
		}
	}

	return components, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
