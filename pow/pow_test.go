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
)

func ExamplePoW_Compute() {
	pow := NewPoW([]byte("lynn9388"), 10)
	pow.Compute()

	hash := hash(pow.Data, pow.Nonce)
	fmt.Println(hex.EncodeToString(hash))
	// Output:
	// 002e7106c737ba4239194db04e477446201ef1dbe9386bdbd885981c762f951b
}

func ExamplePoW_Validate() {
	pow := NewPoW([]byte("lynn9388"), 10)
	if pow.Validate() == false {
		fmt.Println("New generated PoW is not valid")
	}

	pow.Compute()
	if pow.Validate() == true {
		fmt.Println("Computed PoW is valid")
	}
	// Output:
	// New generated PoW is not valid
	// Computed PoW is valid
}
