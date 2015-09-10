// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"internal/testenv"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TODO: move all these tests elsewhere?
// Perhaps teach test/run.go how to run them with a new action verb.
func runTest(t *testing.T, filename string) {
	doTest(t, filename, "run")
}
func buildTest(t *testing.T, filename string) {
	doTest(t, filename, "build")
}
func doTest(t *testing.T, filename string, kind string) {
	if runtime.GOARCH != "amd64" {
		t.Skipf("skipping SSA tests on %s for now", runtime.GOARCH)
	}
	testenv.MustHaveGoBuild(t)
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", kind, filepath.Join("testdata", filename))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	// TODO: set GOGC=off until we have stackmaps
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed: %v:\nOut: %s\nStderr: %s\n", err, &stdout, &stderr)
	}
	if s := stdout.String(); s != "" {
		t.Errorf("Stdout = %s\nWant empty", s)
	}
	if s := stderr.String(); strings.Contains(s, "SSA unimplemented") {
		t.Errorf("Unimplemented message found in stderr:\n%s", s)
	}
}

// TestShortCircuit tests OANDAND and OOROR expressions and short circuiting.
func TestShortCircuit(t *testing.T) { runTest(t, "short_ssa.go") }

// TestBreakContinue tests that continue and break statements do what they say.
func TestBreakContinue(t *testing.T) { runTest(t, "break_ssa.go") }

// TestArithmetic tests that both backends have the same result for arithmetic expressions.
func TestArithmetic(t *testing.T) { runTest(t, "arith_ssa.go") }

// TestFP tests that both backends have the same result for floating point expressions.
func TestFP(t *testing.T) { runTest(t, "fp_ssa.go") }

// TestArithmeticBoundary tests boundary results for arithmetic operations.
func TestArithmeticBoundary(t *testing.T) { runTest(t, "arithBoundary_ssa.go") }

// TestArithmeticConst tests results for arithmetic operations against constants.
func TestArithmeticConst(t *testing.T) { runTest(t, "arithConst_ssa.go") }

func TestChan(t *testing.T) { runTest(t, "chan_ssa.go") }

func TestCompound(t *testing.T) { runTest(t, "compound_ssa.go") }

func TestCtl(t *testing.T) { runTest(t, "ctl_ssa.go") }

func TestFp(t *testing.T) { runTest(t, "fp_ssa.go") }

func TestLoadStore(t *testing.T) { runTest(t, "loadstore_ssa.go") }

func TestMap(t *testing.T) { runTest(t, "map_ssa.go") }

func TestRegalloc(t *testing.T) { runTest(t, "regalloc_ssa.go") }

func TestString(t *testing.T) { runTest(t, "string_ssa.go") }

func TestDeferNoReturn(t *testing.T) { buildTest(t, "deferNoReturn_ssa.go") }

// TestClosure tests closure related behavior.
func TestClosure(t *testing.T) { runTest(t, "closure_ssa.go") }