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

package suite

import (
	"time"

	"github.com/ethereum/go-ethereum/params"
)

// Ether represents the number of wei in one ether, used for Ethereum
// transactions.
const (
	Ether   = params.Ether
	OneGwei = uint64(params.GWei) // 1 Gwei = 1e9 wei
	TenGwei = 10 * OneGwei        // 10 Gwei = 1e10 wei
)

// EtherTransferGasLimit specifies the gas limit for a standard Ethereum
// transfer.
// This is the amount of gas required to perform a basic ether transfer.
const (
	EtherTransferGasLimit uint64 = 21000 // Standard gas limit for ether transfer
)

// DefaultE2ETestTimeout defines the default timeout duration for end-to-end
// tests. This is used to specify how long to wait for a test before considering
// it failed.
const (
	DefaultE2ETestTimeout = 60 * 5 * time.Second // timeout for E2E tests
)