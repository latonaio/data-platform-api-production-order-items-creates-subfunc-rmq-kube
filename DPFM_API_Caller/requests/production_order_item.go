package requests

type ProductionOrderItem struct {
	PlannedOrder              int `json:"PlannedOrder"`
	PlannedOrderItem          int `json:"PlannedOrderItem"`
	ProductionOrderItemNumber int `json:"ProductionOrderItemNumber"`
}
