package subfunction

import (
	api_input_reader "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
)

func (f *SubFunction) ProductionOrderComponent(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ProductionOrderComponent {
	data := make([]*api_processing_data_formatter.ProductionOrderComponent, 0)
	components := make([]*api_processing_data_formatter.PlannedOrderComponent, 0)

	for _, component := range psdc.PlannedOrderComponent {
		plannedOrder := component.PlannedOrder
		plannedOrderItem := component.PlannedOrderItem
		operations := component.Operations
		operationsItem := component.OperationsItem
		billOfMaterial := component.BillOfMaterial
		bomItem := component.BOMItem

		if operations == nil || operationsItem == nil || billOfMaterial == nil || bomItem == nil {
			continue
		}

		if componentContain(components, plannedOrder, plannedOrderItem, *operations, *operationsItem, *billOfMaterial, *bomItem) {
			continue
		}

		components = append(components, &api_processing_data_formatter.PlannedOrderComponent{
			PlannedOrder:     plannedOrder,
			PlannedOrderItem: plannedOrderItem,
			Operations:       operations,
			OperationsItem:   operationsItem,
			BillOfMaterial:   billOfMaterial,
			BOMItem:          bomItem,
		})

		datum := psdc.ConvertToProductionOrderComponent(plannedOrder, plannedOrderItem, *operations, *operationsItem, *billOfMaterial, *bomItem)
		data = append(data, datum)
	}

	return data
}

func componentContain(
	components []*api_processing_data_formatter.PlannedOrderComponent,
	plannedOrder, plannedOrderItem,
	operations, operationsItem,
	billOfMaterial, bOMItem int,
) bool {
	for _, component := range components {
		if plannedOrder == component.PlannedOrder && plannedOrderItem == component.PlannedOrderItem &&
			operations == *component.Operations && operationsItem == *component.OperationsItem &&
			billOfMaterial == *component.BillOfMaterial && bOMItem == *component.BOMItem {
			return true
		}
	}
	return false
}
