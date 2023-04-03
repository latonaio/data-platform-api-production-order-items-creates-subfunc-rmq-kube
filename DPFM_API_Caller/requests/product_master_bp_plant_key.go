package requests

type ProductMasterBPPlantKey struct {
	ComponentProduct                 []string `json:"ComponentProduct"`
	StockConfirmationBusinessPartner []int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           []string `json:"StockConfirmationPlant"`
}
