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

package avx2

import (
	"unsafe"

	"github.com/goshafaq/sonic/internal/native/types"
	"github.com/goshafaq/sonic/internal/rt"
)

var (
	__i64toa func(out unsafe.Pointer, val int64) (ret int)

	__u64toa func(out unsafe.Pointer, val uint64) (ret int)

	__f64toa func(out unsafe.Pointer, val float64) (ret int)

	__f32toa func(out unsafe.Pointer, val float32) (ret int)

	__lspace func(sp unsafe.Pointer, nb int, off int) (ret int)

	__quote func(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn unsafe.Pointer, flags uint64) (ret int)

	__html_escape func(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn unsafe.Pointer) (ret int)

	__unquote func(sp unsafe.Pointer, nb int, dp unsafe.Pointer, ep unsafe.Pointer, flags uint64) (ret int)

	__value func(s unsafe.Pointer, n int, p int, v unsafe.Pointer, flags uint64) (ret int)

	__vstring func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer, flags uint64)

	__vnumber func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer)

	__vsigned func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer)

	__vunsigned func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer)

	__skip_one func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer, flags uint64) (ret int)

	__skip_one_fast func(s unsafe.Pointer, p unsafe.Pointer) (ret int)

	__skip_array func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer, flags uint64) (ret int)

	__skip_object func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer, flags uint64) (ret int)

	__skip_number func(s unsafe.Pointer, p unsafe.Pointer) (ret int)

	__validate_one func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer) (ret int)

	__get_by_path func(s unsafe.Pointer, p unsafe.Pointer, path unsafe.Pointer, m unsafe.Pointer) (ret int)

	__validate_utf8 func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer) (ret int)

	__validate_utf8_fast func(s unsafe.Pointer) (ret int)

	__fsm_exec func(m unsafe.Pointer, s unsafe.Pointer, p unsafe.Pointer, flags uint64) (ret int)
)

//go:nosplit
func i64toa(out *byte, val int64) (ret int) {
	return __i64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func u64toa(out *byte, val uint64) (ret int) {
	return __u64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func f64toa(out *byte, val float64) (ret int) {
	return __f64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func f32toa(out *byte, val float32) (ret int) {
	return __f32toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func lspace(sp unsafe.Pointer, nb int, off int) (ret int) {
	return __lspace(rt.NoEscape(sp), nb, off)
}

//go:nosplit
func quote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) (ret int) {
	return __quote(rt.NoEscape(unsafe.Pointer(sp)), nb, rt.NoEscape(unsafe.Pointer(dp)), rt.NoEscape(unsafe.Pointer(dn)), flags)
}

//go:nosplit
func html_escape(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) (ret int) {
	return __html_escape(rt.NoEscape(unsafe.Pointer(sp)), nb, rt.NoEscape(unsafe.Pointer(dp)), rt.NoEscape(unsafe.Pointer(dn)))
}

//go:nosplit
func unquote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) (ret int) {
	return __unquote(rt.NoEscape(unsafe.Pointer(sp)), nb, rt.NoEscape(unsafe.Pointer(dp)), rt.NoEscape(unsafe.Pointer(ep)), flags)
}

//go:nosplit
func value(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) (ret int) {
	return __value(rt.NoEscape(unsafe.Pointer(s)), n, p, rt.NoEscape(unsafe.Pointer(v)), flags)
}

//go:nosplit
func vstring(s *string, p *int, v *types.JsonState, flags uint64) {
	__vstring(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)), flags)
}

//go:nosplit
func vnumber(s *string, p *int, v *types.JsonState) {
	__vnumber(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)))
}

//go:nosplit
func vsigned(s *string, p *int, v *types.JsonState) {
	__vsigned(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)))
}

//go:nosplit
func vunsigned(s *string, p *int, v *types.JsonState) {
	__vunsigned(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)))
}

//go:nosplit
func skip_one(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
	return __skip_one(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)), flags)
}

//go:nosplit
func skip_one_fast(s *string, p *int) (ret int) {
	return __skip_one_fast(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)))
}

//go:nosplit
func skip_array(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
	return __skip_array(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)), flags)
}

//go:nosplit
func skip_object(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
	return __skip_object(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)), flags)
}

//go:nosplit
func skip_number(s *string, p *int) (ret int) {
	return __skip_number(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)))
}

//go:nosplit
func validate_one(s *string, p *int, m *types.StateMachine) (ret int) {
	return __validate_one(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)))
}

//go:nosplit
func get_by_path(s *string, p *int, path *[]interface{}, m *types.StateMachine) (ret int) {
	return __get_by_path(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(path)), rt.NoEscape(unsafe.Pointer(m)))
}

//go:nosplit
func validate_utf8(s *string, p *int, m *types.StateMachine) (ret int) {
	return __validate_utf8(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)))
}

//go:nosplit
func validate_utf8_fast(s *string) (ret int) {
	return __validate_utf8_fast(rt.NoEscape(unsafe.Pointer(s)))
}

//go:nosplit
func fsm_exec(m *types.StateMachine, s *string, p *int, flags uint64) (ret int) {
	return __fsm_exec(rt.NoEscape(unsafe.Pointer(m)), rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), flags)
}
