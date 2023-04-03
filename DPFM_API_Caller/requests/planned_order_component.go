package requests

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
