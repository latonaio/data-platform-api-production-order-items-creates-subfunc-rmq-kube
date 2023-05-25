package subfunction

import (
	api_input_reader "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-production-order-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) ProductionRoutingHeaderInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionRoutingHeader, error) {
	data := make([]*api_processing_data_formatter.ProductionRoutingHeader, 0)
	var err error

	processType := psdc.ProcessType

	if processType.PluralitySpec {
		data, err = f.ProductionRoutingHeaderByPluralitySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.ProductionRoutingHeaderByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, xerrors.Errorf("OrderIDの絞り込み（一括登録）に必要な入力パラメータが揃っていません。")
	}

	return data, nil
}

func (f *SubFunction) ProductionRoutingHeaderByPluralitySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionRoutingHeader, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductionRoutingHeaderKey()

	productionPlantBusinessPartner := sdc.InputParameters.ProductionPlantBusinessPartner
	productionPlant := sdc.InputParameters.ProductionPlant

	dataKey.ProductionPlantBusinessPartner = append(dataKey.ProductionPlantBusinessPartner, *productionPlantBusinessPartner...)
	dataKey.ProductionPlant = append(dataKey.ProductionPlant, *productionPlant...)

	repeat1 := strings.Repeat("?,", len(dataKey.ProductionPlantBusinessPartner)-1) + "?"
	for _, v := range dataKey.ProductionPlantBusinessPartner {
		args = append(args, v)
	}
	repeat2 := strings.Repeat("?,", len(dataKey.ProductionPlant)-1) + "?"
	for _, v := range dataKey.ProductionPlant {
		args = append(args, v)
	}

	args = append(args, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT BusinessPartner, ProductionRoutingGroup, ProductionRouting, ProductionRoutingInternalVers,
		IsMarkedForDeletion, BillOfOperationsDesc, Plant, BillOfOperationsUsage, BillOfOperationsStatus,
		ResponsiblePlannerGroup, MinimumLotSizeQuantity, MaximumLotSizeQuantity, BillOfOperationsUnit,
		CreationDate, LastChangeDate, ValidityStartDate, ValidityEndDate, ChangeNumber,	PlainLongText, MaterialAssignment
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_routing_header_data
		WHERE BusinessPartner IN ( `+repeat1+` )
		AND Plant IN ( `+repeat2+` )
		AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductionRoutingHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionRoutingHeaderByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionRoutingHeader, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductionRoutingHeaderKey()

	dataKey.ProductionPlantBusinessPartnerTo = sdc.InputParameters.ProductionPlantBusinessPartnerTo
	dataKey.ProductionPlantBusinessPartnerFrom = sdc.InputParameters.ProductionPlantBusinessPartnerFrom
	dataKey.ProductionPlantTo = sdc.InputParameters.ProductionPlantTo
	dataKey.ProductionPlantFrom = sdc.InputParameters.ProductionPlantFrom

	args = append(args, dataKey.ProductionPlantBusinessPartnerFrom, dataKey.ProductionPlantBusinessPartnerTo,
		dataKey.ProductionPlantFrom, dataKey.ProductionPlantTo)

	args = append(args, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT BusinessPartner, ProductionRoutingGroup, ProductionRouting, ProductionRoutingInternalVers,
		IsMarkedForDeletion, BillOfOperationsDesc, Plant, BillOfOperationsUsage, BillOfOperationsStatus,
		ResponsiblePlannerGroup, MinimumLotSizeQuantity, MaximumLotSizeQuantity, BillOfOperationsUnit,
		CreationDate, LastChangeDate, ValidityStartDate, ValidityEndDate, ChangeNumber,	PlainLongText, MaterialAssignment
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_routing_header_data
		WHERE BusinessPartner BETWEEN ? AND ?
		AND Plant BETWEEN ? AND ?
		AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductionRoutingHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionRoutingProductPlantInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionRoutingProductPlant, error) {
	data := make([]*api_processing_data_formatter.ProductionRoutingProductPlant, 0)
	var err error

	processType := psdc.ProcessType

	if processType.PluralitySpec {
		data, err = f.ProductionRoutingProductPlantByPluralitySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.ProductionRoutingProductPlantByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, xerrors.Errorf("OrderIDの絞り込み（一括登録）に必要な入力パラメータが揃っていません。")
	}

	return data, nil
}

func (f *SubFunction) ProductionRoutingProductPlantByPluralitySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionRoutingProductPlant, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductionRoutingProductPlantKey()

	productInHeader := sdc.InputParameters.ProductInHeader

	dataKey.ProductInHeader = append(dataKey.ProductInHeader, *productInHeader...)

	for _, v := range psdc.ProductionRoutingHeader {
		if v.Plant == nil {
			continue
		}
		dataKey.BusinessPartner = append(dataKey.BusinessPartner, v.BusinessPartner)
		dataKey.Plant = append(dataKey.Plant, *v.Plant)
		dataKey.ProductionRoutingGroup = append(dataKey.ProductionRoutingGroup, v.ProductionRoutingGroup)
		dataKey.ProductionRouting = append(dataKey.ProductionRouting, v.ProductionRouting)
	}

	repeat1 := strings.Repeat("?,", len(dataKey.ProductInHeader)-1) + "?"
	for _, v := range dataKey.ProductInHeader {
		args = append(args, v)
	}
	repeat2 := strings.Repeat("(?,?,?,?),", len(dataKey.BusinessPartner)-1) + "(?,?,?,?)"
	for i := range dataKey.BusinessPartner {
		args = append(args, dataKey.BusinessPartner[i], dataKey.Plant[i], dataKey.ProductionRoutingGroup[i], dataKey.ProductionRouting[i])
	}

	// args = append(args, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT BusinessPartner, Product,  Plant, ProductionRoutingGroup, ProductionRouting, 
		ProductionRoutingMatlAssgmt, ProductionRtgMatlAssgmtIntVers, CreationDate, LastChangeDate, 
		ValidityStartDate, ValidityEndDate, ChangeNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_routing_product_plant_data
		WHERE Product IN ( `+repeat1+` )
		AND (BusinessPartner, Plant, ProductionRoutingGroup, ProductionRouting) IN ( `+repeat2+` );`, args...,
	// AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductionRoutingProductPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionRoutingProductPlantByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionRoutingProductPlant, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductionRoutingProductPlantKey()

	dataKey.ProductInHeaderTo = sdc.InputParameters.ProductInHeaderTo
	dataKey.ProductInHeaderFrom = sdc.InputParameters.ProductInHeaderFrom

	for _, v := range psdc.ProductionRoutingHeader {
		if v.Plant == nil {
			continue
		}
		dataKey.BusinessPartner = append(dataKey.BusinessPartner, v.BusinessPartner)
		dataKey.Plant = append(dataKey.Plant, *v.Plant)
		dataKey.ProductionRoutingGroup = append(dataKey.ProductionRoutingGroup, v.ProductionRoutingGroup)
		dataKey.ProductionRouting = append(dataKey.ProductionRouting, v.ProductionRouting)
	}

	args = append(args, dataKey.ProductInHeaderFrom, dataKey.ProductInHeaderTo)

	repeat := strings.Repeat("(?,?,?,?),", len(dataKey.BusinessPartner)-1) + "(?,?,?,?)"
	for i := range dataKey.BusinessPartner {
		args = append(args, dataKey.BusinessPartner[i], dataKey.Plant[i], dataKey.ProductionRoutingGroup[i], dataKey.ProductionRouting[i])
	}

	// args = append(args, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT BusinessPartner, Product,  Plant, ProductionRoutingGroup, ProductionRouting, 
		ProductionRoutingMatlAssgmt, ProductionRtgMatlAssgmtIntVers, CreationDate, LastChangeDate, 
		ValidityStartDate, ValidityEndDate, ChangeNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_routing_product_plant_data
		WHERE Product BETWEEN ? AND ?
		AND (BusinessPartner, Plant, ProductionRoutingGroup, ProductionRouting) IN ( `+repeat+` );`, args...,
	// AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductionRoutingProductPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionRoutingOperation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionRoutingOperation, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductionRoutingOperationKey()

	for _, v := range psdc.ProductionRoutingProductPlant {
		dataKey.ProductionRoutingGroup = append(dataKey.ProductionRoutingGroup, v.ProductionRoutingGroup)
		dataKey.ProductionRouting = append(dataKey.ProductionRouting, v.ProductionRouting)
		dataKey.Plant = append(dataKey.Plant, v.Plant)
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey.ProductionRoutingGroup)-1) + "(?,?,?)"
	for i := range dataKey.ProductionRoutingGroup {
		args = append(args, dataKey.ProductionRoutingGroup[i], dataKey.ProductionRouting[i], dataKey.Plant[i])
	}

	rows, err := f.db.Query(
		`SELECT ProductionRoutingGroup, ProductionRouting, ProductionRoutingSequence, ProductionRoutingOpIntID,
		ProductionRoutingOpIntVersion, Operation, CreationDate, ChangeNumber, ValidityStartDate, ValidityEndDate,
		WorkCenterTypeCode, WorkCenterInternalID, OperationSetupType, OperationSetupGroupCategory, OperationSetupGroup,
		OperationReferenceQuantity, OperationUnit, OpQtyToBaseQtyNmrtr, OpQtyToBaseQtyDnmntr, MaximumWaitDuration,
		MaximumWaitDurationUnit, MinimumWaitDuration, MinimumWaitDurationUnit, StandardQueueDuration,
		StandardQueueDurationUnit, MinimumQueueDuration, MinimumQueueDurationUnit, StandardMoveDuration,
		StandardMoveDurationUnit, MinimumMoveDuration, MinimumMoveDurationUnit, OpIsExtlyProcdWithSubcontrg,
		PlannedDeliveryDuration, MaterialGroup, PurchasingGroup, NumberOfOperationPriceUnits, CostElement,
		OpExternalProcessingPrice, OpExternalProcessingCurrency
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_routing_operation_data
		WHERE (ProductionRoutingGroup, ProductionRouting, Plant) IN ( `+repeat+` )`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductionRoutingOperation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
