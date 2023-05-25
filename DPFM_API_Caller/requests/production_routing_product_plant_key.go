package requests

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
