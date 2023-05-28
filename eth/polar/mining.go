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

package polar

import (
	"context"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// TODO: replace this file with a proper mining object

// Prepare prepares the Polaris for processing a new block at the given height.
func (pl *Polaris) Prepare(ctx context.Context, height uint64) {
	pl.blockchain.Prepare(ctx, height)
}

// ProcessTransaction processes the given transaction and returns the receipt.
func (pl *Polaris) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*core.ExecutionResult, error) {
	return pl.blockchain.ProcessTransaction(ctx, tx)
}

// Finalize finalizes the current block.
func (pl *Polaris) Finalize(ctx context.Context) error {
	return pl.blockchain.Finalize(ctx)
}

func (pl *Polaris) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return pl.blockchain.SendTx(ctx, signedTx)
}
