/*
 * Copyright © 2018 Lynn <lynn9388@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package pow implements the PoW algorithm.
package pow

import (
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/big"
)

// PoW represents a Poof-of-Work.
type PoW struct {
	Data       []byte // data for a proof
	TargetBits uint   // number of leading zero
	Nonce      int64  // random number to adjust the data's hash
}

// NewPoW returns a not computed PoW.
func NewPoW(data []byte, targetBits uint) *PoW {
	return &PoW{Data: data, TargetBits: targetBits, Nonce: 0}
}

// Compute performs a Proof-of-Work.
func (pow *PoW) Compute() {
	target := newTarget(pow.TargetBits)
	nonce := int64(0)
	maxNonce := int64(math.MaxInt64)

	for nonce <= maxNonce {
		var hashInt big.Int
		hashInt.SetBytes(hash(pow.Data, nonce))
		if hashInt.Cmp(target) == -1 {
			pow.Nonce = nonce
			break
		}
		nonce++
	}
}

// Validate checks if the PoW fulfils the requirement of Proof-of-Work.
func (pow *PoW) Validate() bool {
	var hashInt big.Int
	hashInt.SetBytes(hash(pow.Data, pow.Nonce))
	target := newTarget(pow.TargetBits)
	return hashInt.Cmp(target) == -1
}

// newTarget returns a target number with required leading zero.
func newTarget(targetBits uint) *big.Int {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return target
}

// hash returns the hash value of data|nonce.
func hash(data []byte, nonce int64) []byte {
	binaryNonce := make([]byte, 8)
	binary.BigEndian.PutUint64(binaryNonce, uint64(nonce))
	hash := sha256.Sum256(append(data, binaryNonce...))
	return hash[:]
}
