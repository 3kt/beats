// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build linux

package journalctl

import "fmt"

type SeekMode uint

const (
	// SeekHead option seeks to the head of a journal
	SeekHead SeekMode = iota
	// SeekTail option seeks to the tail of a journal
	SeekTail
	// SeekCursor option seeks to the position specified in the cursor
	// SeekSince option seeks to the position specified by the since option
	SeekSince
)

var seekModes = map[string]SeekMode{
	"head":  SeekHead,
	"tail":  SeekTail,
	"since": SeekSince,
}

// Unpack validates and unpack "seek" config options. It returns an error if
// the string is no valid seek mode.
func (m *SeekMode) Unpack(value string) error {
	mode, ok := seekModes[value]
	if !ok {
		return fmt.Errorf("invalid seek mode '%s'", value)
	}

	*m = mode
	return nil
}
