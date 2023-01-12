package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mockibc/x/mockibc/types"
)

func (k Keeper) ReceivedPacketAll(c context.Context, req *types.QueryAllReceivedPacketRequest) (*types.QueryAllReceivedPacketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var receivedPackets []types.ReceivedPacket
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	receivedPacketStore := prefix.NewStore(store, types.KeyPrefix(types.ReceivedPacketKey))

	pageRes, err := query.Paginate(receivedPacketStore, req.Pagination, func(key []byte, value []byte) error {
		var receivedPacket types.ReceivedPacket
		if err := k.cdc.Unmarshal(value, &receivedPacket); err != nil {
			return err
		}

		receivedPackets = append(receivedPackets, receivedPacket)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReceivedPacketResponse{ReceivedPacket: receivedPackets, Pagination: pageRes}, nil
}

func (k Keeper) ReceivedPacket(c context.Context, req *types.QueryGetReceivedPacketRequest) (*types.QueryGetReceivedPacketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	receivedPacket, found := k.GetReceivedPacket(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetReceivedPacketResponse{ReceivedPacket: receivedPacket}, nil
}
