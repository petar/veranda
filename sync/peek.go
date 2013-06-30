// Copyright 2013 Petar Maymounkov
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

package sync

import (
	"container/list"
	"io"
	"sync"
)

/*
	SEND  ––X4,X3,X2,X1-–>	RECV
	================================================
	Send (block)
							Receive	(never-block)
							Peek	(never-block)
							Len		(never-block)
	================================================
							Close	(never-block)
*/

// PeekChan is a chan-like object that supports peeking.
type PeekChan struct {
	scap   int
	slk    sync.Mutex
	scond  *sync.Cond
	sbuf   list.List
	closed bool
}

func MakePeekChan(limit int) *PeekChan {
	y := &PeekChan{}
	y.Init(limit)
	return y
}

func (y *PeekChan) Init(limit int) {
	y.scap = limit
	y.scond = sync.NewCond(&y.slk)
	y.sbuf.Init()
	y.closed = false
}

func (y *PeekChan) canSendOrClosed() bool {
	return y.closed || y.sbuf.Len() < y.scap
}

// Send returns io.ErrUnexpectedEOF if the PeekChan is closed.
func (y *PeekChan) Send(v interface{}) error {
	y.slk.Lock()
	defer y.slk.Unlock()
	for !y.canSendOrClosed() {
		y.scond.Wait()
	}
	if y.closed {
		return io.ErrUnexpectedEOF
	}
	y.sbuf.PushBack(v)
	return nil
}

// Receive returns the next message in the channel.
// Receive will panic if there are no messages in the channel.
func (y *PeekChan) Receive() (interface{}, bool) {
	y.slk.Lock()
	defer y.slk.Unlock()
	if y.closed {
		return nil, false
	}
	v := y.sbuf.Remove(y.sbuf.Front())
	y.scond.Broadcast()
	return v, true
}

// Peek returns a new channel with the current contents of this channel.
func (y *PeekChan) Peek() chan interface{} {
	y.slk.Lock()
	defer y.slk.Unlock()
	ch := make(chan interface{}, y.sbuf.Len())
	for e := y.sbuf.Front(); e != nil; e = e.Next() {
		ch <- e.Value
	}
	return ch
}

// Cap returns the capacity of the channel.
func (y *PeekChan) Cap() int {
	return y.scap
}

// Len returns the number of elements inside the channel
func (y *PeekChan) Len() int {
	y.slk.Lock()
	defer y.slk.Unlock()
	return y.sbuf.Len()
}

// Close closes the PeekChan, unblocking any waiting calls on Send which will return io.ErrUnexpectedEOF.
func (y *PeekChan) Close() {
	y.slk.Lock()
	defer y.slk.Unlock()
	y.closed = true
	y.scond.Broadcast()
}
