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
	"strings"
	"testing"
)

func TestGetNonce(t *testing.T) {
	data := []byte("lynn9388")
	for difficulty := 1; difficulty < 6; difficulty++ {
		nonce := GetNonce(data, uint(difficulty))
		t.Log(Hash(data, nonce))
		if !strings.HasPrefix(Hash(data, nonce), strings.Repeat("0", difficulty)) {
			t.Fail()
		}
	}
}
