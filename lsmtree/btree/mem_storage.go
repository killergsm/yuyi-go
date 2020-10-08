// Copyright 2015 The yuyi Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package btree

import "github.com/satori/go.uuid"

// MaxLength set max capacity of each file to 4M
const MaxLength = 4 * 1024 * 1024

var (
	file = uuid.NewV4().String()
	off  = 0

	cache map[Address][]byte
)

func WriteTo(input []byte) Address {
	size := len(input)
	if off+size > MaxLength {
		// need rotate to another file
		file = uuid.NewV4().String()
		off = 0
	}
	// create new address and cache it.
	res := Address{File: file, Offset: off, Length: size}
	cache[res] = input
	off += size
	return res
}

func ReadFrom(address Address) []byte {
	return cache[address]
}
