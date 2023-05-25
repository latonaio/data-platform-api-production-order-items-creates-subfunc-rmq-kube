package api_processing_data_formatter

type SDC struct {
	MetaData                        *MetaData                        `json:"MetaData"`
	ProcessType                     *ProcessType                     `json:"ProcessType"`
	PlannedOrderHeader              []*PlannedOrderHeader            `json:"PlannedOrderHeader"`
	PlannedOrderItem                []*PlannedOrderItem              `json:"PlannedOrderItem"`
	PlannedOrderComponent           []*PlannedOrderComponent         `json:"PlannedOrderComponent"`
	ProductionOrderItem             []*ProductionOrderItem           `json:"ProductionOrderItem"`
	ExecuteProductAvailabilityCheck *ExecuteProductAvailabilityCheck `json:"ExecuteProductAvailabilityCheck"`
	ProductMasterBPPlant            []*ProductMasterBPPlant          `json:"ProductMasterBPPlant"`
	BatchMasterRecordBatch          []*BatchMasterRecordBatch        `json:"BatchMasterRecordBatch"`
	StockConfirmation               []*StockConfirmation             `json:"StockConfirmation"`
	ItemIsPartiallyConfirmed        []*ItemIsPartiallyConfirmed      `json:"ItemIsPartiallyConfirmed"`
	ItemIsConfirmed                 []*ItemIsConfirmed               `json:"ItemIsConfirmed"`
	ProductAvailabilityIsNotChecked *ProductAvailabilityIsNotChecked `json:"ProductAvailabilityIsNotChecked"`
	InternalBillOfOperations        []*InternalBillOfOperations      `json:"InternalBillOfOperations"`
	TotalQuantity                   []*TotalQuantity                 `json:"TotalQuantity"`
	PlannedScrapQuantityItem        []*PlannedScrapQuantityItem      `json:"PlannedScrapQuantityItem"`
	ProductionOrderComponent        []*ProductionOrderComponent      `json:"ProductionOrderComponent"`
	ProductionRoutingHeader         []*ProductionRoutingHeader       `json:"ProductionRoutingHeader"`
	ProductionRoutingProductPlant   []*ProductionRoutingProductPlant `json:"ProductionRoutingProductPlant"`
	ProductionRoutingOperation      []*ProductionRoutingOperation    `json:"ProductionRoutingOperation"`
	CreationDateItem                *CreationDate                    `json:"CreationDateItem"`
	LastChangeDateItem              *LastChangeDate                  `json:"LastChangeDateItem"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type ProcessType struct {
	BulkProcess       bool `json:"BulkProcess"`
	IndividualProcess bool `json:"IndividualProcess"`
	PluralitySpec     bool `json:"PluralitySpec"`
	RangeSpec         bool `json:"RangeSpec"`
}

// Header
type PlannedOrderHeaderKey struct {
	MRPArea                                 []*string `json:"MRPArea"`
	MRPAreaTo                               *string   `json:"MRPAreaTo"`
	MRPAreaFrom                             *string   `json:"MRPAreaFrom"`
	OwnerProductionPlantBusinessPartner     []*int    `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerTo   *int      `json:"OwnerProductionPlantBusinessPartnerTo"`
	OwnerProductionPlantBusinessPartnerFrom *int      `json:"OwnerProductionPlantBusinessPartnerFrom"`
	OwnerProductionPlant                    []*string `json:"OwnerProductionPlant"`
	OwnerProductionPlantTo                  *string   `json:"OwnerProductionPlantTo"`
	OwnerProductionPlantFrom                *string   `json:"OwnerProductionPlantFrom"`
	ProductInHeader                         []*string `json:"ProductInHeader"`
	ProductInHeaderTo                       *string   `json:"ProductInHeaderTo"`
	ProductInHeaderFrom                     *string   `json:"ProductInHeaderFrom"`
	PlannedOrderType                        string    `json:"PlannedOrderType"`
	PlannedOrderIsReleased                  bool      `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                     bool      `json:"IsMarkedForDeletion"`
}

type PlannedOrderHeader struct {
	PlannedOrder                             int      `json:"PlannedOrder"`
	PlannedOrderType                         *string  `json:"PlannedOrderType"`
	Product                                  *string  `json:"Product"`
	ProductDeliverFromParty                  *int     `json:"ProductDeliverFromParty"`
	ProductDeliverToParty                    *int     `json:"ProductDeliverToParty"`
	OriginIssuingPlant                       *string  `json:"OriginIssuingPlant"`
	OriginIssuingPlantStorageLocation        *string  `json:"OriginIssuingPlantStorageLocation"`
	DestinationReceivingPlant                *string  `json:"DestinationReceivingPlant"`
	DestinationReceivingPlantStorageLocation *string  `json:"DestinationReceivingPlantStorageLocation"`
	OwnerProductionPlantBusinessPartner      *int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     *string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantStorageLocation      *string  `json:"OwnerProductionPlantStorageLocation"`
	MRPArea                                  *string  `json:"MRPArea"`
	MRPController                            *string  `json:"MRPController"`
	PlannedOrderQuantityInBaseUnit           *float32 `json:"PlannedOrderQuantityInBaseUnit"`
	PlannedOrderPlannedScrapQtyInBaseUnit    *float32 `json:"PlannedOrderPlannedScrapQtyInBaseUnit"`
	PlannedOrderOriginIssuingUnit            *string  `json:"PlannedOrderOriginIssuingUnit"`
	PlannedOrderDestinationReceivingUnit     *string  `json:"PlannedOrderDestinationReceivingUnit"`
	PlannedOrderOriginIssuingQuantity        *float32 `json:"PlannedOrderOriginIssuingQuantity"`
	PlannedOrderDestinationReceivingQuantity *float32 `json:"PlannedOrderDestinationReceivingQuantity"`
	PlannedOrderPlannedStartDate             *string  `json:"PlannedOrderPlannedStartDate"`
	PlannedOrderPlannedStartTime             *string  `json:"PlannedOrderPlannedStartTime"`
	PlannedOrderPlannedEndDate               *string  `json:"PlannedOrderPlannedEndDate"`
	PlannedOrderPlannedEndTime               *string  `json:"PlannedOrderPlannedEndTime"`
	LastChangeDateTime                       *string  `json:"LastChangeDateTime"`
	OrderID                                  *int     `json:"OrderID"`
	OrderItem                                *int     `json:"OrderItem"`
	ProductBuyer                             *int     `json:"ProductBuyer"`
	ProductSeller                            *int     `json:"ProductSeller"`
	Project                                  *string  `json:"Project"`
	Reservation                              *int     `json:"Reservation"`
	ReservationItem                          *int     `json:"ReservationItem"`
	PlannedOrderLongText                     *string  `json:"PlannedOrderLongText"`
	PlannedOrderIsFixed                      *bool    `json:"PlannedOrderIsFixed"`
	PlannedOrderBOMIsFixed                   *bool    `json:"PlannedOrderBOMIsFixed"`
	LastScheduledDate                        *string  `json:"LastScheduledDate"`
	ScheduledBasicEndDate                    *string  `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime                    *string  `json:"ScheduledBasicEndTime"`
	ScheduledBasicStartDate                  *string  `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime                  *string  `json:"ScheduledBasicStartTime"`
	SchedulingType                           *string  `json:"SchedulingType"`
	PlannedOrderIsReleased                   *bool    `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

type PlannedOrderItemKey struct {
	PlannedOrder                       []int     `json:"PlannedOrder"`
	MRPArea                            []*string `json:"MRPArea"`
	MRPAreaTo                          *string   `json:"MRPAreaTo"`
	MRPAreaFrom                        *string   `json:"MRPAreaFrom"`
	ProductionPlantBusinessPartner     []*int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerTo   *int      `json:"ProductionPlantBusinessPartnerTo"`
	ProductionPlantBusinessPartnerFrom *int      `json:"ProductionPlantBusinessPartnerFrom"`
	ProductionPlant                    []*string `json:"ProductionPlant"`
	ProductionPlantTo                  *string   `json:"ProductionPlantTo"`
	ProductionPlantFrom                *string   `json:"ProductionPlantFrom"`
	ProductionPlantStorageLocation     []*string `json:"ProductionPlantStorageLocation"`
	ProductionPlantStorageLocationTo   *string   `json:"ProductionPlantStorageLocationTo"`
	ProductionPlantStorageLocationFrom *string   `json:"ProductionPlantStorageLocationFrom"`
	ProductInItem                      []*string `json:"ProductInItem"`
	ProductInItemTo                    *string   `json:"ProductInItemTo"`
	ProductInItemFrom                  *string   `json:"ProductInItemFrom"`
	PlannedOrderIsReleased             bool      `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                bool      `json:"IsMarkedForDeletion"`
}

type PlannedOrderItem struct {
	PlannedOrder                          int      `json:"PlannedOrder"`
	PlannedOrderItem                      int      `json:"PlannedOrderItem"`
	Product                               *string  `json:"Product"`
	ProductDeliverFromParty               *int     `json:"ProductDeliverFromParty"`
	ProductDeliverToParty                 *int     `json:"ProductDeliverToParty"`
	IssuingPlant                          *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation           *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlant                        *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation         *string  `json:"ReceivingPlantStorageLocation"`
	ProductionPlantBusinessPartner        *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                       *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation        *string  `json:"ProductionPlantStorageLocation"`
	BaseUnit                              *string  `json:"BaseUnit"`
	MRPArea                               *string  `json:"MRPArea"`
	MRPController                         *string  `json:"MRPController"`
	PlannedOrderQuantityInBaseUnit        *float32 `json:"PlannedOrderQuantityInBaseUnit"`
	PlannedOrderPlannedScrapQtyInBaseUnit *float32 `json:"PlannedOrderPlannedScrapQtyInBaseUnit"`
	PlannedOrderIssuingUnit               *string  `json:"PlannedOrderIssuingUnit"`
	PlannedOrderReceivingUnit             *string  `json:"PlannedOrderReceivingUnit"`
	PlannedOrderIssuingQuantity           *float32 `json:"PlannedOrderIssuingQuantity"`
	PlannedOrderReceivingQuantity         *float32 `json:"PlannedOrderReceivingQuantity"`
	PlannedOrderPlannedStartDate          *string  `json:"PlannedOrderPlannedStartDate"`
	PlannedOrderPlannedStartTime          *string  `json:"PlannedOrderPlannedStartTime"`
	PlannedOrderPlannedEndDate            *string  `json:"PlannedOrderPlannedEndDate"`
	PlannedOrderPlannedEndTime            *string  `json:"PlannedOrderPlannedEndTime"`
	LastChangeDateTime                    *string  `json:"LastChangeDateTime"`
	OrderID                               *int     `json:"OrderID"`
	OrderItem                             *int     `json:"OrderItem"`
	ProductBuyer                          *int     `json:"ProductBuyer"`
	ProductSeller                         *int     `json:"ProductSeller"`
	Project                               *string  `json:"Project"`
	Reservation                           *int     `json:"Reservation"`
	ReservationItem                       *int     `json:"ReservationItem"`
	PlannedOrderLongText                  *string  `json:"PlannedOrderLongText"`
	PlannedOrderIsFixed                   *bool    `json:"PlannedOrderIsFixed"`
	PlannedOrderBOMIsFixed                *bool    `json:"PlannedOrderBOMIsFixed"`
	LastScheduledDate                     *string  `json:"LastScheduledDate"`
	ScheduledBasicEndDate                 *string  `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime                 *string  `json:"ScheduledBasicEndTime"`
	ScheduledBasicStartDate               *string  `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime               *string  `json:"ScheduledBasicStartTime"`
	SchedulingType                        *string  `json:"SchedulingType"`
	PlannedOrderIsReleased                *bool    `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                   *bool    `json:"IsMarkedForDeletion"`
}

type PlannedOrderComponentKey struct {
	PlannedOrder        []int `json:"PlannedOrder"`
	PlannedOrderItem    []int `json:"PlannedOrderItem"`
	IsMarkedForDeletion bool  `json:"IsMarkedForDeletion"`
}

type PlannedOrderComponent struct {
	PlannedOrder                     int      `json:"PlannedOrder"`
	PlannedOrderItem                 int      `json:"PlannedOrderItem"`
	BillOfMaterial                   *int     `json:"BillOfMaterial"`
	BOMItem                          *int     `json:"BOMItem"`
	Operations                       *int     `json:"Operations"`
	OperationsItem                   *int     `json:"OperationsItem"`
	Reservation                      *int     `json:"Reservation"`
	ReservationItem                  *int     `json:"ReservationItem"`
	ComponentProduct                 *string  `json:"ComponentProduct"`
	ComponentProductDeliverFromParty *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverToParty   *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductBuyer            *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller           *int     `json:"ComponentProductSeller"`
	ComponentProductRequirementDate  *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime  *string  `json:"ComponentProductRequirementTime"`
	ComponentProductRequiredQuantity *float32 `json:"ComponentProductRequiredQuantity"`
	ComponentProductBusinessPartner  *int     `json:"ComponentProductBusinessPartner"`
	BaseUnit                         *string  `json:"BaseUnit"`
	MRPArea                          *string  `json:"MRPArea"`
	MRPController                    *string  `json:"MRPController"`
	StockConfirmationPartnerFunction *string  `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch      *string  `json:"StockConfirmationPlantBatch"`
	StorageLocationForMRP            *string  `json:"StorageLocationForMRP"`
	ComponentWithdrawnQuantity       *float32 `json:"ComponentWithdrawnQuantity"`
	ComponentScrapInPercent          *float32 `json:"ComponentScrapInPercent"`
	OperationScrapInPercent          *float32 `json:"OperationScrapInPercent"`
	QuantityIsFixed                  *bool    `json:"QuantityIsFixed"`
	LastChangeDateTime               *string  `json:"LastChangeDateTime"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
}

// Item
type ProductionOrderItem struct {
	PlannedOrder              int `json:"PlannedOrder"`
	PlannedOrderItem          int `json:"PlannedOrderItem"`
	ProductionOrderItemNumber int `json:"ProductionOrderItemNumber"`
}

type ExecuteProductAvailabilityCheck struct {
	ExecuteProductAvailabilityCheck bool `json:"ExecuteProductAvailabilityCheck"`
}

type ProductMasterBPPlantKey struct {
	ComponentProduct                 []string `json:"ComponentProduct"`
	StockConfirmationBusinessPartner []int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           []string `json:"StockConfirmationPlant"`
}

type ProductMasterBPPlant struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	IsBatchManagementRequired *bool   `json:"IsBatchManagementRequired"`
	BatchManagementPolicy     *string `json:"BatchManagementPolicy"`
}

type StockConfirmationKey struct {
	PlannedOrder                                 int     `json:"PlannedOrder"`
	PlannedOrderItem                             int     `json:"PlannedOrderItem"`
	Product                                      string  `json:"Product"`
	StockConfirmationBusinessPartner             int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	PlannedOrderIssuingQuantity                  float32 `json:"PlannedOrderIssuingQuantity"`
	ComponentProductRequirementDate              string  `json:"ComponentProductRequirementDate"`
	StockConfirmationIsLotUnit                   bool    `json:"StockConfirmationIsLotUnit"`
	StockConfirmationIsOrdinary                  bool    `json:"StockConfirmationIsOrdinary"`
}

type StockConfirmation struct {
	PlannedOrder                    int     `json:"PlannedOrder"`
	PlannedOrderItem                int     `json:"PlannedOrderItem"`
	BusinessPartner                 int     `json:"BusinessPartner"`
	Product                         string  `json:"Product"`
	Plant                           string  `json:"Plant"`
	Batch                           string  `json:"Batch"`
	RequestedQuantity               float32 `json:"RequestedQuantity"`
	ProductStockAvailabilityDate    string  `json:"ProductStockAvailabilityDate"`
	OrderID                         int     `json:"OrderID"`
	OrderItem                       int     `json:"OrderItem"`
	Project                         string  `json:"Project"`
	InventoryStockType              string  `json:"InventoryStockType"`
	InventorySpecialStockType       string  `json:"InventorySpecialStockType"`
	AvailableProductStock           float32 `json:"AvailableProductStock"`
	CheckedQuantity                 float32 `json:"CheckedQuantity"`
	CheckedDate                     string  `json:"CheckedDate"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
	Suffix                          string  `json:"Suffix"`
	StockConfirmationIsLotUnit      bool    `json:"StockConfirmationIsLotUnit"`
	StockConfirmationIsOrdinary     bool    `json:"StockConfirmationIsOrdinary"`
}

type ProductAvailabilityCheck struct {
	ConnectionKey                 string `json:"connection_key"`
	Result                        bool   `json:"result"`
	RedisKey                      string `json:"redis_key"`
	Filepath                      string `json:"filepath"`
	APIStatusCode                 int    `json:"api_status_code"`
	RuntimeSessionID              string `json:"runtime_session_id"`
	BusinessPartnerID             *int   `json:"business_partner"`
	ServiceLabel                  string `json:"service_label"`
	ProductStockAvailabilityCheck struct {
		Product                         *string  `json:"Product"`
		BusinessPartner                 *int     `json:"BusinessPartner"`
		Plant                           *string  `json:"Plant"`
		StorageLocation                 *string  `json:"StorageLocation"`
		StorageBin                      *string  `json:"StorageBin"`
		Batch                           *string  `json:"Batch"`
		RequestedQuantity               *float32 `json:"RequestedQuantity"`
		ProductStockAvailabilityDate    *string  `json:"ProductStockAvailabilityDate"`
		InventoryStockType              *string  `json:"InventoryStockType"`
		InventorySpecialStockType       *string  `json:"InventorySpecialStockType"`
		AvailableProductStock           *float32 `json:"AvailableProductStock"`
		CheckedQuantity                 *float32 `json:"CheckedQuantity"`
		CheckedDate                     *string  `json:"CheckedDate"`
		OpenConfirmedQuantityInBaseUnit *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
		StockIsFullyChecked             *bool    `json:"StockIsFullyChecked"`
		Suffix                          *string  `json:"Suffix"`
	} `json:"ProductStockAvailabilityCheck"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	ProductStockCode string   `json:"product_stock_code"`
	Deleted          bool     `json:"deleted"`
}

type BatchMasterRecordBatchKey struct {
	ComponentProduct                 []string `json:"ComponentProduct"`
	StockConfirmationBusinessPartner []int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           []string `json:"StockConfirmationPlant"`
	IsMarkedForDeletion              bool     `json:"IsMarkedForDeletion"`
}

type BatchMasterRecordBatch struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	Batch               string `json:"Batch"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityStartTime   string `json:"ValidityStartTime"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	ValidityEndTime     string `json:"ValidityEndTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ItemIsPartiallyConfirmed struct {
	PlannedOrder                    int     `json:"PlannedOrder"`
	PlannedOrderItem                int     `json:"PlannedOrderItem"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	ItemIsPartiallyConfirmed        bool    `json:"ItemIsPartiallyConfirmed"`
}

type ItemIsConfirmed struct {
	PlannedOrder        int  `json:"PlannedOrder"`
	PlannedOrderItem    int  `json:"PlannedOrderItem"`
	StockIsFullyChecked bool `json:"StockIsFullyChecked"`
	ItemIsConfirmed     bool `json:"ItemIsConfirmed"`
}

type ProductAvailabilityIsNotChecked struct {
	ProductAvailabilityIsNotChecked *bool `json:"ProductAvailabilityIsNotChecked"`
}

type InternalBillOfOperations struct {
	PlannedOrder             int  `json:"PlannedOrder"`
	PlannedOrderItem         int  `json:"PlannedOrderItem"`
	InternalBillOfOperations *int `json:"InternalBillOfOperations"`
}

type TotalQuantity struct {
	PlannedOrder     int      `json:"PlannedOrder"`
	PlannedOrderItem int      `json:"PlannedOrderItem"`
	TotalQuantity    *float32 `json:"TotalQuantity"`
}

type PlannedScrapQuantityItem struct {
	PlannedOrder            int      `json:"PlannedOrder"`
	PlannedOrderItem        int      `json:"PlannedOrderItem"`
	ComponentScrapInPercent float32  `json:"ComponentScrapInPercent"`
	TotalQuantity           *float32 `json:"TotalQuantity"`
	PlannedScrapQuantity    *float32 `json:"PlannedScrapQuantity"`
}

// Component
type ProductionOrderComponent struct {
	PlannedOrder     int `json:"PlannedOrder"`
	PlannedOrderItem int `json:"PlannedOrderItem"`
	Operations       int `json:"Operations"`
	OperationsItem   int `json:"OperationsItem"`
	BillOfMaterial   int `json:"BillOfMaterial"`
	BOMItem          int `json:"BOMItem"`
}

// Operation
type ProductionRoutingHeaderKey struct {
	ProductionPlantBusinessPartner     []*int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerTo   *int      `json:"ProductionPlantBusinessPartnerTo"`
	ProductionPlantBusinessPartnerFrom *int      `json:"ProductionPlantBusinessPartnerFrom"`
	ProductionPlant                    []*string `json:"ProductionPlant"`
	ProductionPlantTo                  *string   `json:"ProductionPlantTo"`
	ProductionPlantFrom                *string   `json:"ProductionPlantFrom"`
	IsMarkedForDeletion                bool      `json:"IsMarkedForDeletion"`
}

type ProductionRoutingHeader struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	ProductionRoutingGroup        string  `json:"ProductionRoutingGroup"`
	ProductionRouting             string  `json:"ProductionRouting"`
	ProductionRoutingInternalVers string  `json:"ProductionRoutingInternalVers"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
	BillOfOperationsDesc          *string `json:"BillOfOperationsDesc"`
	Plant                         *string `json:"Plant"`
	BillOfOperationsUsage         *string `json:"BillOfOperationsUsage"`
	BillOfOperationsStatus        *string `json:"BillOfOperationsStatus"`
	ResponsiblePlannerGroup       *string `json:"ResponsiblePlannerGroup"`
	MinimumLotSizeQuantity        *string `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity        *string `json:"MaximumLotSizeQuantity"`
	BillOfOperationsUnit          *string `json:"BillOfOperationsUnit"`
	CreationDate                  *string `json:"CreationDate"`
	LastChangeDate                *string `json:"LastChangeDate"`
	ValidityStartDate             *string `json:"ValidityStartDate"`
	ValidityEndDate               *string `json:"ValidityEndDate"`
	ChangeNumber                  *string `json:"ChangeNumber"`
	PlainLongText                 *string `json:"PlainLongText"`
	MaterialAssignment            *string `json:"MaterialAssignment"`
}

type ProductionRoutingProductPlantKey struct {
	ProductInHeader        []*string `json:"ProductInHeader"`
	ProductInHeaderTo      *string   `json:"ProductInHeaderTo"`
	ProductInHeaderFrom    *string   `json:"ProductInHeaderFrom"`
	BusinessPartner        []int     `json:"BusinessPartner"`
	Plant                  []string  `json:"Plant"`
	ProductionRoutingGroup []string  `json:"ProductionRoutingGroup"`
	ProductionRouting      []string  `json:"ProductionRouting"`
	IsMarkedForDeletion    bool      `json:"IsMarkedForDeletion"`
}

type ProductionRoutingProductPlant struct {
	BusinessPartner                int     `json:"BusinessPartner"`
	Product                        string  `json:"Product"`
	Plant                          string  `json:"Plant"`
	ProductionRoutingGroup         string  `json:"ProductionRoutingGroup"`
	ProductionRouting              string  `json:"ProductionRouting"`
	ProductionRoutingMatlAssgmt    string  `json:"ProductionRoutingMatlAssgmt"`
	ProductionRtgMatlAssgmtIntVers string  `json:"ProductionRtgMatlAssgmtIntVers"`
	CreationDate                   *string `json:"CreationDate"`
	LastChangeDate                 *string `json:"LastChangeDate"`
	ValidityStartDate              *string `json:"ValidityStartDate"`
	ValidityEndDate                *string `json:"ValidityEndDate"`
	ChangeNumber                   *string `json:"ChangeNumber"`
}

type ProductionRoutingOperationKey struct {
	ProductionRoutingGroup []string `json:"ProductionRoutingGroup"`
	ProductionRouting      []string `json:"ProductionRouting"`
	Plant                  []string `json:"Plant"`
}

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

// 日付等の処理
type CreationDate struct {
	CreationDate string `json:"CreationDate"`
}

type LastChangeDate struct {
	LastChangeDate string `json:"LastChangeDate"`
}
