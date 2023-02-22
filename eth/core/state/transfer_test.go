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

package state_test

import (
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/state"
	vmmock "pkg.berachain.dev/stargazer/eth/core/vm/mock"
)

var _ = Describe("Transfers", func() {
	var sdb *vmmock.StargazerStateDBMock
	var addr common.Address

	BeforeEach(func() {
		sdb = vmmock.NewEmptyStateDB()
	})

	Context("Test CanTransfer", func() {
		It("should return true if the account has enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			ok := state.CanTransfer(sdb, addr, big.NewInt(100))
			Expect(ok).To(BeTrue())
		})

		It("should return false if the account does not have enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			ok := state.CanTransfer(sdb, addr, big.NewInt(101))
			Expect(ok).To(BeFalse())
		})
	})

	Context("Test Transfer", func() {
		It("should state.Transfer the amount if the account has enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			sdb.SubBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(0)
				}
			}
			sdb.AddBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(100)
				}
			}
			state.Transfer(sdb, addr, addr, big.NewInt(100))
			Expect(sdb.GetBalanceFunc(addr).Cmp(big.NewInt(100))).To(Equal(0))
		})

		It("should not state.Transfer the amount if the account does not have enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			sdb.SubBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(0)
				}
			}
			sdb.AddBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(100)
				}
			}
			state.Transfer(sdb, addr, addr, big.NewInt(101))
			Expect(sdb.GetBalanceFunc(addr).Cmp(big.NewInt(100))).To(Equal(0))
		})
	})
})
