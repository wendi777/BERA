// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for the full licensing terms.
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

package events_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/testutil"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state/events"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state/events/mock"
)

var _ = Describe("Manager", func() {
	var cem state.ControllableEventManager
	var ctx sdk.Context
	var ldb *mock.LogsDBMock

	BeforeEach(func() {
		ldb = mock.NewEmptyLogsDB()

		ctx = testutil.NewContext()
		ctx.EventManager().EmitEvent(sdk.NewEvent("1"))

		cem = events.NewManagerFrom(ctx.EventManager(), mock.NewPrecompileLogFactory())
		ctx = ctx.WithEventManager(cem)
		Expect(len(ctx.EventManager().Events())).To(Equal(1))
		Expect(cem.Events()).To(HaveLen(1))
	})

	It("should have the right registry key", func() {
		Expect(cem.RegistryKey()).To(Equal("events"))
	})

	It("should correctly snapshot/revert", func() {
		ctx.EventManager().EmitEvent(sdk.NewEvent("2"))
		Expect(len(ctx.EventManager().Events())).To(Equal(2))

		snap := cem.Snapshot()
		ctx.EventManager().EmitEvent(sdk.NewEvent("3"))
		Expect(len(ctx.EventManager().Events())).To(Equal(3))

		cem.RevertToSnapshot(snap)
		Expect(len(ctx.EventManager().Events())).To(Equal(2))
	})

	It("should not build eth logs when not in precompile", func() {
		ctx.EventManager().EmitEvent(sdk.NewEvent("2"))
		Expect(len(ctx.EventManager().Events())).To(Equal(2))
		Expect(len(ldb.AddLogCalls())).To(Equal(0))

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent("3"),
			sdk.NewEvent("4"),
		})
		Expect(len(ctx.EventManager().Events())).To(Equal(4))
		Expect(len(ldb.AddLogCalls())).To(Equal(0))
	})

	It("should panic when building eth logs fails", func() {
		cem.BeginPrecompileExecution(ldb)

		Expect(func() {
			ctx.EventManager().EmitEvent(sdk.NewEvent("non-eth-event"))
		}).To(Panic())
	})

	It("should build eth logs from cosmos events during precompile", func() {
		cem.BeginPrecompileExecution(ldb)

		ctx.EventManager().EmitEvent(sdk.NewEvent("2"))
		Expect(len(ctx.EventManager().Events())).To(Equal(2))
		Expect(len(ldb.AddLogCalls())).To(Equal(1))

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent("3"),
			sdk.NewEvent("4"),
		})
		Expect(len(ctx.EventManager().Events())).To(Equal(4))
		Expect(len(ldb.AddLogCalls())).To(Equal(3))

		cem.EndPrecompileExecution()

		Expect(func() { cem.Finalize() }).ToNot(Panic())
	})
})
