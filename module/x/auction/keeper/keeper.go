package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	// NOTE: If you add anything to this struct, add a nil check to ValidateMembers below!
	storeKey   sdk.StoreKey // Unexposed key to access store from sdk.Context
	paramSpace paramtypes.Subspace

	// NOTE: If you add anything to this struct, add a nil check to ValidateMembers below!
	cdc           codec.BinaryCodec // The wire codec for binary encoding/decoding.
	BankKeeper    *bankkeeper.BaseKeeper
	AccountKeeper *authkeeper.AccountKeeper
	DistKeeper    *distrkeeper.Keeper
}

// GetParams get params
func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.ParamsKey))
	params := types.Params{}
	err := k.cdc.Unmarshal(bz, &params)
	return params, err
}

// SetParams set params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set([]byte(types.ParamsKey), bz)
	return nil
}

// SendToCommunityPool handles incorrect SendToCosmos calls to the community pool, since the calls
// have already been made on Ethereum there's nothing we can do to reverse them, and we should at least
// make use of the tokens which would otherwise be lost
func (k Keeper) SendToCommunityPool(ctx sdk.Context, coins sdk.Coins) error {
	if err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, distrtypes.ModuleName, coins); err != nil {
		return sdkerrors.Wrap(err, "transfer to community pool failed")
	}
	feePool := k.DistKeeper.GetFeePool(ctx)
	feePool.CommunityPool = feePool.CommunityPool.Add(sdk.NewDecCoinsFromCoins(coins...)...)
	k.DistKeeper.SetFeePool(ctx, feePool)
	return nil
}

func (k Keeper) SendFromCommunityPool(ctx sdk.Context, coins sdk.Coins) error {
	if err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, distrtypes.ModuleName, types.ModuleName, coins); err != nil {
		return sdkerrors.Wrap(err, "transfer to auction module failed")
	}
	return nil
}
