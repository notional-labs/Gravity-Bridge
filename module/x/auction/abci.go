package auction

import (
	"fmt"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction/keeper"
	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

func startMewAuctionPeriod(ctx sdk.Context, params types.Params, k keeper.Keeper, bk types.BankKeeper, ak types.AccountKeeper) error {
	auctionRate := params.AuctionRate

	increamentId, err := k.IncreamentAuctionPeriodId(ctx)
	if err != nil {
		panic(err)
	}

	newAuctionPeriods := types.AuctionPeriod{
		Id:               increamentId,
		StartBlockHeight: uint64(ctx.BlockHeight()),
		EndBlockHeight:   uint64(ctx.BlockHeight()) + params.AuctionPeriod,
	}

	// Set new auction period to store
	k.SetAuctionPeriod(ctx, newAuctionPeriods)

	for token := range params.AllowTokens {
		balance := bk.GetBalance(ctx, ak.GetModuleAccount(ctx, distrtypes.ModuleName).GetAddress(), token)

		// Compute auction amount to send to auction module account
		amount := sdk.NewDecFromInt(balance.Amount).Mul(auctionRate).TruncateInt()

		sdkcoin := sdk.NewCoin(token, amount)

		// Send fund from community pool to auction module
		err := k.SendFromCommunityPool(ctx, sdk.Coins{sdkcoin})
		if err != nil {
			return err
		}
		newId, err := k.IncreamentAuctionId(ctx, increamentId)
		if err != nil {
			return err
		}

		newAuction := types.Auction{
			Id:              newId,
			AuctionPeriodId: increamentId,
			AuctionAmount:   &sdkcoin,
			Status:          1,
			HighestBid:      nil,
		}

		// Update auction in auction period auction list
		err = k.AddNewAuctionToAuctionPeriod(ctx, increamentId, newAuction)
		if err != nil {
			return err
		}

		k.CreateNewBidQueue(ctx, newId)
	}

	return nil

}

func endAuctionPeriod(
	ctx sdk.Context,
	params types.Params,
	latestAuctionPeriod types.AuctionPeriod,
	k keeper.Keeper,
	bk types.BankKeeper,
	ak types.AccountKeeper,
) error {
	for _, auction := range k.GetAllAuctionsByPeriodID(ctx, latestAuctionPeriod.Id) {
		// Update auction status to finished
		auction.Status = 2
		k.SetAuction(ctx, auction)

		// If no bid return fund back to community pool
		if auction.HighestBid == nil {
			err := k.SendToCommunityPool(ctx, sdk.Coins{*auction.AuctionAmount})
			if err != nil {
				panic(err)
			}
			continue
		}

		// Send in the winning token to the highest bidder address
		err := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(auction.HighestBid.BidderAddress), sdk.Coins{*auction.AuctionAmount})
		if err != nil {
			panic(err)
		}

		// Delete the bid queue when the auction period has ended
		k.ClearQueue(ctx, auction.Id)
	}

	balances := bk.GetAllBalances(ctx, ak.GetModuleAccount(ctx, types.ModuleName).GetAddress())

	// Empty the rest of the auction module balances back to community pool
	err := k.SendFromCommunityPool(ctx, balances)
	if err != nil {
		panic(err)
	}
	return nil
}

// Go through all bid entries of auctions
// get the bid with highest amount and lock token from respective bidder
// and return lock token to bidder if someone bid a higher amount
func processBidEntries(
	ctx sdk.Context,
	params types.Params,
	k keeper.Keeper,
	latestAuctionPeriod types.AuctionPeriod,
) {
	for _, auction := range k.GetAllAuctionsByPeriodID(ctx, latestAuctionPeriod.Id) {
		bidsQueue, found := k.GetBidsQueue(ctx, auction.Id)
		if !found {
			continue
		}

		// Get new highest bid from bids queue
		// TODO: Need find better way to implement this
		oldHighestBid := auction.HighestBid
		var newHighestBid types.Bid
		if oldHighestBid == nil {
			newHighestBid, found = findHighestBid(ctx, bidsQueue)
			if !found {
				continue
			}
		} else {
			newHighestBid, found = findHighestBidCompareWithOldHighestBid(ctx, bidsQueue, *oldHighestBid)
			if !found {
				continue
			}
		}

		if oldHighestBid != nil && oldHighestBid.BidderAddress == newHighestBid.BidderAddress {
			bidAmountGap := newHighestBid.BidAmount.Sub(*oldHighestBid.BidAmount)
			// Send the added amount to auction module
			err := k.LockBidAmount(ctx, newHighestBid.BidderAddress, bidAmountGap)
			if err != nil {
				panic(fmt.Sprintf("Fail to lock bid token from address %s", newHighestBid.BidderAddress))
			}
		} else {
			// Return fund to the pervious highest bidder
			err := k.ReturnPrevioudBidAmount(ctx, oldHighestBid.BidderAddress, *oldHighestBid.BidAmount)
			if err != nil {
				panic(fmt.Sprintf("Fail to return lock token to address %s", oldHighestBid.BidderAddress))
			}

			err = k.LockBidAmount(ctx, newHighestBid.BidderAddress, *newHighestBid.BidAmount)
			if err != nil {
				panic(fmt.Sprintf("Fail to lock bid token from address %s", newHighestBid.BidderAddress))
			}

		}
		auction.HighestBid = &newHighestBid

		// Update the new highest bid entry
		k.SetAuction(ctx, auction)
	}
}

func findHighestBidCompareWithOldHighestBid(ctx sdk.Context, bidsQueue types.BidsQueue, highestBid types.Bid) (bid types.Bid, found bool) {
	// Set initial highest bid
	newHighestBid := highestBid
	found = false

	for _, bid := range bidsQueue.Queue {
		if !bid.BidAmount.IsLT(*newHighestBid.BidAmount) {
			newHighestBid = *bid
			found = true
		}
	}

	return newHighestBid, found
}

func findHighestBid(ctx sdk.Context, bidsQueue types.BidsQueue) (bid types.Bid, found bool) {
	// Set initial highest bid
	newHighestBid := *bidsQueue.Queue[0]
	found = false

	for _, bid := range bidsQueue.Queue {
		if !bid.BidAmount.IsLT(*newHighestBid.BidAmount) {
			newHighestBid = *bid
			found = true
		}
	}

	return newHighestBid, found
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper, bk types.BankKeeper, ak types.AccountKeeper) {
	params := k.GetParams(ctx)

	// An initial estimateNextBlockHeight need to be set as a starting point
	estimateNextBlockHeight, found := k.GetEstimateAuctionPeriodBlockHeight(ctx)
	if !found {
		panic("Cannot find estimate block height for next auction period")
	}

	if uint64(ctx.BlockHeight()) == estimateNextBlockHeight.Height {
		// Set estimate block height for next auction periods
		k.SetEstimateAuctionPeriodBlockHeight(ctx, uint64(ctx.BlockHeight())+params.AuctionEpoch)

		err := startMewAuctionPeriod(ctx, params, k, bk, ak)
		if err != nil {
			return
		}
	}
}

func EndBlocker(ctx sdk.Context, k keeper.Keeper, bk types.BankKeeper, ak types.AccountKeeper) {
	params := k.GetParams(ctx)

	lastestAuctionPeriods, found := k.GetLatestAuctionPeriod(ctx)
	if !found {
		return
	}

	// Process bid entries during the duration of the auction period
	if lastestAuctionPeriods.EndBlockHeight >= uint64(ctx.BlockHeight()) && lastestAuctionPeriods.StartBlockHeight <= uint64(ctx.BlockHeight()) {
		processBidEntries(ctx, params, k, *lastestAuctionPeriods)
	}

	if lastestAuctionPeriods.EndBlockHeight == uint64(ctx.BlockHeight()) {
		err := endAuctionPeriod(ctx, params, *lastestAuctionPeriods, k, bk, ak)
		if err != nil {
			return
		}
	}
}
