/*
Copyright 2021 Erigon contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package bitutil

import (
	"math"
	"testing"
)

func TestSelect64(t *testing.T) {
	if res := Select64(5270498307387724361, 14); res != 41 {
		panic(res)
	}
	if res := Select64(5270498307387724361, 6); res != 18 {
		panic(res)
	}
	if res := Select64(uint64(math.MaxUint64), 62); res != 62 {
		panic(res)
	}
	if res := Select64(210498307387724361, 14); res != 35 {
		panic(res)
	}
}
