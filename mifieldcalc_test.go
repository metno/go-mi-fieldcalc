package mifieldcalc

import (
	"math"
	"testing"
)

const (
	UNDEF = (float32)(12356789.0)
	T0    = 273.15
)

func near(a float32, b float32, tolerance float32) bool {
	return math.Abs(float64(a-b)) < float64(tolerance)
}

func TestAbsHum(test *testing.T) {

	f_abshum := []float32{(float32)(2 * UNDEF)}
	f_airtmp := []float32{293.16}
	f_relhum := []float32{0.8}
	fDefined := SOME_DEFINED

	err := Abshum(1, 1, f_airtmp, f_relhum, f_abshum, &fDefined, UNDEF)
	if err != nil {
		test.Fatalf("error in abshum: %v\n", err)
	}

	exp := float32(13.8400)
	if !near(f_abshum[0], exp, 1e-4) {
		test.Errorf("abshum actual %f expected %f\n", f_abshum[0], exp)
	}
	if fDefined != ALL_DEFINED {
		test.Errorf("abshum expeted all defined got %d\n", fDefined)
	}
}

func TestXLevelHum(test *testing.T) {

	type levelhum_params_t struct {
		ca     int     // 'compute' for alevelhum
		cp     int     // 'compute' for plevelhum
		t      float32 // temperature input
		humin  float32 // humidity input
		p      float32 // pressure input
		expect float32 // expected output
		near   float32 // max deviation from expected
	}

	levelhum_params := []levelhum_params_t{
		// alevelhum/hlevelhum and plevelhum have compute numbers >= 5 switched
		{1, 1, 30.68 + T0, .025, 1013, 91.9, 0.1},
		{2, 2, 302.71, .025, 1013, 91.9, 0.1},
		{3, 3, 30.68 + T0, 55, 1013, 0.014963, .000001},
		{4, 4, 302.71, 55, 1013, 0.014963, .000001},
		{5, 7, 30.68 + T0, .015, 1013, 20.6, 0.1},
		{6, 8, 302.71, .015, 1013, 20.6, 0.1},
		{7, 5, 30.68 + T0, 55, 1013, 20.6, 0.1},
		{8, 6, 302.71, 55, 1013, 20.6, 0.1},
	}

	for i := range levelhum_params {
		tc := &levelhum_params[i]

		tc_t := []float32{tc.t}
		tc_humin := []float32{tc.humin}
		tc_p := []float32{tc.p}

		// test alevelhum

		out := []float32{(float32)(2 * UNDEF)}
		fDefined := SOME_DEFINED
		err := Alevelhum(1, 1, tc_t, tc_humin, tc_p, tc.ca, out, &fDefined, UNDEF)
		if err != nil {
			test.Fatalf("error in alevelhum: %v\n", err)
		}

		if !near(out[0], tc.expect, tc.near) {
			test.Errorf("alevelhum actual %f expected %f\n", out[0], tc.expect)
		}
		if fDefined != ALL_DEFINED {
			test.Errorf("alevelhum expected all defined got %d\n", fDefined)
		}

		// test plevelhum

		out[0] = 2 * UNDEF
		fDefined = SOME_DEFINED
		err = Plevelhum(1, 1, tc_t, tc_humin, tc.p, tc.cp, out, &fDefined, UNDEF)
		if err != nil {
			test.Fatalf("error in plevelhum: %v\n", err)
		}

		if !near(out[0], tc.expect, tc.near) {
			test.Errorf("plevelhum actual %f expected %f\n", out[0], tc.expect)
		}
		if fDefined != ALL_DEFINED {
			test.Errorf("plevelhum expected all defined got %d\n", fDefined)
		}
	}
}
