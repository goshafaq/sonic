// Code generated by Makefile, DO NOT EDIT.

// Code generated by Makefile, DO NOT EDIT.

/*
 * Copyright 2021 ByteDance Inc.
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

package avx

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastInt_IntToString(t *testing.T) {
	var buf [32]byte
	assert.Equal(t, "0", string(buf[:i64toa(&buf[0], 0)]))
	assert.Equal(t, "1", string(buf[:i64toa(&buf[0], 1)]))
	assert.Equal(t, "12", string(buf[:i64toa(&buf[0], 12)]))
	assert.Equal(t, "123", string(buf[:i64toa(&buf[0], 123)]))
	assert.Equal(t, "1234", string(buf[:i64toa(&buf[0], 1234)]))
	assert.Equal(t, "12345", string(buf[:i64toa(&buf[0], 12345)]))
	assert.Equal(t, "123456", string(buf[:i64toa(&buf[0], 123456)]))
	assert.Equal(t, "1234567", string(buf[:i64toa(&buf[0], 1234567)]))
	assert.Equal(t, "12345678", string(buf[:i64toa(&buf[0], 12345678)]))
	assert.Equal(t, "123456789", string(buf[:i64toa(&buf[0], 123456789)]))
	assert.Equal(t, "1234567890", string(buf[:i64toa(&buf[0], 1234567890)]))
	assert.Equal(t, "12345678901", string(buf[:i64toa(&buf[0], 12345678901)]))
	assert.Equal(t, "123456789012", string(buf[:i64toa(&buf[0], 123456789012)]))
	assert.Equal(t, "1234567890123", string(buf[:i64toa(&buf[0], 1234567890123)]))
	assert.Equal(t, "12345678901234", string(buf[:i64toa(&buf[0], 12345678901234)]))
	assert.Equal(t, "123456789012345", string(buf[:i64toa(&buf[0], 123456789012345)]))
	assert.Equal(t, "1234567890123456", string(buf[:i64toa(&buf[0], 1234567890123456)]))
	assert.Equal(t, "12345678901234567", string(buf[:i64toa(&buf[0], 12345678901234567)]))
	assert.Equal(t, "123456789012345678", string(buf[:i64toa(&buf[0], 123456789012345678)]))
	assert.Equal(t, "1234567890123456789", string(buf[:i64toa(&buf[0], 1234567890123456789)]))
	assert.Equal(t, "9223372036854775807", string(buf[:i64toa(&buf[0], 9223372036854775807)]))
	assert.Equal(t, "-1", string(buf[:i64toa(&buf[0], -1)]))
	assert.Equal(t, "-12", string(buf[:i64toa(&buf[0], -12)]))
	assert.Equal(t, "-123", string(buf[:i64toa(&buf[0], -123)]))
	assert.Equal(t, "-1234", string(buf[:i64toa(&buf[0], -1234)]))
	assert.Equal(t, "-12345", string(buf[:i64toa(&buf[0], -12345)]))
	assert.Equal(t, "-123456", string(buf[:i64toa(&buf[0], -123456)]))
	assert.Equal(t, "-1234567", string(buf[:i64toa(&buf[0], -1234567)]))
	assert.Equal(t, "-12345678", string(buf[:i64toa(&buf[0], -12345678)]))
	assert.Equal(t, "-123456789", string(buf[:i64toa(&buf[0], -123456789)]))
	assert.Equal(t, "-1234567890", string(buf[:i64toa(&buf[0], -1234567890)]))
	assert.Equal(t, "-12345678901", string(buf[:i64toa(&buf[0], -12345678901)]))
	assert.Equal(t, "-123456789012", string(buf[:i64toa(&buf[0], -123456789012)]))
	assert.Equal(t, "-1234567890123", string(buf[:i64toa(&buf[0], -1234567890123)]))
	assert.Equal(t, "-12345678901234", string(buf[:i64toa(&buf[0], -12345678901234)]))
	assert.Equal(t, "-123456789012345", string(buf[:i64toa(&buf[0], -123456789012345)]))
	assert.Equal(t, "-1234567890123456", string(buf[:i64toa(&buf[0], -1234567890123456)]))
	assert.Equal(t, "-12345678901234567", string(buf[:i64toa(&buf[0], -12345678901234567)]))
	assert.Equal(t, "-123456789012345678", string(buf[:i64toa(&buf[0], -123456789012345678)]))
	assert.Equal(t, "-1234567890123456789", string(buf[:i64toa(&buf[0], -1234567890123456789)]))
	assert.Equal(t, "-9223372036854775808", string(buf[:i64toa(&buf[0], -9223372036854775808)]))
}

func TestFastInt_UintToString(t *testing.T) {
	var buf [32]byte
	assert.Equal(t, "0", string(buf[:u64toa(&buf[0], 0)]))
	assert.Equal(t, "1", string(buf[:u64toa(&buf[0], 1)]))
	assert.Equal(t, "12", string(buf[:u64toa(&buf[0], 12)]))
	assert.Equal(t, "123", string(buf[:u64toa(&buf[0], 123)]))
	assert.Equal(t, "1234", string(buf[:u64toa(&buf[0], 1234)]))
	assert.Equal(t, "12345", string(buf[:u64toa(&buf[0], 12345)]))
	assert.Equal(t, "123456", string(buf[:u64toa(&buf[0], 123456)]))
	assert.Equal(t, "1234567", string(buf[:u64toa(&buf[0], 1234567)]))
	assert.Equal(t, "12345678", string(buf[:u64toa(&buf[0], 12345678)]))
	assert.Equal(t, "123456789", string(buf[:u64toa(&buf[0], 123456789)]))
	assert.Equal(t, "1234567890", string(buf[:u64toa(&buf[0], 1234567890)]))
	assert.Equal(t, "12345678901", string(buf[:u64toa(&buf[0], 12345678901)]))
	assert.Equal(t, "123456789012", string(buf[:u64toa(&buf[0], 123456789012)]))
	assert.Equal(t, "1234567890123", string(buf[:u64toa(&buf[0], 1234567890123)]))
	assert.Equal(t, "12345678901234", string(buf[:u64toa(&buf[0], 12345678901234)]))
	assert.Equal(t, "123456789012345", string(buf[:u64toa(&buf[0], 123456789012345)]))
	assert.Equal(t, "1234567890123456", string(buf[:u64toa(&buf[0], 1234567890123456)]))
	assert.Equal(t, "12345678901234567", string(buf[:u64toa(&buf[0], 12345678901234567)]))
	assert.Equal(t, "123456789012345678", string(buf[:u64toa(&buf[0], 123456789012345678)]))
	assert.Equal(t, "1234567890123456789", string(buf[:u64toa(&buf[0], 1234567890123456789)]))
	assert.Equal(t, "12345678901234567890", string(buf[:u64toa(&buf[0], 12345678901234567890)]))
	assert.Equal(t, "18446744073709551615", string(buf[:u64toa(&buf[0], 18446744073709551615)]))
}

func BenchmarkFastInt_IntToString(b *testing.B) {
	benchmarks := []struct {
		name string
		test func(*testing.B)
	}{{
		name: "StdLib-Positive",
		test: func(b *testing.B) {
			var buf [32]byte
			for i := 0; i < b.N; i++ {
				strconv.AppendInt(buf[:0], int64(i), 10)
			}
		},
	}, {
		name: "StdLib-Negative",
		test: func(b *testing.B) {
			var buf [32]byte
			for i := 0; i < b.N; i++ {
				strconv.AppendInt(buf[:0], -int64(i), 10)
			}
		},
	}, {
		name: "FastInt-Positive",
		test: func(b *testing.B) {
			var buf [32]byte
			for i := 0; i < b.N; i++ {
				i64toa(&buf[0], int64(i))
			}
		},
	}, {
		name: "FastInt-Negative",
		test: func(b *testing.B) {
			var buf [32]byte
			for i := 0; i < b.N; i++ {
				i64toa(&buf[0], -int64(i))
			}
		},
	}}
	for _, bm := range benchmarks {
		b.Run(bm.name, bm.test)
	}
}

type utoaBench struct {
	name string
	num  uint64
}

func BenchmarkFastInt_UintToString(b *testing.B) {
	maxUint := "18446744073709551615"
	benchs := make([]utoaBench, len(maxUint)+1)
	benchs[0].name = "Zero"
	benchs[0].num = 0
	for i := 1; i <= len(maxUint); i++ {
		benchs[i].name = strconv.FormatInt(int64(i), 10) + "-Digs"
		benchs[i].num, _ = strconv.ParseUint(string(maxUint[:i]), 10, 64)
	}

	for _, t := range benchs {
		benchmarks := []struct {
			name string
			test func(*testing.B)
		}{{
			name: "StdLib",
			test: func(b *testing.B) {
				var buf [32]byte
				for i := 0; i < b.N; i++ {
					strconv.AppendUint(buf[:0], t.num, 10)
				}
			},
		}, {
			name: "FastInt",
			test: func(b *testing.B) {
				var buf [32]byte
				for i := 0; i < b.N; i++ {
					u64toa(&buf[0], t.num)
				}
			},
		}}
		for _, bm := range benchmarks {
			name := fmt.Sprintf("%s_%s", bm.name, t.name)
			b.Run(name, bm.test)
		}
	}
}
