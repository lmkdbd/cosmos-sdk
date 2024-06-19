package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/testmodule/types"
)

func (k *Keeper) EndBlocker(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	new_params := types.Params{
		TestParam: "test",
	}
	_ = k.Params.Set(sdkCtx, new_params)
	// if err != nil {
	// 	return err
	// }
	return nil
}
