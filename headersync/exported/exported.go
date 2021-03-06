/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/polynetwork/cosmos-poly-module/headersync/internal/types"
	polytype "github.com/polynetwork/poly/core/types"
)

// DelegationI delegation bond for a delegated proof of stake system
type HeaderSyncKeeper interface {
	ProcessHeader(ctx sdk.Context, header *polytype.Header, headerProof []byte, curHeader *polytype.Header) error
	GetConsensusPeers(ctx sdk.Context, chainId uint64) (*types.ConsensusPeers, error)
}
