package api_processing_data_formatter

type SDC struct {
	MetaData                        *MetaData                        `json:"MetaData"`
	ProcessType                     *ProcessType                     `json:"ProcessType"`
	PlannedOrderHeder               []*PlannedOrderHeder             `json:"PlannedOrderHeder"`
	PlannedOrderItem                []*PlannedOrderItem              `json:"PlannedOrderItem"`
	PlannedOrderComponent           []*PlannedOrderComponent         `json:"PlannedOrderComponent"`
	ProductionOrderItem             []*ProductionOrderItem           `json:"ProductionOrderItem"`
	ProductMasterBPPlant            []*ProductMasterBPPlant          `json:"ProductMasterBPPlant"`
	BatchMasterRecordBatch          []*BatchMasterRecordBatch        `json:"BatchMasterRecordBatch"`
	StockConfirmation               []*StockConfirmation             `json:"StockConfirmation"`
	ProductAvailabilityIsNotChecked *ProductAvailabilityIsNotChecked `json:"ProductAvailabilityIsNotChecked"`
	InternalBillOfOperations        []*InternalBillOfOperations      `json:"InternalBillOfOperations"`
	TotalQuantity                   []*TotalQuantity                 `json:"TotalQuantity"`
	PlannedScrapQuantityItem        []*PlannedScrapQuantityItem      `json:"PlannedScrapQuantityItem"`
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
type PlannedOrderHederKey struct {
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

type PlannedOrderHeder struct {
	PlannedOrder                             int     `json:"PlannedOrder"`
	PlannedOrderType                         *string `json:"PlannedOrderType"`
	Product                                  *string `json:"Product"`
	ProductDeliverFromParty                  *int    `json:"ProductDeliverFromParty"`
	ProductDeliverToParty                    *int    `json:"ProductDeliverToParty"`
	OriginIssuingPlant                       *string `json:"OriginIssuingPlant"`
	OriginIssuingPlantStorageLocation        *string `json:"OriginIssuingPlantStorageLocation"`
	DestinationReceivingPlant                *string `json:"DestinationReceivingPlant"`
	DestinationReceivingPlantStorageLocation *string `json:"DestinationReceivingPlantStorageLocation"`
	OwnerProductionPlantBusinessPartner      *int    `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     *string `json:"OwnerProductionPlant"`
	OwnerProductionPlantStorageLocation      *string `json:"OwnerProductionPlantStorageLocation"`
	// BaseUnit                                 *string  `json:"BaseUnit"`
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
	PlannedOrderSequence             *string  `json:"PlannedOrderSequence"`
	PlannedOrderOperation            *string  `json:"PlannedOrderOperation"`
	OrderInternalBillOfOperations    *string  `json:"OrderInternalBillOfOperations"`
	BillOfMaterial                   *int     `json:"BillOfMaterial"`
	BOMItem                          *int     `json:"BOMItem"`
	BOMItemDescription               *string  `json:"BOMItemDescription"`
	BillOfMaterialCategory           *string  `json:"BillOfMaterialCategory"`
	BillOfMaterialItemNumber         *int     `json:"BillOfMaterialItemNumber"`
	BillOfMaterialInternalID         *int     `json:"BillOfMaterialInternalID"`
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
	ProductionOrderItemNumber int `json:"ProductionOrderItemNumber"`
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
		BusinessPartner                 *int     `json:"BusinessPartner"`
		Product                         *string  `json:"Product"`
		Plant                           *string  `json:"Plant"`
		Batch                           *string  `json:"Batch"`
		RequestedQuantity               *float32 `json:"RequestedQuantity"`
		ProductStockAvailabilityDate    *string  `json:"ProductStockAvailabilityDate"`
		OrderID                         *int     `json:"OrderID"`
		OrderItem                       *int     `json:"OrderItem"`
		Project                         *string  `json:"Project"`
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

type ProductAvailabilityIsNotChecked struct {
	ProductAvailabilityIsNotChecked *bool `json:"ProductAvailabilityIsNotChecked"`
}

type InternalBillOfOperations struct {
	PlannedOrder             int     `json:"PlannedOrder"`
	PlannedOrderItem         int     `json:"PlannedOrderItem"`
	InternalBillOfOperations *string `json:"InternalBillOfOperations"`
}

type TotalQuantity struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Product                      string  `json:"Product"`
	Plant                        string  `json:"Plant"`
	Batch                        string  `json:"Batch"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	TotalQuantity                float32 `json:"TotalQuantity"`
}

type PlannedScrapQuantityItem struct {
	PlannedOrder            int     `json:"PlannedOrder"`
	PlannedOrderItem        int     `json:"PlannedOrderItem"`
	ComponentScrapInPercent float32 `json:"ComponentScrapInPercent"`
	TotalQuantity           float32 `json:"TotalQuantity"`
	PlannedScrapQuantity    float32 `json:"PlannedScrapQuantity"`
}

// 日付等の処理
type CreationDate struct {
	CreationDate string `json:"CreationDate"`
}

type LastChangeDate struct {
	LastChangeDate string `json:"LastChangeDate"`
}
