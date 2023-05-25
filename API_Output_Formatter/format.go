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

	items := make([]*Item, 0, len(psdc.ProductionOrderItem))
	for i := range psdc.ProductionOrderItem {
		item := &Item{}
		inputItem := sdc.Header.Item[0]

		// 入力ファイル
		item, err = jsonTypeConversion(item, inputItem)
		if err != nil {
			return nil, err
		}

		plannedOrder := psdc.ProductionOrderItem[i].PlannedOrder
		plannedOrderItem := psdc.ProductionOrderItem[i].PlannedOrderItem

		quantityIdx := -1
		for j, v := range psdc.TotalQuantity {
			if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
				quantityIdx = j
				break
			}
		}
		if quantityIdx == -1 {
			continue
		}

		sQuantityIdx := -1
		for j, v := range psdc.PlannedScrapQuantityItem {
			if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
				sQuantityIdx = j
				break
			}
		}
		if sQuantityIdx == -1 {
			continue
		}

		item.ProductionOrder = sdc.Header.ProductionOrder
		item.ProductionOrderItem = psdc.ProductionOrderItem[i].ProductionOrderItemNumber
		item.CreationDate = psdc.CreationDateItem.CreationDate
		item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
		item.ItemIsReleased = getBoolPtr(true)
		item.ItemIsLocked = getBoolPtr(false)
		item.ItemIsMarkedForDeletion = getBoolPtr(false)
		item.ProductionOrderHasGeneratedOperations = getBoolPtr(true)
		item.ProductAvailabilityIsNotChecked = psdc.ProductAvailabilityIsNotChecked.ProductAvailabilityIsNotChecked
		// item.PrecedingItem = //TBA
		// item.FollowingItem = //TBA
		item.Product = psdc.PlannedOrderItem[i].Product
		item.ProductionPlant = getValue(psdc.PlannedOrderItem[i].ProductionPlant)
		item.ProductionPlantBusinessPartner = getValue(psdc.PlannedOrderItem[i].ProductionPlantBusinessPartner)
		item.ProductionPlantStorageLocation = psdc.PlannedOrderItem[i].ProductionPlantStorageLocation
		item.MRPArea = psdc.PlannedOrderItem[i].MRPArea
		item.MRPController = psdc.PlannedOrderItem[i].MRPController
		// item.ProductionVersion = getStringPtr("0001") //DBがstring型からint型に変更
		item.PlannedOrder = &psdc.PlannedOrderItem[i].PlannedOrder
		item.OrderID = psdc.PlannedOrderItem[i].OrderID
		item.OrderItem = psdc.PlannedOrderItem[i].OrderItem
		// item.MinimumLotSizeQuantity = //TBA
		// item.StandardLotSizeQuantity = //TBA
		// item.LotSizeRoundingQuantity = //TBA
		// item.MaximumLotSizeQuantity = //TBA
		// item.LotSizeIsFixed =
		item.ProductionOrderPlannedStartDate = psdc.PlannedOrderItem[i].PlannedOrderPlannedStartDate
		item.ProductionOrderPlannedStartTime = psdc.PlannedOrderItem[i].PlannedOrderPlannedStartTime
		item.ProductionOrderPlannedEndDate = psdc.PlannedOrderItem[i].PlannedOrderPlannedEndDate
		item.ProductionOrderPlannedEndTime = psdc.PlannedOrderItem[i].PlannedOrderPlannedEndTime
		item.ProductionUnit = psdc.PlannedOrderItem[i].PlannedOrderIssuingUnit
		item.TotalQuantity = getValue(psdc.TotalQuantity[quantityIdx].TotalQuantity)
		item.PlannedScrapQuantity = psdc.PlannedScrapQuantityItem[sQuantityIdx].PlannedScrapQuantity
		item.ConfirmedYieldQuantity = psdc.TotalQuantity[quantityIdx].TotalQuantity

		if psdc.ItemIsPartiallyConfirmed == nil || len(psdc.ItemIsPartiallyConfirmed) == 0 {
			item.ItemIsPartiallyConfirmed = nil
		} else {
			pConfirmedIdx := -1
			for j, v := range psdc.ItemIsPartiallyConfirmed {
				if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
					pConfirmedIdx = j
					break
				}
			}
			if pConfirmedIdx == -1 {
				continue
			}
			item.ItemIsPartiallyConfirmed = &psdc.ItemIsPartiallyConfirmed[pConfirmedIdx].ItemIsPartiallyConfirmed
		}

		if psdc.ItemIsConfirmed == nil || len(psdc.ItemIsConfirmed) == 0 {
			item.ItemIsConfirmed = nil
		} else {
			confirmedIdx := -1
			for j, v := range psdc.ItemIsConfirmed {
				if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
					confirmedIdx = j
					break
				}
			}
			if confirmedIdx == -1 {
				continue
			}
			item.ItemIsConfirmed = &psdc.ItemIsConfirmed[confirmedIdx].ItemIsConfirmed
		}

		items = append(items, item)
	}

	return items, nil
}

func ConvertToComponent(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemComponent, error) {
	var err error

	components := make([]*ItemComponent, 0, len(psdc.ProductionOrderComponent))
	for i := range psdc.ProductionOrderComponent {
		inputComponent := sdc.Header.Item[0].ItemComponent[0]

		component := &ItemComponent{}

		// 入力ファイル
		component, err = jsonTypeConversion(component, inputComponent)
		if err != nil {
			return nil, err
		}

		plannedOrder := psdc.ProductionOrderComponent[i].PlannedOrder
		plannedOrderItem := psdc.ProductionOrderComponent[i].PlannedOrderItem

		orderItemIdx := -1
		for j, v := range psdc.ProductionOrderItem {
			if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
				orderItemIdx = j
				break
			}
		}
		if orderItemIdx == -1 {
			continue
		}

		itemIdx := -1
		for j, v := range psdc.PlannedOrderItem {
			if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
				itemIdx = j
				break
			}
		}
		if itemIdx == -1 {
			continue
		}

		operations := psdc.ProductionOrderComponent[i].Operations
		operationsItem := psdc.ProductionOrderComponent[i].OperationsItem
		billOfMaterial := psdc.ProductionOrderComponent[i].BillOfMaterial
		bomItem := psdc.ProductionOrderComponent[i].BOMItem

		componentIdx := -1
		for j, v := range psdc.PlannedOrderComponent {
			if v.Operations == nil || v.OperationsItem == nil || v.BillOfMaterial == nil || v.BOMItem == nil {
				continue
			}
			if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem &&
				*v.Operations == operations && *v.OperationsItem == operationsItem &&
				*v.BillOfMaterial == billOfMaterial && *v.BOMItem == bomItem {
				componentIdx = j
				break
			}
		}
		if componentIdx == -1 {
			continue
		}

		businessPartner := psdc.PlannedOrderComponent[i].StockConfirmationBusinessPartner
		product := psdc.PlannedOrderComponent[i].ComponentProduct
		plant := psdc.PlannedOrderComponent[i].StockConfirmationPlant
		batch := psdc.PlannedOrderComponent[i].StockConfirmationPlantBatch
		productStockAvailabilityDate := psdc.PlannedOrderComponent[i].ComponentProductRequirementDate

		if businessPartner == nil || product == nil || plant == nil || batch == nil || productStockAvailabilityDate == nil {
			continue
		}

		component.ProductionOrder = sdc.Header.ProductionOrder
		component.ProductionOrderItem = psdc.ProductionOrderItem[orderItemIdx].ProductionOrderItemNumber
		// component.Operations = //TBA
		// component.OperationsItem = //TBA
		component.BillOfMaterial = getValue(psdc.PlannedOrderComponent[componentIdx].BillOfMaterial)
		component.BillOfMaterialItem = getValue(psdc.PlannedOrderComponent[componentIdx].BOMItem)
		component.Reservation = psdc.PlannedOrderComponent[componentIdx].Reservation
		component.ReservationItem = psdc.PlannedOrderComponent[componentIdx].ReservationItem
		component.ComponentProduct = psdc.PlannedOrderComponent[componentIdx].ComponentProduct
		component.ComponentProductRequirementDate = psdc.PlannedOrderComponent[componentIdx].ComponentProductRequirementDate
		component.ComponentProductRequirementTime = psdc.PlannedOrderComponent[componentIdx].ComponentProductRequirementTime
		component.ComponentProductIsMarkedForBackflush = getBoolPtr(false)
		component.ComponentProductBusinessPartner = psdc.PlannedOrderComponent[componentIdx].ComponentProductBusinessPartner
		component.StockConfirmationPlant = psdc.PlannedOrderComponent[componentIdx].StockConfirmationPlant
		component.PlannedOrder = &psdc.PlannedOrderComponent[componentIdx].PlannedOrder
		component.OrderID = psdc.PlannedOrderItem[itemIdx].OrderID     //1-1-3にない
		component.OrderItem = psdc.PlannedOrderItem[itemIdx].OrderItem //1-1-3にない
		// component.BOMItemDescription = //planned_order_componentからBOMItemDescriptionがなくなった
		component.StorageLocation = sdc.Header.Item[0].ProductionPlantStorageLocation
		component.ProductCompIsAlternativeItem = getBoolPtr(false)
		component.CostingPolicy = getStringPtr("S")
		component.ComponentScrapInPercent = psdc.PlannedOrderComponent[componentIdx].ComponentScrapInPercent
		component.OperationScrapInPercent = psdc.PlannedOrderComponent[componentIdx].OperationScrapInPercent
		component.BaseUnit = psdc.PlannedOrderComponent[componentIdx].BaseUnit
		component.RequiredQuantity = psdc.PlannedOrderComponent[componentIdx].ComponentProductRequiredQuantity
		component.WithdrawnQuantity = psdc.PlannedOrderComponent[componentIdx].ComponentWithdrawnQuantity
		component.ProductCompOriginalQuantity = psdc.PlannedOrderComponent[componentIdx].ComponentProductRequiredQuantity
		component.IsMarkedForDeletion = getBoolPtr(false)

		if psdc.StockConfirmation == nil || len(psdc.StockConfirmation) == 0 {
			component.Batch = nil
			component.ConfirmedAvailableQuantity = nil
		} else {
			stockConfIdx := -1
			for j, v := range psdc.StockConfirmation {
				if len(v.Batch) == 0 {
					if v.BusinessPartner == *businessPartner && v.Product == *product &&
						v.Plant == *plant && v.ProductStockAvailabilityDate == *productStockAvailabilityDate {
						stockConfIdx = j
						break
					}
				} else {
					if v.BusinessPartner == *businessPartner && v.Product == *product &&
						v.Plant == *plant && v.Batch == *batch && v.ProductStockAvailabilityDate == *productStockAvailabilityDate {
						stockConfIdx = j
						break
					}
				}
			}
			if stockConfIdx == -1 {
				continue
			}
			component.Batch = &psdc.StockConfirmation[stockConfIdx].Batch
			component.ConfirmedAvailableQuantity = &psdc.StockConfirmation[stockConfIdx].OpenConfirmedQuantityInBaseUnit
		}

		components = append(components, component)
	}

	return components, nil
}

func ConvertToComponentStockConfirmation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemComponentStockConfirmation, error) {
	var err error

	componentStockConfirmations := make([]*ItemComponentStockConfirmation, 0, len(psdc.ProductionOrderComponent))
	for i := range psdc.ProductionOrderComponent {
		inputComponentStockConfirmation := sdc.Header.Item[0].ItemComponent[0].ItemComponentStockConfirmation[0]

		componentStockConfirmation := &ItemComponentStockConfirmation{}

		// 入力ファイル
		componentStockConfirmation, err = jsonTypeConversion(componentStockConfirmation, inputComponentStockConfirmation)
		if err != nil {
			return nil, err
		}

		plannedOrder := psdc.ProductionOrderComponent[i].PlannedOrder
		plannedOrderItem := psdc.ProductionOrderComponent[i].PlannedOrderItem

		orderItemIdx := -1
		for j, v := range psdc.ProductionOrderItem {
			if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
				orderItemIdx = j
				break
			}
		}
		if orderItemIdx == -1 {
			continue
		}

		componentStockConfirmation.ProductionOrder = sdc.Header.ProductionOrder
		componentStockConfirmation.ProductionOrderItem = psdc.ProductionOrderItem[orderItemIdx].ProductionOrderItemNumber
		componentStockConfirmation.IsMarkedForDeletion = getBoolPtr(false)

		componentStockConfirmations = append(componentStockConfirmations, componentStockConfirmation)
	}

	return componentStockConfirmations, nil
}

func ConvertToComponentCosting(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemComponentCosting, error) {
	var err error

	componentCostings := make([]*ItemComponentCosting, 0, len(psdc.ProductionOrderComponent))
	for i := range psdc.ProductionOrderComponent {
		inputComponentCosting := sdc.Header.Item[0].ItemComponent[0].ItemComponentCosting[0]

		componentCosting := &ItemComponentCosting{}

		// 入力ファイル
		componentCosting, err = jsonTypeConversion(componentCosting, inputComponentCosting)
		if err != nil {
			return nil, err
		}

		plannedOrder := psdc.ProductionOrderComponent[i].PlannedOrder
		plannedOrderItem := psdc.ProductionOrderComponent[i].PlannedOrderItem

		orderItemIdx := -1
		for j, v := range psdc.ProductionOrderItem {
			if v.PlannedOrder == plannedOrder && v.PlannedOrderItem == plannedOrderItem {
				orderItemIdx = j
				break
			}
		}
		if orderItemIdx == -1 {
			continue
		}

		componentCosting.ProductionOrder = sdc.Header.ProductionOrder
		componentCosting.ProductionOrderItem = psdc.ProductionOrderItem[orderItemIdx].ProductionOrderItemNumber
		componentCosting.IsMarkedForDeletion = getBoolPtr(false)

		componentCostings = append(componentCostings, componentCosting)
	}

	return componentCostings, nil
}

func ConvertToOperations(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemOperations, error) {
	var err error

	length := 0
	for _, item := range sdc.Header.Item {
		length += len(item.ItemOperations)

	}

	operations := make([]*ItemOperations, 0, length)
	for i, item := range sdc.Header.Item {
		for j := range item.ItemOperations {
			operation := &ItemOperations{}
			inputOperation := item.ItemOperations[j]

			// 入力ファイル
			operation, err = jsonTypeConversion(operation, inputOperation)
			if err != nil {
				return nil, err
			}

			operation.ProductionOrder = sdc.Header.ProductionOrder
			operation.ProductionOrderItem = psdc.ProductionOrderItem[i].ProductionOrderItemNumber
			operation.OperationIsReleased = getBoolPtr(false)
			operation.OperationIsPartiallyConfirmed = getBoolPtr(false)
			operation.OperationIsConfirmed = getBoolPtr(false)
			operation.OperationIsClosed = getBoolPtr(false)
			operation.OperationIsMarkedForDeletion = getBoolPtr(false)
			// operation.OperationUnit =  //2-3-3

			operations = append(operations, operation)
		}

	}

	return operations, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func getValue[T any](ptr *T) T {
	var zero T
	if ptr == nil {
		return zero
	}
	return *ptr
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
