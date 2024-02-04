//go:build !byollvm && linux && !llvm14 && !llvm15 && !llvm16

package llvm

// #cgo CPPFLAGS: -I/usr/include/llvm-17 -I/usr/include/llvm-c-17 -I/usr/lib/llvm/17/include -I/usr/lib/llvm/17/include/llvm-c -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++17
// #cgo LDFLAGS: -L/usr/lib/llvm-17/lib -L/usr/lib/llvm/17/lib64  -lLLVM-17
import "C"

type run_build_sh int
