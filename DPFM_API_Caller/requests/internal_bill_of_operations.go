package requests

type InternalBillOfOperations struct {
	PlannedOrder             int  `json:"PlannedOrder"`
	PlannedOrderItem         int  `json:"PlannedOrderItem"`
	InternalBillOfOperations *int `json:"InternalBillOfOperations"`
}
