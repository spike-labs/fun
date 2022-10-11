package request

type DiversifyFundsService struct {
	ContractAddress string `json:"contract_address" form:"contract_address"`
	MinBalance      string `json:"min_balance" form:"min_balance"`
}
