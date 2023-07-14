package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAuctionPeriodByID returns the auction period with the given id.
func (k Keeper) GetAuctionPeriodByID(ctx sdk.Context, id uint64) (val types.AuctionPeriod, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.KeyPrefixAuctionPeriod))
	bz := store.Get(uint64ToBytes(id))
	if bz == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(bz, &val)
	return val, true
}

// GetAllAuctionPeriods returns all auction periods.
func (k Keeper) GetAllAuctionPeriods(ctx sdk.Context) []types.AuctionPeriod {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.KeyPrefixAuctionPeriod))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	var auctionPeriods []types.AuctionPeriod
	for ; iterator.Valid(); iterator.Next() {
		var auctionPeriod types.AuctionPeriod
		k.cdc.MustUnmarshal(iterator.Value(), &auctionPeriod)
		auctionPeriods = append(auctionPeriods, auctionPeriod)
	}

	return auctionPeriods
}

// GetAuctionPeriodsByAuctionId returns all auction periods.
func (k Keeper) GetAuctionPeriodsByAuctionId(ctx sdk.Context, auctionId uint64) []types.AuctionPeriod {
	auctions := k.GetAllAuctions(ctx)
	auctionPeriodIds := []uint64{}
	auctionPeriodsFound := []types.AuctionPeriod{}
	for _, auction := range auctions {
		if auction.Id == auctionId {
			auctionPeriodIds = append(auctionPeriodIds, auction.AuctionPeriodId)
		}
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.KeyPrefixAuctionPeriod))

	for _, auctionPeriodId := range auctionPeriodIds {
		var autionPeriod types.AuctionPeriod
		bz := store.Get(uint64ToBytes(auctionPeriodId))
		k.cdc.MustUnmarshal(bz, &autionPeriod)
		
		auctionPeriodsFound = append(auctionPeriodsFound, autionPeriod)
	}
	return auctionPeriodsFound
}

// GetLatestAuctionPeriod returns the latest auction period.
func (k Keeper) GetLatestAuctionPeriod(ctx sdk.Context) (*types.AuctionPeriod, bool) {
	auctionPeriods := k.GetAllAuctionPeriods(ctx)
	if len(auctionPeriods) == 0 {
		return nil, false
	}
	return &auctionPeriods[len(auctionPeriods)-1], true
}

// SetAuctionPeriod sets the given auction period.
func (k Keeper) SetAuctionPeriod(ctx sdk.Context, auctionPeriod types.AuctionPeriod) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.KeyPrefixAuctionPeriod))

	bz := k.cdc.MustMarshal(&auctionPeriod)
	store.Set(uint64ToBytes(auctionPeriod.Id), bz)
}

func (k Keeper) IncreamentAuctionPeriodId(ctx sdk.Context) (uint64, error) {
	lastAuctionPeriod, found := k.GetLatestAuctionPeriod(ctx)
	if !found {
		return 0, fmt.Errorf("An initial auction period must be set during upgrade handler")
	}
	return lastAuctionPeriod.Id + 1, nil
}

// Helper function to convert uint64 to bytes.
func uint64ToBytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}
