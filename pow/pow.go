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
	"encoding/hex"
	"strings"
)

// GetNonce returns a nonce fulfilled the requirement of PoW.
func GetNonce(data []byte, difficulty uint) []byte {
	nonce := make([]byte, 4)
	for i := uint32(1); !Fulfill(data, nonce, difficulty); i++ {
		binary.LittleEndian.PutUint32(nonce, i)
	}
	return nonce
}

// Fulfill checks if the nonce fulfils the requirement of PoW.
func Fulfill(data []byte, nonce []byte, difficulty uint) bool {
	hash := Hash(data, nonce)
	prefix := strings.Repeat("0", int(difficulty))
	return strings.HasPrefix(hash, prefix)
}

// Hash returns the SHA256 Hash values in hexadecimal.
func Hash(data []byte, nonce []byte) string {
	hash := sha256.Sum256(append(data, nonce...))
	return hex.EncodeToString(hash[:])
}
