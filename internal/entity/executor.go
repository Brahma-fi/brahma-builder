package entity

type SignAndExecuteRequest struct {
	Operation  uint8  `json:"operation"`
	Subaccount string `json:"subaccount"`
	ChainID    uint64 `json:"chainID"`
	To         string `json:"to"`
	Value      string `json:"value"`
	Data       string `json:"data"`
}
