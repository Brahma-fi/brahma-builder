package morpho

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
)

const (
	slippage = 0.9995
)

type Config struct {
	FeeReceiver       string            `json:"feeReceiver"`
	BaseURL           string            `json:"baseURL"`
	BaseFeesInUSD     float64           `json:"baseFeesInUSD"`
	YieldFees         float64           `json:"yieldFees"`
	BundlerAddress    string            `json:"bundlerAddress"`
	FeeConfig         map[string]string `json:"feeConfig"`
	WhitelistedVaults []string          `json:"whitelistedVaults"`
}

func ParseConfig(raw map[string]any) (*Config, error) {
	cfg := &Config{}
	if err := mapstructure.Decode(raw, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

type StrategyParams struct {
	BaseToken common.Address `json:"baseToken"`
}
