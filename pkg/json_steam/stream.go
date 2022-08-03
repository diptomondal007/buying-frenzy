// Licensed to Dipto Mondal under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Dipto Mondal licenses this file to you under
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

package json_steam

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Stream holds the things needed for a json stream
type Stream struct {
	c     chan Entry
	value chan interface{}
	want  chan struct{}
}

// Entry is data holder for each entry in data file
type Entry struct {
	Data  interface{}
	Error error
}

// NewJSONStreamer returns a new stream instance
func NewJSONStreamer() Stream {
	return Stream{
		c:     make(chan Entry),
		value: make(chan interface{}),
		want:  make(chan struct{}),
	}
}

// Watch ...
func (s Stream) Watch() <-chan Entry {
	return s.c
}

// Want ...
func (s Stream) Want() <-chan struct{} {
	return s.want
}

// Value ...
func (s Stream) Value() chan<- interface{} {
	return s.value
}

// Start ...
func (s Stream) Start(file string) {
	defer close(s.c)

	f, err := os.Open(file)
	if err != nil {
		s.c <- Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}
	defer f.Close()

	de := json.NewDecoder(f)
	if _, err = de.Token(); err != nil {
		s.c <- Entry{Error: fmt.Errorf("error decoding json: %w", err)}
		return
	}

	i := 1
	for de.More() {
		s.want <- struct{}{}
		v := <-s.value

		if err := de.Decode(&v); err != nil {
			s.c <- Entry{Error: fmt.Errorf("decode line %d: %w", i, err)}
			return
		}
		s.c <- Entry{
			Data:  v,
			Error: nil,
		}
		i++
	}

	// Read closing delimiter. `]` or `}`
	if _, err := de.Token(); err != nil {
		s.c <- Entry{Error: fmt.Errorf("decode closing delimiter: %w", err)}
		return
	}

	log.Println(">>>>>>>>>>>>> reading data successful")
}
