package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "mockibc/testutil/keeper"
	"mockibc/testutil/nullify"
	"mockibc/x/mockibc/keeper"
	"mockibc/x/mockibc/types"
)

func createNReceivedPacket(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ReceivedPacket {
	items := make([]types.ReceivedPacket, n)
	for i := range items {
		items[i].Id = keeper.AppendReceivedPacket(ctx, items[i])
	}
	return items
}

func TestReceivedPacketGet(t *testing.T) {
	keeper, ctx := keepertest.MockibcKeeper(t)
	items := createNReceivedPacket(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetReceivedPacket(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestReceivedPacketRemove(t *testing.T) {
	keeper, ctx := keepertest.MockibcKeeper(t)
	items := createNReceivedPacket(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveReceivedPacket(ctx, item.Id)
		_, found := keeper.GetReceivedPacket(ctx, item.Id)
		require.False(t, found)
	}
}

func TestReceivedPacketGetAll(t *testing.T) {
	keeper, ctx := keepertest.MockibcKeeper(t)
	items := createNReceivedPacket(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllReceivedPacket(ctx)),
	)
}

func TestReceivedPacketCount(t *testing.T) {
	keeper, ctx := keepertest.MockibcKeeper(t)
	items := createNReceivedPacket(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetReceivedPacketCount(ctx))
}
