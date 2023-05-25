package requests

type ProductionOrderComponent struct {
	PlannedOrder     int `json:"PlannedOrder"`
	PlannedOrderItem int `json:"PlannedOrderItem"`
	Operations       int `json:"Operations"`
	OperationsItem   int `json:"OperationsItem"`
	BillOfMaterial   int `json:"BillOfMaterial"`
	BOMItem          int `json:"BOMItem"`
}
