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

func (k msgServer) Bid(ctx context.Context, msg *types.MsgBid) (res *types.MsgBidResponse, err error) {
	err = msg.ValidateBasic()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Key not valid")
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	bidsQueue, found := k.GetBidsQueue(sdkCtx, msg.AuctionId)
	if !found {
		return nil, fmt.Errorf("Bids queue for auction with id %v", msg.AuctionId)
	}

	for i, bid := range bidsQueue.Queue {
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

	return &types.MsgBidResponse{}, nil
}
