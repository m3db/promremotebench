// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/uber-go/tally"
	"go.uber.org/zap/zaptest"
)

const testResult = `{
	"status" : "success",
	"data" : {
	   "resultType" : "matrix",
	   "result" : [
		  {
			 "metric" : {
				"__name__" : "up",
				"job" : "prometheus",
				"instance" : "localhost:9090"
			 },
			 "values" : [
				[ 1435781430.781, "5" ],
				[ 1435781445.781, "8" ],
				[ 1435781460.781, "1" ]
			 ]
		  }
	   ]
	}
 }`

const testResult2 = `
 {
    "status": "success",
    "data": {
        "resultType": "matrix",
        "result": [
            {
                "metric": {},
                "values": [
                    [
                        1567192552,
                        "2298081297676.0757"
                    ],
                    [
                        1567192562,
                        "2298619694338.755"
                    ],
                    [
                        1567192572,
                        "2299091718680.2"
                    ],
                    [
                        1567192582,
						"2299531757764.756"
                    ]
                ],
                "step_size_ms": 10000
            }
        ]
    }
}`

func TestValidateQuery(t *testing.T) {
	query := newQueryExecutor(queryExecutorOptions{
		Scope:  tally.NoopScope,
		Logger: zaptest.NewLogger(t),
	})
	data := []byte(testResult)
	require.True(t, query.validateQuery(
		Datapoints{
			Datapoint{
				Timestamp: promTimestampToTime(1435781430781),
				Value:     5,
			},
			Datapoint{
				Timestamp: promTimestampToTime(1435781445781),
				Value:     8,
			},
			Datapoint{
				Timestamp: promTimestampToTime(1435781460781),
				Value:     1,
			},
		},
		data,
		"foo",
	))

	data = []byte(testResult2)
	require.True(t, query.validateQuery(
		Datapoints{
			Datapoint{
				Timestamp: time.Now(),
				Value:     2297275298167.599,
			},
			Datapoint{
				Timestamp: time.Now(),
				Value:     2297369112899.5605,
			},
			Datapoint{
				Timestamp: time.Now(),
				Value:     2297325109692.599,
			},
			Datapoint{
				Timestamp: time.Now(),
				Value:     2298081297676.075,
			},
			Datapoint{
				Timestamp: time.Now(),
				Value:     2298619694338.755,
			},
			Datapoint{
				Timestamp: time.Now(),
				Value:     2299091718680.2,
			},
			Datapoint{
				Timestamp: time.Now(),
				Value:     2299531757764.756,
			},
		},
		data,
		"foo",
	))

	require.False(t, query.validateQuery(
		Datapoints{
			Datapoint{
				Timestamp: time.Now(),
				Value:     1,
			},
			Datapoint{
				Timestamp: time.Now(),
				Value:     2,
			},
		},
		data,
		"foo",
	))
}
