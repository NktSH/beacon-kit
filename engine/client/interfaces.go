// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package client

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	gethcoretypes "github.com/ethereum/go-ethereum/core/types"
	enginetypes "github.com/itsdevbear/bolaris/engine/types"
	enginev1 "github.com/itsdevbear/bolaris/engine/types/v1"
	"github.com/itsdevbear/bolaris/types/consensus"
	"github.com/itsdevbear/bolaris/types/consensus/primitives"
)

// Caller defines a client that can interact with an Ethereum
// execution node's engine engineClient via JSON-RPC.
type Caller interface {
	// Generic Methods
	//
	// Start
	Start(context.Context)

	// Status returns the status of the execution client.
	Status() error

	// Engine API Related Methods
	//
	// NewPayload creates a new payload for the Ethereum execution node.
	NewPayload(ctx context.Context, payload enginetypes.ExecutionPayload,
		versionedHashes []common.Hash, parentBlockRoot *common.Hash) ([]byte, error)

	// ForkchoiceUpdated updates the fork choice of the Ethereum execution node.
	ForkchoiceUpdated(
		ctx context.Context, state *enginev1.ForkchoiceState,
		attrs enginetypes.PayloadAttributer,
	) (*enginev1.PayloadIDBytes, []byte, error)

	// GetPayload retrieves the payload from the Ethereum execution node.
	GetPayload(
		ctx context.Context, payloadID primitives.PayloadID, slot primitives.Slot,
	) (enginetypes.ExecutionPayload, *enginev1.BlobsBundle, bool, error)

	// ExecutionBlockByHash retrieves the execution block by its hash.
	ExecutionBlockByHash(ctx context.Context, hash common.Hash,
		withTxs bool) (*enginev1.ExecutionBlock, error)

	// Eth Namespace Methods
	//
	// BlockByHash retrieves the block by its hash.
	HeaderByHash(
		ctx context.Context,
		hash common.Hash,
	) (*gethcoretypes.Header, error)
}

// ExecutionPayloadRebuilder specifies a service capable of reassembling a
// complete beacon block, inclusive of an execution payload, using a signed
// beacon block and access to an execution client's engine API.
type PayloadReconstructor interface {
	ReconstructFullBlock(
		ctx context.Context, blindedBlock consensus.ReadOnlyBeaconKitBlock,
	) (consensus.BeaconKitBlock, error)
}