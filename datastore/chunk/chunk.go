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

package chunk

import (
	"os"
	"time"

	"github.com/google/uuid"
)

type ChunkType int8

const (
	wal   ChunkType = 1
	btree ChunkType = 2
)

var (
	// maxCapacity set max capacity of each file to 512k
	maxCapacity = 512 * 1024
	folder      = os.TempDir()
)

type chunk struct {
	name        uuid.UUID
	chunkType   ChunkType
	createdTime int64
	sealedTime  int64

	capacity     int
	sealed       bool
	sealedLength int
}

func newChunk(chunkType ChunkType) (*chunk, error) {
	name := uuid.New()
	_, err := os.Create(chunkFileName(name, chunkType))
	if err != nil {
		return nil, err
	}
	// Todo: write metadata for the chunk
	c := &chunk{
		name:        name,
		createdTime: time.Now().UnixNano(),
		capacity:    maxCapacity,
		sealed:      false,
	}
	return c, nil
}

func chunkFileName(name uuid.UUID, chunkType ChunkType) string {
	return folder + string(os.PathSeparator) + name.String()
}
