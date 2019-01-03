// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/riscv64"
)

func ginsnop() {
	// Hardware nop is ADD $0, ZERO
	p := gc.Prog(riscv64.AADD)
	p.From.Type = obj.TYPE_CONST
	p.From3 = &obj.Addr{Type: obj.TYPE_REG, Reg: riscv64.REG_ZERO}
	p.To = *p.From3
}
