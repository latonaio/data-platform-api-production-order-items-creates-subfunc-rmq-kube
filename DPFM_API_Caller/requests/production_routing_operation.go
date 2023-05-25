package requests

type ProductionRoutingOperation struct {
	ProductionRoutingGroup        string  `json:"ProductionRoutingGroup"`
	ProductionRouting             string  `json:"ProductionRouting"`
	ProductionRoutingSequence     string  `json:"ProductionRoutingSequence"`
	ProductionRoutingOpIntID      string  `json:"ProductionRoutingOpIntID"`
	ProductionRoutingOpIntVersion string  `json:"ProductionRoutingOpIntVersion"`
	Operation                     *string `json:"Operation"`
	CreationDate                  *string `json:"CreationDate"`
	ChangeNumber                  *string `json:"ChangeNumber"`
	ValidityStartDate             *string `json:"ValidityStartDate"`
	ValidityEndDate               *string `json:"ValidityEndDate"`
	WorkCenterTypeCode            *string `json:"WorkCenterTypeCode"`
	WorkCenterInternalID          *string `json:"WorkCenterInternalID"`
	OperationSetupType            *string `json:"OperationSetupType"`
	OperationSetupGroupCategory   *string `json:"OperationSetupGroupCategory"`
	OperationSetupGroup           *string `json:"OperationSetupGroup"`
	OperationReferenceQuantity    *string `json:"OperationReferenceQuantity"`
	OperationUnit                 *string `json:"OperationUnit"`
	OpQtyToBaseQtyNmrtr           *string `json:"OpQtyToBaseQtyNmrtr"`
	OpQtyToBaseQtyDnmntr          *string `json:"OpQtyToBaseQtyDnmntr"`
	MaximumWaitDuration           *string `json:"MaximumWaitDuration"`
	MaximumWaitDurationUnit       *string `json:"MaximumWaitDurationUnit"`
	MinimumWaitDuration           *string `json:"MinimumWaitDuration"`
	MinimumWaitDurationUnit       *string `json:"MinimumWaitDurationUnit"`
	StandardQueueDuration         *string `json:"StandardQueueDuration"`
	StandardQueueDurationUnit     *string `json:"StandardQueueDurationUnit"`
	MinimumQueueDuration          *string `json:"MinimumQueueDuration"`
	MinimumQueueDurationUnit      *string `json:"MinimumQueueDurationUnit"`
	StandardMoveDuration          *string `json:"StandardMoveDuration"`
	StandardMoveDurationUnit      *string `json:"StandardMoveDurationUnit"`
	MinimumMoveDuration           *string `json:"MinimumMoveDuration"`
	MinimumMoveDurationUnit       *string `json:"MinimumMoveDurationUnit"`
	OpIsExtlyProcdWithSubcontrg   *bool   `json:"OpIsExtlyProcdWithSubcontrg"`
	PlannedDeliveryDuration       *int    `json:"PlannedDeliveryDuration"`
	MaterialGroup                 *string `json:"MaterialGroup"`
	PurchasingGroup               *string `json:"PurchasingGroup"`
	NumberOfOperationPriceUnits   *string `json:"NumberOfOperationPriceUnits"`
	CostElement                   *string `json:"CostElement"`
	OpExternalProcessingPrice     *string `json:"OpExternalProcessingPrice"`
	OpExternalProcessingCurrency  *string `json:"OpExternalProcessingCurrency"`
}
