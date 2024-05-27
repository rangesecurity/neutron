package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/neutron-org/neutron/v4/x/dynamicfees/types"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"
)

var _ feemarkettypes.DenomResolver = Keeper{}

func (k Keeper) ConvertToDenom(ctx sdk.Context, coin sdk.DecCoin, denom string) (sdk.DecCoin, error) {
	params := k.GetParams(ctx)
	for _, c := range params.NtrnPrices {
		if c.Denom == denom {
			return sdk.NewDecCoinFromDec(denom, coin.Amount.Mul(c.Amount)), nil
		}
	}
	return sdk.DecCoin{}, types.ErrUnknownDenom
}

func (k Keeper) AllowedDenoms(ctx sdk.Context) ([]string, error) {
	params := k.GetParams(ctx)
	denoms := make([]string, 0, params.NtrnPrices.Len())
	for _, coin := range params.NtrnPrices {
		denoms = append(denoms, coin.Denom)
	}
	return denoms, nil
}