package mockibc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mockibc/testutil/keeper"
	"mockibc/testutil/nullify"
	"mockibc/x/mockibc"
	"mockibc/x/mockibc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		ReceivedPacketList: []types.ReceivedPacket{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ReceivedPacketCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MockibcKeeper(t)
	mockibc.InitGenesis(ctx, *k, genesisState)
	got := mockibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.ReceivedPacketList, got.ReceivedPacketList)
	require.Equal(t, genesisState.ReceivedPacketCount, got.ReceivedPacketCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
