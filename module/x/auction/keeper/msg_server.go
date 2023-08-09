package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the gov MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Bid msg add a bid entry to the queue to be process by the end of each block
func (k msgServer) Bid(ctx context.Context, msg *types.MsgBid) (res *types.MsgBidResponse, err error) {
	err = msg.ValidateBasic()

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Key not valid")
	}

	if msg.Amount.Amount.Uint64() < params.MinBidAmount {
		return nil, types.ErrInvalidBidAmount
	}

	bidsQueue, found := k.GetBidsQueue(sdkCtx, msg.AuctionId)
	if !found {
		return nil, fmt.Errorf("Bids queue for auction with id %v", msg.AuctionId)
	}

	// Fetch current auction period
	latestAuctionPeriod, found := k.GetLatestAuctionPeriod(sdkCtx)
	if !found {
		return nil, types.ErrNoPreviousAuctionPeriod
	}

	currentAuction, found := k.GetAuctionByPeriodIDAndAuctionId(sdkCtx, latestAuctionPeriod.Id, msg.AuctionId)
	if !found {
		return nil, types.ErrAuctionNotFound
	}
	highestBid := currentAuction.HighestBid

	// If highest bid exist need to check the bid gap
	if highestBid != nil && (msg.Amount.Sub(*highestBid.BidAmount)).Amount.Uint64() < params.BidGap {
		return nil, types.ErrInvalidBidAmountGap
	}

	if len(bidsQueue.Queue) == 0 {
		// For empty queue just add the new entry
		newBid := &types.Bid{
			AuctionId:     msg.AuctionId,
			BidAmount:     msg.Amount,
			BidderAddress: msg.Bidder,
		}
		k.AddBidToQueue(sdkCtx, *newBid, &bidsQueue)
	} else {
		for i, bid := range bidsQueue.Queue {
			// Check if bid entry from exact bidder exist yet
			if bid.AuctionId == msg.AuctionId && bid.BidderAddress == msg.Bidder {
				// Update bid amount of old entry
				bid.BidAmount = msg.Amount

				bidsQueue.Queue[i] = bid

				k.SetBidsQueue(sdkCtx, bidsQueue, msg.AuctionId)
			} else {
				newBid := &types.Bid{
					AuctionId:     msg.AuctionId,
					BidAmount:     msg.Amount,
					BidderAddress: msg.Bidder,
				}
				k.AddBidToQueue(sdkCtx, *newBid, &bidsQueue)
			}
		}
	}

	return &types.MsgBidResponse{}, nil
}
