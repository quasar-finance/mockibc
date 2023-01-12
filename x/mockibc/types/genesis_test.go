package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"mockibc/x/mockibc/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated receivedPacket",
			genState: &types.GenesisState{
				ReceivedPacketList: []types.ReceivedPacket{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid receivedPacket count",
			genState: &types.GenesisState{
				ReceivedPacketList: []types.ReceivedPacket{
					{
						Id: 1,
					},
				},
				ReceivedPacketCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
