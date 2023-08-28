package auction_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"reflect"
	"testing"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction"
	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction/types"

	"github.com/stretchr/testify/suite"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/app/apptesting"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

type TestSuite struct {
	apptesting.AppTestHelper
	suite.Suite
}

// Test helpers
func (suite *TestSuite) SetupTest() {
	suite.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestBeginBlockerAndEndBlockerAuction() {
	suite.SetupTest()
	ctx := suite.Ctx
	// params set
	defaultAuctionEpoch := uint64(4)
	defaultAuctionPeriod := uint64(2)
	defaultMinBidAmount := uint64(1000)
	defaultBidGap := uint64(100)
	auctionRate := sdk.NewDecWithPrec(2, 1) // 20%
	allowTokens := map[string]bool{
		"atomm": true,
	}
	params := types.NewParams(defaultAuctionEpoch, defaultAuctionPeriod, defaultMinBidAmount, defaultBidGap, auctionRate, allowTokens)
	suite.App.GetAuctionKeeper().SetParams(ctx, params)

	// set community pool
	coinsSet := []sdk.Coin{}
	for token := range params.AllowTokens {
		sdkcoin := sdk.NewCoin(token, sdk.NewIntFromUint64(1_000_000_000))
		coinsSet = append(coinsSet, sdkcoin)

	}

	suite.FundModule(ctx, distrtypes.ModuleName, coinsSet)

	coins_dist := []sdk.Coin{}
	for token := range params.AllowTokens {
		balance := suite.App.GetBankKeeper().GetBalance(ctx, suite.App.GetAccountKeeper().GetModuleAccount(ctx, distrtypes.ModuleName).GetAddress(), token)
		coins_dist = append(coins_dist, balance)

	}

	// set a Auction finish (Auction has ended.)
	CoinAuction := sdk.NewCoin("atomm", sdk.NewIntFromUint64(0))
	auctionPeriod_Set := types.AuctionPeriod{Id: 1, StartBlockHeight: 0, EndBlockHeight: 3}
	auction_Set := types.Auction{
		Id:            1,
		AuctionAmount: CoinAuction,
		Status:        types.AuctionStatus_AUCTION_STATUS_FINISH,
		// nolint: exhaustruct
		HighestBid:      &types.Bid{AuctionId: 1, BidAmount: CoinAuction},
		AuctionPeriodId: auctionPeriod_Set.Id,
	}
	suite.App.GetAuctionKeeper().SetAuctionPeriod(ctx, auctionPeriod_Set)
	err := suite.App.GetAuctionKeeper().AddNewAuctionToAuctionPeriod(ctx, auctionPeriod_Set.Id, auction_Set)
	suite.Require().NoError(err)

	println("============================begin block=================================")
	suite.App.GetAuctionKeeper().SetEstimateAuctionPeriodBlockHeight(ctx, uint64(ctx.BlockHeight()))

	auction.BeginBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())

	coins_auc := []sdk.Coin{}
	for token := range params.AllowTokens {
		balance := suite.App.GetBankKeeper().GetBalance(ctx, suite.App.GetAccountKeeper().GetModuleAccount(ctx, types.ModuleName).GetAddress(), token)
		coins_auc = append(coins_auc, balance)

	}
	fmt.Printf("coin auction module mid:%v \n", coins_auc)

	coins_new := []sdk.Coin{}
	for token := range params.AllowTokens {
		balance := suite.App.GetBankKeeper().GetBalance(ctx, suite.App.GetAccountKeeper().GetModuleAccount(ctx, distrtypes.ModuleName).GetAddress(), token)
		coins_new = append(coins_new, balance)

	}
	fmt.Printf("coin dist module mid:%v \n", coins_new)
	println("============================end block=============================")
	ctx = ctx.WithBlockHeight(3)
	auction.EndBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())

	coins_auc = []sdk.Coin{}
	for token := range params.AllowTokens {
		balance := suite.App.GetBankKeeper().GetBalance(ctx, suite.App.GetAccountKeeper().GetModuleAccount(ctx, types.ModuleName).GetAddress(), token)
		coins_auc = append(coins_auc, balance)

	}
	fmt.Printf("coin auction module end:%v \n", coins_auc)

	coins_new = []sdk.Coin{}
	for token := range params.AllowTokens {
		balance := suite.App.GetBankKeeper().GetBalance(ctx, suite.App.GetAccountKeeper().GetModuleAccount(ctx, distrtypes.ModuleName).GetAddress(), token)
		coins_new = append(coins_new, balance)

	}
	fmt.Printf("coin dist module end:%v \n", coins_new)
}

func (suite *TestSuite) TestBeginBlocker() {
	previousAuctionPeriod := types.AuctionPeriod{Id: 1, StartBlockHeight: 0, EndBlockHeight: 4}
	expectAmount := sdk.NewCoin("atom", sdk.NewInt(20_000_000))

	testCases := map[string]struct {
		ctxHeight             int64
		expectPanic           bool
		expectAuction         types.Auction
		previousAuctionPeriod *types.AuctionPeriod
		communityBalances     sdk.Coins
	}{
		"Not meet the next auction period": {
			ctxHeight:   4,
			expectPanic: false,
		},
		"Meet the next auction period, no previous auction period": {
			ctxHeight:   5,
			expectPanic: true,
		},
		"Meet the next auction period, community pool has zero balances": {
			ctxHeight:             5,
			expectPanic:           false,
			previousAuctionPeriod: &previousAuctionPeriod,
		},
		"Meet the next auction period, community pool balances truncate to zero": {
			ctxHeight:             5,
			expectPanic:           false,
			previousAuctionPeriod: &previousAuctionPeriod,
			communityBalances:     sdk.NewCoins(sdk.NewCoin("atom", sdk.NewInt(4))),
		},
		"Meet the next auction period, create new auction period": {
			ctxHeight:   5,
			expectPanic: false,
			expectAuction: types.Auction{
				Id:              1,
				AuctionAmount:   expectAmount,
				Status:          types.AuctionStatus_AUCTION_STATUS_IN_PROGRESS,
				AuctionPeriodId: 2,
			},
			previousAuctionPeriod: &previousAuctionPeriod,
			communityBalances:     sdk.NewCoins(sdk.NewCoin("atom", sdk.NewInt(100_000_000))),
		},
	}

	for name, tc := range testCases {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx

			// Set params
			allowTokens := map[string]bool{
				"atom": true,
			}
			params := types.NewParams(uint64(4), uint64(2), uint64(1000), uint64(100), sdk.NewDecWithPrec(2, 1), allowTokens)
			suite.App.GetAuctionKeeper().SetParams(ctx, params)

			// Try to begin block without initial estimateNextBlockHeight set
			suite.Require().Panics(func() {
				auction.BeginBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())
			})

			// Set next auction period at block 5
			suite.App.GetAuctionKeeper().SetEstimateAuctionPeriodBlockHeight(ctx, 5)

			ctx = ctx.WithBlockHeight(tc.ctxHeight)

			if tc.previousAuctionPeriod != nil {
				suite.App.GetAuctionKeeper().SetAuctionPeriod(ctx, *tc.previousAuctionPeriod)
			}

			if tc.communityBalances != nil {
				suite.FundModule(ctx, distrtypes.ModuleName, tc.communityBalances)
				suite.App.GetDistriKeeper().SetFeePool(ctx, distrtypes.FeePool{CommunityPool: sdk.NewDecCoinsFromCoins(tc.communityBalances...)})
			}

			if !tc.expectPanic {
				suite.Require().NotPanics(func() {
					auction.BeginBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())
				})
				if tc.previousAuctionPeriod != nil {
					if !reflect.DeepEqual(tc.expectAuction, types.Auction{}) {
						auctions := suite.App.GetAuctionKeeper().GetAllAuctionsByPeriodID(ctx, tc.previousAuctionPeriod.Id+1)
						// Should contain 1 aution for atom token
						suite.Equal(len(auctions), 1)
						auction := auctions[0]
						suite.Equal(auction, tc.expectAuction)

					} else {
						auctions := suite.App.GetAuctionKeeper().GetAllAuctionsByPeriodID(ctx, tc.previousAuctionPeriod.Id+1)
						// Should not cotain any aution
						suite.Equal(len(auctions), 0)
					}
				}
			} else {
				suite.Require().Panics(func() {
					auction.BeginBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())
				})
			}
		})
	}

}

func (suite *TestSuite) TestEndBlocker() {
	_, _, addr0 := testdata.KeyTestPubAddr()
	_, _, addr1 := testdata.KeyTestPubAddr()

	testCases := map[string]struct {
		ctxHeight         int64
		queue             types.BidsQueue
		expectPanic       bool
		currentHighestBid *types.Bid
		newHighestBid     *types.Bid
		expectedCommunity sdk.Coins
	}{
		"In the process duration, empty bid queue": {
			ctxHeight: 5,
			queue: types.BidsQueue{
				Queue: []*types.Bid{},
			},
			expectPanic: false,
		},
		"In the process duration, no highest bid existing, insufficient balances": {
			ctxHeight: 5,
			queue: types.BidsQueue{
				Queue: []*types.Bid{
					{
						AuctionId:     1,
						BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_100_000)),
						BidderAddress: suite.TestAccs[0].String(),
					},
				},
			},
			expectPanic: false,
			expectedCommunity: sdk.NewCoins(sdk.NewCoin("atom", sdk.NewInt(10_000_000))),
		},
		"In the process duration, no highest bid existing, enough balances": {
			ctxHeight: 5,
			queue: types.BidsQueue{
				Queue: []*types.Bid{
					{
						AuctionId:     1,
						BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_000_000)),
						BidderAddress: addr0.String(),
					},
				},
			},
			expectPanic: false,
			newHighestBid: &types.Bid{
				AuctionId:     1,
				BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_000_000)),
				BidderAddress: addr0.String(),
			},
		},
		"In the process duration, highest bid existing, new bid > highest": {
			ctxHeight: 5,
			currentHighestBid: &types.Bid{
				AuctionId:     1,
				BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_000_000)),
				BidderAddress: addr0.String(),
			},
			queue: types.BidsQueue{
				Queue: []*types.Bid{
					{
						AuctionId:     1,
						BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_100_000)),
						BidderAddress: addr0.String(),
					},
				},
			},
			expectPanic: false,
			newHighestBid: &types.Bid{
				AuctionId:     1,
				BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_100_000)),
				BidderAddress: addr0.String(),
			},
		},
		"In the process duration, highest bid existing, 2 new bids with same amount": {
			ctxHeight: 5,
			currentHighestBid: &types.Bid{
				AuctionId:     1,
				BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_000_000)),
				BidderAddress: addr0.String(),
			},
			queue: types.BidsQueue{
				Queue: []*types.Bid{
					{
						AuctionId:     1,
						BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_100_000)),
						BidderAddress: addr0.String(),
					},
					{
						AuctionId:     1,
						BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_100_000)),
						BidderAddress: addr1.String(),
					},
					
				},
			},
			expectPanic: false,
			newHighestBid: &types.Bid{
				AuctionId:     1,
				BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_100_000)),
				BidderAddress: addr0.String(),
			},
		},
		"In the process end, no winning": {
			ctxHeight: 7,
			queue: types.BidsQueue{
				Queue: []*types.Bid{},
			},
			expectPanic: false,
		},
		"In the process end, has winning": {
			ctxHeight: 7,
			currentHighestBid: &types.Bid{
				AuctionId:     1,
				BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_000_000)),
				BidderAddress: addr0.String(),
			},
			queue: types.BidsQueue{
				Queue: []*types.Bid{},
			},
			expectPanic: false,
			newHighestBid: &types.Bid{
				AuctionId:     1,
				BidAmount:     sdk.NewCoin("stake", sdk.NewInt(1_000_000)),
				BidderAddress: addr0.String(),
			},
		},
	}

	for name, tc := range testCases {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx

			// Set params
			allowTokens := map[string]bool{
				"atom": true,
			}
			params := types.NewParams(uint64(4), uint64(2), uint64(1000), uint64(100), sdk.NewDecWithPrec(2, 1), allowTokens)
			suite.App.GetAuctionKeeper().SetParams(ctx, params)

			// Fund module & account
			suite.FundModule(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1_000_000))))
			suite.FundModule(ctx, distrtypes.ModuleName, sdk.NewCoins(sdk.NewCoin("atom", sdk.NewInt(10_000_000))))
			suite.App.GetDistriKeeper().SetFeePool(ctx, distrtypes.FeePool{CommunityPool: sdk.NewDecCoinsFromCoins(sdk.NewCoin("atom", sdk.NewInt(10_000_000)))})
			suite.FundAccount(ctx, sdk.MustAccAddressFromBech32(addr0.String()), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(2_000_000))))

			// Set next auction period at block 5 and set up auction
			suite.App.GetAuctionKeeper().SetEstimateAuctionPeriodBlockHeight(ctx, 5)
			ctx = ctx.WithBlockHeight(5)
			previousAuctionPeriod := types.AuctionPeriod{Id: 1, StartBlockHeight: 0, EndBlockHeight: 4}
			suite.App.GetAuctionKeeper().SetAuctionPeriod(ctx, previousAuctionPeriod)
			auction.BeginBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())

			ctx = ctx.WithBlockHeight(tc.ctxHeight)
			if tc.queue.Queue != nil {
				suite.App.GetAuctionKeeper().SetBidsQueue(ctx, tc.queue, 1)
			}

			if tc.currentHighestBid != nil {
				auction, _ := suite.App.GetAuctionKeeper().GetAuctionByPeriodIDAndAuctionId(ctx, 2, 1)
				auction.HighestBid = tc.currentHighestBid
				suite.App.GetAuctionKeeper().SetAuction(ctx, auction)
			}

			if !tc.expectPanic {
				auction.EndBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())
				auction, found := suite.App.GetAuctionKeeper().GetAuctionByPeriodIDAndAuctionId(ctx, 2, 1)
				suite.Require().True(found)
				if tc.newHighestBid != nil {
					suite.Require().Equal(tc.newHighestBid, auction.HighestBid)
				} else {
					suite.Require().Nil(auction.HighestBid)
				}
			}

		})
	}

}
