package types

import (
	tmcrypto "github.com/tendermint/tendermint/crypto"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
)

var _ clientexported.Header = Header{}

// ClientType defines that the Header is a Solo Machine.
func (Header) ClientType() clientexported.ClientType {
	return clientexported.SoloMachine
}

// GetHeight returns the current sequence number as the height.
func (h Header) GetHeight() uint64 {
	return h.Sequence
}

// GetPubKey unmarshals the new public key into a tmcrypto.PubKey type.
func (h Header) GetPubKey() tmcrypto.PubKey {
	publicKey, ok := h.NewPublicKey.GetCachedValue().(tmcrypto.PubKey)
	if !ok {
		return nil
	}

	return publicKey
}

// ValidateBasic ensures that the sequence, signature and public key have all
// been initialized.
func (h Header) ValidateBasic() error {
	if h.Sequence == 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "sequence number cannot be zero")
	}

	if len(h.Signature) == 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "signature cannot be empty")
	}

	if h.NewPublicKey == nil || h.GetPubKey() == nil || len(h.GetPubKey().Bytes()) == 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "new public key cannot be empty")
	}

	return nil
}