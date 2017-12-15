// Copyright © 2017 Ricardo Aravena <raravena@branch.io>
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

package scp

import (
	"io"
	"os"
	"testing"
	"time"
)

func TestNewProgressBar(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		size     int64
		outpipe  io.Writer
		expected ProgressBar
	}{
		{name: "Basic Progress Bar",
			size:    1000,
			subject: "testprogress",
			outpipe: os.Stdout,
			expected: ProgressBar{
				Out:       os.Stdout,
				Format:    DEFAULTFORMAT,
				Subject:   "testprogress",
				StartTime: time.Now(),
				Size:      1000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := NewProgressBar(tt.subject, tt.size)
			if returned.Subject != tt.expected.Subject {
				t.Errorf("Value received: %v expected %v", returned.Subject, tt.expected.Subject)
			}
			if returned.Size != tt.expected.Size {
				t.Errorf("Value received: %v expected %v", returned.Size, tt.expected.Size)
			}
			if returned.Out != tt.expected.Out {
				t.Errorf("Value received: %v expected %v", returned.Out, tt.expected.Out)
			}
			if returned.Format != tt.expected.Format {
				t.Errorf("Value received: %v expected %v", returned.Format, tt.expected.Format)
			}
		})
	}
}

func TestNewProgressBarTo(t *testing.T) {
	tests := []struct {
		name     string
		subject  string
		size     int64
		update   int64
		outpipe  io.Writer
		expected ProgressBar
	}{
		{name: "Basic To Progress Bar",
			size:    1000,
			subject: "testprogress",
			outpipe: os.Stdout,
			update:  10,
			expected: ProgressBar{
				Out:       os.Stdout,
				Format:    DEFAULTFORMAT,
				Subject:   "testprogress",
				StartTime: time.Now(),
				Size:      1000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := NewProgressBarTo(tt.subject, tt.size, tt.outpipe)
			if returned.Subject != tt.expected.Subject {
				t.Errorf("Value received: %v expected %v", returned.Subject, tt.expected.Subject)
			}
			if returned.Size != tt.expected.Size {
				t.Errorf("Value received: %v expected %v", returned.Size, tt.expected.Size)
			}
			if returned.Out != tt.expected.Out {
				t.Errorf("Value received: %v expected %v", returned.Out, tt.expected.Out)
			}
			if returned.Format != tt.expected.Format {
				t.Errorf("Value received: %v expected %v", returned.Format, tt.expected.Format)
			}
			// For coverage
			returned.Update(tt.update)
		})
	}
}
