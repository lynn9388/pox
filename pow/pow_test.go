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

package pow

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestPoW_Compute(t *testing.T) {
	for i := 0; i < 20; i++ {
		pow := NewPoW([]byte("lynn9388"), uint(i))
		pow.Compute()

		hash := hash(pow.Data, pow.Nonce)
		fmt.Println(hex.EncodeToString(hash))
	}
}

func TestPoW_Validate(t *testing.T) {
	for i := 10; i < 20; i++ {
		pow := NewPoW([]byte("lynn9388"), uint(i))
		if pow.Validate() == true {
			t.Error("failed to validate new generated PoW")
		}

		pow.Compute()
		if pow.Validate() == false {
			t.Error("failed to validate computed PoW")
		}
	}
}
