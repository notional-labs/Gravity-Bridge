package auction_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"testing"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction"
	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/auction/types"

	"github.com/stretchr/testify/suite"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/app/apptesting"
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
		// fmt.Printf("%v \n", token)
		balance := suite.App.GetBankKeeper().GetBalance(ctx, suite.App.GetAccountKeeper().GetModuleAccount(ctx, distrtypes.ModuleName).GetAddress(), token)
		coins_dist = append(coins_dist, balance)

	}
	fmt.Printf("param: %v\n", params)
	fmt.Printf("coin dist module begin:%v \n", coins_dist)

	// set a Auction finish (Auction has ended.)
	CoinAuction := sdk.NewCoin("atomm", sdk.NewIntFromUint64(0))
	auctionPeriod_Set := types.AuctionPeriod{Id: 1, StartBlockHeight: 0, EndBlockHeight: 3}
	auction_Set := types.Auction{
		Id:            1,
		AuctionAmount: &CoinAuction,
		Status:        types.AuctionStatus_AUCTION_STATUS_FINISH,
		// nolint: exhaustruct
		HighestBid:      &types.Bid{AuctionId: 1, BidAmount: &CoinAuction},
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

	testCases := map[string]struct {
		ctxHeight             int64
		expectPanic           bool
		expectNewAuction      bool
		previousAuctionPeriod *types.AuctionPeriod
		communityBalances     sdk.Coins
	}{
		"Not meet the next auction period": {
			ctxHeight:        4,
			expectPanic:      false,
			expectNewAuction: false,
		},
		"Meet the next auction period, no previous auction period": {
			ctxHeight:   5,
			expectPanic: true,
		},
		"Meet the next auction period, community pool has zero balances": {
			ctxHeight:             5,
			expectPanic:           false,
			expectNewAuction:      false,
			previousAuctionPeriod: &previousAuctionPeriod,
		},
		"Meet the next auction period, community pool balances truncate to zero": {
			ctxHeight:             5,
			expectPanic:           false,
			expectNewAuction:      false,
			previousAuctionPeriod: &previousAuctionPeriod,
			communityBalances:     sdk.NewCoins(sdk.NewCoin("atom", sdk.NewInt(4))),
		},
		"Meet the next auction period, create new auction period": {
			ctxHeight:             5,
			expectPanic:           false,
			expectNewAuction:      true,
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
				feePool := suite.App.GetDistriKeeper().GetFeePool(ctx)
				fmt.Println("feePool", feePool)
			}

			if !tc.expectPanic {
				suite.Require().NotPanics(func() {
					auction.BeginBlocker(ctx, suite.App.GetAuctionKeeper(), suite.App.GetBankKeeper(), suite.App.GetAccountKeeper())
				})
				if tc.previousAuctionPeriod != nil {
					if tc.expectNewAuction {
						auctions := suite.App.GetAuctionKeeper().GetAllAuctionsByPeriodID(ctx, tc.previousAuctionPeriod.Id+1)
						// Should contain 1 aution for atom token
						suite.Equal(len(auctions), 1)
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
