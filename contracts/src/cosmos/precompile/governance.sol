// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
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

pragma solidity ^0.8.4;

interface IGovernanceModule {
    ////////////////////////////////////////// Write Methods /////////////////////////////////////////////
    /**
     * @dev Submit a proposal to the governance module. Returns the proposal id.
     * @param proposal The proposal to submit.
     * @param message The message to submit with the proposal.
     */
    function submitProposal(
        bytes calldata proposal,
        bytes calldata message
    ) external returns (uint64);

    /**
     * @dev Cancel a proposal. Returns the cancled time and height.
     *   burned.
     * @param proposalId The id of the proposal to cancel.
     */
    function cancelProposal(
        uint64 proposalId
    ) external returns (uint64, uint64);

    /**
     * @dev Vote on a proposal.
     * @param proposalId The id of the proposal to vote on.
     * @param option The option to vote on.
     * @param metadata The metadata to attach to the vote.
     */
    function vote(
        uint64 proposalId,
        int32 option,
        string memory metadata
    ) external returns (bool);

    /**
     * @dev Vote on a proposal with weights.
     * @param proposalId The id of the proposal to vote on.
     * @param options The options to vote on.
     * @param metadata The metadata to attach to the vote.
     */
    function voteWeighted(
        uint64 proposalId,
        WeightedVoteOption[] calldata options,
        string calldata metadata
    ) external returns (bool);

    ////////////////////////////////////////// Read Methods /////////////////////////////////////////////

    /**
     * @dev Get the proposal with the given id.
     */
    function getProposal(
        uint64 proposalId
    ) external view returns (Proposal memory);

    /**
     * @dev Get proposals with a given status.
     * @param proposalStatus The status of the proposals to get.
     */
    function getProposals(
        int32 proposalStatus
    ) external view returns (Proposal[] memory);

    ////////////////////////////////////////// Structs ///////////////////////////////////////////////////

    /**
     * @dev Represents a cosmos coin.
     * Note: this struct is generated as go struct that is then used in the precompile.
     */
    struct Coin {
        uint64 amount;
        string denom;
    }

    /**
     * @dev Represents a governance module `WeightedVoteOption`.
     * Note: this struct is generated in generated/i_staking_module.abigen.go
     */
    struct WeightedVoteOption {
        int32 voteOption;
        string weight;
    }

    /**
     * @dev Represents a governance module `Proposal`.
     * Note: this struct is generated in generated/i_staking_module.abigen.go
     */
    struct Proposal {
        uint64 id;
        bytes message;
        int32 status;
        TallyResult finalTallyResult;
        uint64 submitTime;
        uint64 depositEndTime;
        Coin[] totalDeposit;
        uint64 votingStartTime;
        uint64 votingEndTime;
        string metadata;
        string title;
        string summary;
        string proposer;
    }

    /**
     * @dev Represents a governance module `TallyResult`.
     * Note: this struct is generated in generated/i_staking_module.abigen.go
     */
    struct TallyResult {
        string yesCount;
        string abstainCount;
        string noCount;
        string noWithVetoCount;
    }

    // /**
    //  * @dev Emitted by the governance module when `submitProposal` is called.
    //  * @param proposalId The id of the proposal.
    //  * @param proposalMessages The messages of the proposal.
    //  */
    // event SubmitProposal(
    //     uint64 indexed proposalId,
    //     string indexed proposalMessages
    // );

    // /**
    //  * @dev Emitted by the governance module when `submitProposal` is called.
    //  * @param votingPeriodStart The start of the voting period.
    //  */
    // event SubmitProposal(uint64 indexed votingPeriodStart);

    // /**
    //  * @dev Emitted by the governance module when `submitProposal` is called.
    //  * @param proposalId The id of the proposal.
    //  * @param votingPeriodStart is the time stamp of when voting started.
    //  */
    // event ProposalDeposit(uint64 indexed proposalId, uint64 votingPeriodStart);

    // /**
    //  * @dev Emitted by the governance module when `AddVote` is called in the msg server.
    //  * @param proposalId The id of the proposal.
    //  */
    // event ProposalVote(uint64 indexed proposalId);

    // /**
    //  * @dev Emitted by the governance module when `cancelProposal` is called.
    //  * @param sender The sender of the cancel proposal.
    //  * @param proposalId The id of the proposal.
    //  */
    // event CancelProposal(address indexed sender, uint64 indexed proposalId);
}
