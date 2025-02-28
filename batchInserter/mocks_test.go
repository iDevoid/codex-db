/**
 * Copyright 2019 Comcast Cable Communications Management, LLC
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
 *
 */

package batchInserter

import (
	"github.com/stretchr/testify/mock"
	"github.com/xmidt-org/codex-db"
)

type mockInserter struct {
	mock.Mock
}

func (c *mockInserter) InsertRecords(records ...db.Record) error {
	args := c.Called(records)
	return args.Error(0)
}
