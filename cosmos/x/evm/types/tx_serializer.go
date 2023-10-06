// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/ethereum/go-ethereum/beacon/engine"

	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// TxSerializer provides an interface to serialize ethereum transactions
// to sdk.Tx's and bytes that can be used by CometBFT.
type TxSerializer interface {
	TxToSdkTx(signedTx *coretypes.Transaction) (sdk.Tx, error)
	TxToSdkTxBytes(signedTx *coretypes.Transaction) ([]byte, error)
	PayloadToSdkTxBytes(payload *engine.ExecutionPayloadEnvelope) ([]byte, error)
	PayloadToSdkTx(payload *engine.ExecutionPayloadEnvelope) (sdk.Tx, error)
}

// serializer implements TxSerializer.
type serializer struct {
	txConfig client.TxConfig
}

// NewSerializer returns a new instance of TxSerializer.
func NewSerializer(txConfig client.TxConfig) TxSerializer {
	return &serializer{
		txConfig: txConfig,
	}
}

// SerializeToBytes converts an Ethereum transaction to Cosmos formatted
// txBytes which allows for it to broadcast it to CometBFT.
func (s *serializer) PayloadToSdkTxBytes(
	payload *engine.ExecutionPayloadEnvelope,
) ([]byte, error) {
	// First, we convert the Ethereum transaction to a Cosmos transaction.
	cosmosTx, err := s.PayloadToSdkTx(payload)
	if err != nil {
		return nil, err
	}

	// Then we use the clientCtx.TxConfig.TxEncoder() to encode the Cosmos transaction into bytes.
	return s.txConfig.TxEncoder()(cosmosTx)
}

// PayloadToSdkTx converts an ExecutionPayloadEnvelope to an sdk.Tx.
func (s *serializer) PayloadToSdkTx(payload *engine.ExecutionPayloadEnvelope) (sdk.Tx, error) {
	var err error
	// TODO: do we really need to use extensions for anything? Since we
	// are using the standard ante handler stuff I don't think we actually need to.
	tx := s.txConfig.NewTxBuilder()

	// Set the tx gas limit to the block gas limit in the payload
	tx.SetGasLimit(payload.ExecutionPayload.GasLimit)

	bz, err := payload.MarshalJSON()
	if err != nil {
		return nil, err
	}

	wp := &WrappedPayloadEnvelope{
		Data: bz,
	}

	// Lastly, we set the signature. We can pull the sequence from the nonce of the ethereum tx.
	if err = tx.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: 0,
			Data: &signingtypes.SingleSignatureData{
				// We retrieve the hash of the signed transaction from the ethereum transaction
				// objects, as this was the bytes that were signed. We pass these into the
				// SingleSignatureData as the SignModeHandler needs to know what data was signed
				// over so that it can verify the signature in the ante handler.
				Signature: []byte{0x01},
			},
			PubKey: &secp256k1.PubKey{Key: []byte{0x01}},
		},
	); err != nil {
		return nil, err
	}

	// Lastly, we inject the signed ethereum transaction as a message into the Cosmos Tx.
	if err = tx.SetMsgs(wp); err != nil {
		return nil, err
	}

	// Finally, we return the Cosmos Tx.
	return tx.GetTx(), nil
}

// TxToSdkTx converts an ethereum transaction to a Cosmos native transaction.
func (s *serializer) TxToSdkTx(signedTx *coretypes.Transaction) (sdk.Tx, error) {
	var err error
	// are using the standard ante handler stuff I don't think we actually need to.
	tx := s.txConfig.NewTxBuilder()

	// Create the WrappedEthereumTransaction message.
	wrappedEthTx, err := WrapTx(signedTx)
	if err != nil {
		return nil, err
	}

	// Lastly, we set the signature. We can pull the sequence from the nonce of the ethereum tx.
	if err = tx.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: signedTx.Nonce(),
			Data: &signingtypes.SingleSignatureData{
				// We retrieve the hash of the signed transaction from the ethereum transaction
				// objects, as this was the bytes that were signed. We pass these into the
				// SingleSignatureData as the SignModeHandler needs to know what data was signed
				// over so that it can verify the signature in the ante handler.
				Signature: []byte{0x0},
			},
			PubKey: &secp256k1.PubKey{Key: []byte{0x0}},
		},
	); err != nil {
		return nil, err
	}

	// Lastly, we inject the signed ethereum transaction as a message into the Cosmos Tx.
	if err = tx.SetMsgs(wrappedEthTx); err != nil {
		return nil, err
	}

	// Finally, we return the Cosmos Tx.
	return tx.GetTx(), nil
}

// SerializeToBytes converts an Ethereum transaction to Cosmos formatted txBytes which allows for
// it to broadcast it to CometBFT.
func (s *serializer) TxToSdkTxBytes(signedTx *coretypes.Transaction) ([]byte, error) {
	// First, we convert the Ethereum transaction to a Cosmos transaction.
	cosmosTx, err := s.TxToSdkTx(signedTx)
	if err != nil {
		return nil, err
	}

	// Then we use the clientCtx.TxConfig.TxEncoder() to encode the Cosmos transaction into bytes.
	return s.txConfig.TxEncoder()(cosmosTx)
}
