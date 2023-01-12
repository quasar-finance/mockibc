package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mockibc/x/mockibc/types"
)

// GetReceivedPacketCount get the total number of receivedPacket
func (k Keeper) GetReceivedPacketCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReceivedPacketCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetReceivedPacketCount set the total number of receivedPacket
func (k Keeper) SetReceivedPacketCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReceivedPacketCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendReceivedPacket appends a receivedPacket in the store with a new id and update the count
func (k Keeper) AppendReceivedPacket(
	ctx sdk.Context,
	receivedPacket types.ReceivedPacket,
) uint64 {
	// Create the receivedPacket
	count := k.GetReceivedPacketCount(ctx)

	// Set the ID of the appended value
	receivedPacket.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReceivedPacketKey))
	appendedValue := k.cdc.MustMarshal(&receivedPacket)
	store.Set(GetReceivedPacketIDBytes(receivedPacket.Id), appendedValue)

	// Update receivedPacket count
	k.SetReceivedPacketCount(ctx, count+1)

	return count
}

// SetReceivedPacket set a specific receivedPacket in the store
func (k Keeper) SetReceivedPacket(ctx sdk.Context, receivedPacket types.ReceivedPacket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReceivedPacketKey))
	b := k.cdc.MustMarshal(&receivedPacket)
	store.Set(GetReceivedPacketIDBytes(receivedPacket.Id), b)
}

// GetReceivedPacket returns a receivedPacket from its id
func (k Keeper) GetReceivedPacket(ctx sdk.Context, id uint64) (val types.ReceivedPacket, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReceivedPacketKey))
	b := store.Get(GetReceivedPacketIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveReceivedPacket removes a receivedPacket from the store
func (k Keeper) RemoveReceivedPacket(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReceivedPacketKey))
	store.Delete(GetReceivedPacketIDBytes(id))
}

// GetAllReceivedPacket returns all receivedPacket
func (k Keeper) GetAllReceivedPacket(ctx sdk.Context) (list []types.ReceivedPacket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReceivedPacketKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ReceivedPacket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetReceivedPacketIDBytes returns the byte representation of the ID
func GetReceivedPacketIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetReceivedPacketIDFromBytes returns ID in uint64 format from a byte array
func GetReceivedPacketIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
