package requests

type ProductionRoutingOperationKey struct {
	ProductionRoutingGroup []string `json:"ProductionRoutingGroup"`
	ProductionRouting      []string `json:"ProductionRouting"`
	Plant                  []string `json:"Plant"`
}
