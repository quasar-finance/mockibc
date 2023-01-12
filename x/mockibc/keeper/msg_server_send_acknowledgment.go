package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mockibc/x/mockibc/types"
)

func (k msgServer) SendAcknowledgment(goCtx context.Context, msg *types.MsgSendAcknowledgment) (*types.MsgSendAcknowledgmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendAcknowledgmentResponse{}, nil
}
