package requests

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
