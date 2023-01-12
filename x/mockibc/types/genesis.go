package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:             PortID,
		ReceivedPacketList: []ReceivedPacket{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in receivedPacket
	receivedPacketIdMap := make(map[uint64]bool)
	receivedPacketCount := gs.GetReceivedPacketCount()
	for _, elem := range gs.ReceivedPacketList {
		if _, ok := receivedPacketIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for receivedPacket")
		}
		if elem.Id >= receivedPacketCount {
			return fmt.Errorf("receivedPacket id should be lower or equal than the last id")
		}
		receivedPacketIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
