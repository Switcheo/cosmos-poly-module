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

package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/polynetwork/cosmos-poly-module/lockproxypip1/internal/types"
)

const (
	// default paramspace for params keeper
	DefaultParamspace = types.ModuleName
)

var (
	OperatorToLockProxyKey = []byte{0x01}
	BindChainIdPrefix      = []byte{0x02}
	RegistryPrefix         = []byte{0x05}
)

func GetOperatorToLockProxyKey(operator sdk.AccAddress) []byte {
	return append(OperatorToLockProxyKey, operator...)
}

func GetRegistryKey(lockProxyHash []byte, assetHash []byte, nativeChainId uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) []byte {
	nativeChainIdBz := make([]byte, 8)
	binary.LittleEndian.PutUint64(nativeChainIdBz, nativeChainId)
	return append(append(append(append(append(RegistryPrefix, lockProxyHash...), assetHash...), nativeChainIdBz...), nativeLockProxyHash...), nativeAssetHash...)
}

func GetBindChainIdKey(proxyHash []byte, toChainId uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, toChainId)
	return append(append(BindChainIdPrefix, proxyHash...), b...)
}
