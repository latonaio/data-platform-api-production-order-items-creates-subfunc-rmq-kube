package subfunction

import (
	api_input_reader "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) ProductionOrderItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ProductionOrderItem {
	data := psdc.ConvertToProductionOrderItem(sdc)

	return data
}

func (f *SubFunction) ProductMasterBPPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterBPPlant, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductMasterBPPlantKey()

	for _, component := range psdc.PlannedOrderComponent {
		if component.ComponentProduct == nil || component.StockConfirmationBusinessPartner == nil || component.StockConfirmationPlant == nil {
			continue
		}
		dataKey.ComponentProduct = append(dataKey.ComponentProduct, *component.ComponentProduct)
		dataKey.StockConfirmationBusinessPartner = append(dataKey.StockConfirmationBusinessPartner, *component.StockConfirmationBusinessPartner)
		dataKey.StockConfirmationPlant = append(dataKey.StockConfirmationPlant, *component.StockConfirmationPlant)
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey.ComponentProduct)-1) + "(?,?,?)"
	for i := range dataKey.ComponentProduct {
		args = append(args, dataKey.ComponentProduct[i], dataKey.StockConfirmationBusinessPartner[i], dataKey.StockConfirmationPlant[i])
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, IsBatchManagementRequired, BatchManagementPolicy
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductMasterBPPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) BatchMasterRecordBatch(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.BatchMasterRecordBatch, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToBatchMasterRecordBatchKey()

	for _, component := range psdc.PlannedOrderComponent {
		if component.ComponentProduct == nil || component.StockConfirmationBusinessPartner == nil || component.StockConfirmationPlant == nil {
			continue
		}
		dataKey.ComponentProduct = append(dataKey.ComponentProduct, *component.ComponentProduct)
		dataKey.StockConfirmationBusinessPartner = append(dataKey.StockConfirmationBusinessPartner, *component.StockConfirmationBusinessPartner)
		dataKey.StockConfirmationPlant = append(dataKey.StockConfirmationPlant, *component.StockConfirmationPlant)
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey.ComponentProduct)-1) + "(?,?,?)"
	for i := range dataKey.ComponentProduct {
		args = append(args, dataKey.ComponentProduct[i], dataKey.StockConfirmationBusinessPartner[i], dataKey.StockConfirmationPlant[i])
	}
	args = append(args, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, Batch, ValidityStartDate, ValidityStartTime,
		ValidityEndDate, ValidityEndTime, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_batch_master_record_batch_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` )
		AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToBatchMasterRecordBatch(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) StockConfirmation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.StockConfirmation, error) {
	var err error
	dataKey := make([]*api_processing_data_formatter.StockConfirmationKey, 0)
	data := make([]*api_processing_data_formatter.StockConfirmation, 0)

	for _, component := range psdc.PlannedOrderComponent {
		datumKey := psdc.ConvertToStockConfirmationKey()

		if component.ComponentProduct == nil || component.StockConfirmationBusinessPartner == nil || component.StockConfirmationPlant == nil {
			continue
		}
		product := *component.ComponentProduct
		stockConfirmationBusinessPartner := *component.StockConfirmationBusinessPartner
		stockConfirmationPlant := *component.StockConfirmationPlant

		isBatchManagementRequired := new(bool)
		isBatchManagementRequired = nil
		for _, v := range psdc.ProductMasterBPPlant {
			if v.Product == product && v.BusinessPartner == stockConfirmationBusinessPartner && v.Plant == stockConfirmationPlant {
				if v.IsBatchManagementRequired == nil {
					continue
				}
				isBatchManagementRequired = v.IsBatchManagementRequired
				break
			}
		}

		if component.ComponentProduct == nil || component.StockConfirmationBusinessPartner == nil || component.StockConfirmationPlant == nil || component.ComponentProductRequirementDate == nil {
			continue
		}
		datumKey.Product = *component.ComponentProduct
		datumKey.StockConfirmationBusinessPartner = *component.StockConfirmationBusinessPartner
		datumKey.StockConfirmationPlant = *component.StockConfirmationPlant
		datumKey.ComponentProductRequirementDate = *component.ComponentProductRequirementDate

		if isBatchManagementRequired == nil {
			continue
		}
		if !*isBatchManagementRequired {
			datumKey.StockConfirmationIsOrdinary = true
		} else if *isBatchManagementRequired {
			for _, v := range psdc.BatchMasterRecordBatch {
				if v.Product == product && v.BusinessPartner == stockConfirmationBusinessPartner && v.Plant == stockConfirmationPlant {
					datumKey.StockConfirmationPlantBatch = v.Batch
					datumKey.StockConfirmationPlantBatchValidityStartDate = v.ValidityStartDate
					datumKey.StockConfirmationPlantBatchValidityEndDate = v.ValidityEndDate
					break
				}
			}
			datumKey.StockConfirmationIsLotUnit = true
		}

		plannedOrder := component.PlannedOrder
		plannedOrderItem := component.PlannedOrderItem
		for _, item := range psdc.PlannedOrderItem {
			if component.PlannedOrder == plannedOrder && component.PlannedOrderItem == plannedOrderItem {
				if item.PlannedOrderIssuingQuantity == nil {
					break
				}
				datumKey.PlannedOrderIssuingQuantity = *item.PlannedOrderIssuingQuantity
			}
		}

		dataKey = append(dataKey, datumKey)
	}

	for _, v := range dataKey {
		req, err := jsonTypeConversion[api_processing_data_formatter.ProductAvailabilityCheck](sdc)
		if err != nil {
			err = xerrors.Errorf("request create error: %w", err)
			return nil, err
		}

		if v.StockConfirmationIsOrdinary {
			req.ProductStockAvailabilityCheck.BusinessPartner = &v.StockConfirmationBusinessPartner
			req.ProductStockAvailabilityCheck.Product = &v.Product
			req.ProductStockAvailabilityCheck.Plant = &v.StockConfirmationPlant
			req.ProductStockAvailabilityCheck.RequestedQuantity = &v.PlannedOrderIssuingQuantity
			req.ProductStockAvailabilityCheck.ProductStockAvailabilityDate = &v.ComponentProductRequirementDate
		} else if v.StockConfirmationIsLotUnit {
			req.ProductStockAvailabilityCheck.BusinessPartner = &v.StockConfirmationBusinessPartner
			req.ProductStockAvailabilityCheck.Product = &v.Product
			req.ProductStockAvailabilityCheck.Plant = &v.StockConfirmationPlant
			req.ProductStockAvailabilityCheck.Batch = &v.StockConfirmationPlantBatch
			req.ProductStockAvailabilityCheck.RequestedQuantity = &v.PlannedOrderIssuingQuantity
			req.ProductStockAvailabilityCheck.ProductStockAvailabilityDate = &v.ComponentProductRequirementDate
		} else {
			continue
		}

		res, err := f.rmq.SessionKeepRequest(f.ctx, "data-platform-function-product-stock-availability-check-queue", req)
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil, err
		}
		res.Success()

		datum, err := psdc.ConvertToStockConfirmation(res.Data(), v.StockConfirmationIsOrdinary, v.StockConfirmationIsLotUnit)
		if err != nil {
			return nil, err
		}

		data = append(data, datum)
	}

	return data, err
}

func (f *SubFunction) ProductAvailabilityIsNotChecked(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.ProductAvailabilityIsNotChecked {
	productAvailabilityIsNotChecked := getBoolPtr(false)

	if len(psdc.StockConfirmation) > 0 {
		productAvailabilityIsNotChecked = getBoolPtr(true)
	}

	data := psdc.ConvertToProductAvailabilityIsNotChecked(productAvailabilityIsNotChecked)

	return data
}

func (f *SubFunction) InternalBillOfOperations(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.InternalBillOfOperations {
	data := make([]*api_processing_data_formatter.InternalBillOfOperations, 0)

	for _, component := range psdc.PlannedOrderComponent {
		plannedOrder := component.PlannedOrder
		plannedOrderItem := component.PlannedOrderItem
		internalBillOfOperations := component.OrderInternalBillOfOperations
		datum := psdc.ConvertToInternalBillOfOperations(plannedOrder, plannedOrderItem, internalBillOfOperations)

		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) TotalQuantity(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.TotalQuantity {
	data := make([]*api_processing_data_formatter.TotalQuantity, 0)

	for _, stockConfirmation := range psdc.StockConfirmation {
		datum := psdc.ConvertToTotalQuantity(stockConfirmation)

		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) PlannedScrapQuantityItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.PlannedScrapQuantityItem {
	data := make([]*api_processing_data_formatter.PlannedScrapQuantityItem, 0)

	for _, component := range psdc.PlannedOrderComponent {
		plannedOrder := component.PlannedOrder
		plannedOrderItem := component.PlannedOrderItem

		if component.ComponentScrapInPercent == nil {
			continue
		}
		componentScrapInPercent := *component.ComponentScrapInPercent

		if component.ComponentProduct == nil || component.StockConfirmationBusinessPartner == nil || component.StockConfirmationPlant == nil || component.StockConfirmationPlantBatch == nil || component.ComponentProductRequirementDate == nil {
			continue
		}

		businessPartner := *component.ComponentProductDeliverFromParty
		product := *component.ComponentProduct
		plant := *component.StockConfirmationPlant
		batch := *component.StockConfirmationPlantBatch
		productStockAvailabilityDate := *component.ComponentProductRequirementDate

		totalQuantityValue := float32(0)
		for _, totalQuantity := range psdc.TotalQuantity {
			if len(totalQuantity.Batch) == 0 {
				if totalQuantity.BusinessPartner == businessPartner && totalQuantity.Product == product && totalQuantity.Plant == plant && totalQuantity.ProductStockAvailabilityDate == productStockAvailabilityDate {
					totalQuantityValue = totalQuantity.TotalQuantity
					break
				}
			} else {
				if totalQuantity.BusinessPartner == businessPartner && totalQuantity.Product == product && totalQuantity.Plant == plant && totalQuantity.Batch == batch && totalQuantity.ProductStockAvailabilityDate == productStockAvailabilityDate {
					totalQuantityValue = totalQuantity.TotalQuantity
					break
				}
			}
		}

		plannedScrapQuantity := componentScrapInPercent * totalQuantityValue / 100

		datum := psdc.ConvertToPlannedScrapQuantityItem(plannedOrder, plannedOrderItem, componentScrapInPercent, totalQuantityValue, plannedScrapQuantity)

		data = append(data, datum)
	}

	return data
}

func jsonTypeConversion[T any](data interface{}) (T, error) {
	var dist T
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

// 日付等の処理
func (f *SubFunction) CreationDateItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationDate {
	data := psdc.ConvertToCreationDate(getSystemDate())

	return data
}

func (f *SubFunction) LastChangeDateItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeDate {
	data := psdc.ConvertToLastChangeDate(getSystemDate())

	return data
}

func getBoolPtr(b bool) *bool {
	return &b
}
