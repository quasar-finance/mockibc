package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendAcknowledgment = "send_acknowledgment"

var _ sdk.Msg = &MsgSendAcknowledgment{}

func NewMsgSendAcknowledgment(sender string) *MsgSendAcknowledgment {
	return &MsgSendAcknowledgment{
		Sender: sender,
	}
}

func (msg *MsgSendAcknowledgment) Route() string {
	return RouterKey
}

func (msg *MsgSendAcknowledgment) Type() string {
	return TypeMsgSendAcknowledgment
}

func (msg *MsgSendAcknowledgment) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendAcknowledgment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendAcknowledgment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
