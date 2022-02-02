package v2

import (
	"fmt"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

// ToProtoUnsafe returns the pointer value of the underlying
// beacon state proto object, bypassing immutability. Use with care.
func (b *BeaconState) ToProtoUnsafe() interface{} {
	if b == nil {
		return nil
	}

	gvrCopy := b.genesisValidatorsRoot

	return &ethpb.BeaconStateAltair{
		GenesisTime:                 b.genesisTime,
		GenesisValidatorsRoot:       gvrCopy[:],
		Slot:                        b.slot,
		Fork:                        b.fork,
		LatestBlockHeader:           b.latestBlockHeader,
		BlockRoots:                  b.blockRoots.Slice(),
		StateRoots:                  b.stateRoots.Slice(),
		HistoricalRoots:             b.historicalRoots.Slice(),
		Eth1Data:                    b.eth1Data,
		Eth1DataVotes:               b.eth1DataVotes,
		Eth1DepositIndex:            b.eth1DepositIndex,
		Validators:                  b.validators,
		Balances:                    b.balances,
		RandaoMixes:                 b.randaoMixes.Slice(),
		Slashings:                   b.slashings,
		PreviousEpochParticipation:  b.previousEpochParticipation,
		CurrentEpochParticipation:   b.currentEpochParticipation,
		JustificationBits:           b.justificationBits,
		PreviousJustifiedCheckpoint: b.previousJustifiedCheckpoint,
		CurrentJustifiedCheckpoint:  b.currentJustifiedCheckpoint,
		FinalizedCheckpoint:         b.finalizedCheckpoint,
		InactivityScores:            b.inactivityScores,
		CurrentSyncCommittee:        b.currentSyncCommittee,
		NextSyncCommittee:           b.nextSyncCommittee,
	}
}

// ToProto the beacon state into a protobuf for usage.
func (b *BeaconState) ToProto() interface{} {
	if b == nil {
		return nil
	}

	b.lock.RLock()
	defer b.lock.RUnlock()

	gvrCopy := b.genesisValidatorsRoot

	return &ethpb.BeaconStateAltair{
		GenesisTime:                 b.genesisTime,
		GenesisValidatorsRoot:       gvrCopy[:],
		Slot:                        b.slot,
		Fork:                        b.forkVal(),
		LatestBlockHeader:           b.latestBlockHeaderVal(),
		BlockRoots:                  b.blockRoots.Slice(),
		StateRoots:                  b.stateRoots.Slice(),
		HistoricalRoots:             b.historicalRoots.Slice(),
		Eth1Data:                    b.eth1DataVal(),
		Eth1DataVotes:               b.eth1DataVotesVal(),
		Eth1DepositIndex:            b.eth1DepositIndex,
		Validators:                  b.validatorsVal(),
		Balances:                    b.balancesVal(),
		RandaoMixes:                 b.randaoMixes.Slice(),
		Slashings:                   b.slashingsVal(),
		PreviousEpochParticipation:  b.previousEpochParticipationVal(),
		CurrentEpochParticipation:   b.currentEpochParticipationVal(),
		JustificationBits:           b.justificationBitsVal(),
		PreviousJustifiedCheckpoint: b.previousJustifiedCheckpointVal(),
		CurrentJustifiedCheckpoint:  b.currentJustifiedCheckpointVal(),
		FinalizedCheckpoint:         b.finalizedCheckpointVal(),
		InactivityScores:            b.inactivityScoresVal(),
		CurrentSyncCommittee:        b.currentSyncCommitteeVal(),
		NextSyncCommittee:           b.nextSyncCommitteeVal(),
	}
}

// StateRoots kept track of in the beacon state.
func (b *BeaconState) StateRoots() [][]byte {
	if b.stateRoots == nil {
		return nil
	}

	b.lock.RLock()
	defer b.lock.RUnlock()

	return b.stateRoots.Slice()
}

// StateRootAtIndex retrieves a specific state root based on an
// input index value.
func (b *BeaconState) StateRootAtIndex(idx uint64) ([]byte, error) {
	if b.stateRoots == nil {
		return nil, nil
	}

	b.lock.RLock()
	defer b.lock.RUnlock()

	r, err := b.stateRootAtIndex(idx)
	if err != nil {
		return nil, err
	}
	return r[:], nil
}

// stateRootAtIndex retrieves a specific state root based on an
// input index value.
// This assumes that a lock is already held on BeaconState.
func (b *BeaconState) stateRootAtIndex(idx uint64) ([32]byte, error) {
	if uint64(len(b.stateRoots)) <= idx {
		return [32]byte{}, fmt.Errorf("index %d out of range", idx)
	}

	return b.stateRoots[idx], nil
}

// ProtobufBeaconState transforms an input into beacon state hard fork 1 in the form of protobuf.
// Error is returned if the input is not type protobuf beacon state.
func ProtobufBeaconState(s interface{}) (*ethpb.BeaconStateAltair, error) {
	pbState, ok := s.(*ethpb.BeaconStateAltair)
	if !ok {
		return nil, errors.New("input is not type pb.BeaconStateAltair")
	}
	return pbState, nil
}

// InnerStateUnsafe returns the pointer value of the underlying
// beacon state proto object, bypassing immutability. Use with care.
func (b *BeaconState) InnerStateUnsafe() interface{} {
	return b.ToProtoUnsafe()
}

// CloneInnerState the beacon state into a protobuf for usage.
func (b *BeaconState) CloneInnerState() interface{} {
	return b.ToProto()
}