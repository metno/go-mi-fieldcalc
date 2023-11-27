package mifieldcalc

/*
#cgo LDFLAGS: -lmi-fieldcalc

#include <mfc_wrapper.h>
*/
import "C"

import (
	"errors"
)

const (
	ALL_DEFINED  = 0
	NONE_DEFINED = 1
	SOME_DEFINED = 2
)

func Abshum(nx int, ny int, f_airtmp []float32, f_relhum []float32, f_abshum []float32, fDefined *int, undef float32) error {
	size := nx * ny
	if len(f_airtmp) != size || len(f_relhum) != size || len(f_abshum) != size {
		return errors.New("slice size does not match nx*ny")
	}

	fdef := (C.int)(*fDefined)
	if ok := C.mifieldcalc_abshum((C.int)(nx), (C.int)(ny), (*C.float)(&f_airtmp[0]), (*C.float)(&f_relhum[0]), (*C.float)(&f_abshum[0]), &fdef, (C.float)(undef)); ok == 0 {
		return errors.New("abshum failed")
	}
	*fDefined = (int)(fdef)
	return nil
}

// Pressure level humidity calculations:
//
//	compute=1  : temp. (Kelvin) og spes. fukt. -> rel. fuktighet (%)
//	compute=2  : pot. temp. og spesifikk fukt. -> rel. fuktighet (%)
//	compute=3  : temp. (Kelvin) og rel. fukt.  -> spesifikk fukt. (kg/kg)
//	compute=4  : pot. temp. og  relativ fukt.  -> spesifikk fukt. (kg/kg)
//	compute=5  : temp. (Kelvin) og rel. fukt.  -> duggpunkt, Td (Celsius)
//	compute=6  : pot. temp. og  relativ fukt.  -> duggpunkt, Td (Celsius)
//	compute=7  : temp. (Kelvin) og spes. fukt. -> duggpunkt, Td (Celsius)
//	compute=8  : pot. temp. og spesifikk fukt. -> duggpunkt, Td (Celsius)
//	compute=9  : temp. (Kelvin) og rel. fukt.  -> duggpunkt, Td (Kelvin)
//	compute=10 : pot. temp. og  relativ fukt.  -> duggpunkt, Td (Kelvin)
//	compute=11 : temp. (Kelvin) og spes. fukt. -> duggpunkt, Td (Kelvin)
//	compute=12 : pot. temp. og spesifikk fukt. -> duggpunkt, Td (Kelvin)
func Plevelhum(nx int, ny int, t []float32, huminp []float32, p float32, compute int, humout []float32, fDefined *int, undef float32) error {
	size := nx * ny
	if len(t) != size || len(huminp) != size || len(humout) != size {
		return errors.New("slice size does not match nx*ny")
	}

	fdef := (C.int)(*fDefined)
	if ok := C.mifieldcalc_plevelhum((C.int)(nx), (C.int)(ny), (*C.float)(&t[0]), (*C.float)(&huminp[0]), (C.float)(p), (C.int)(compute), (*C.float)(&humout[0]), &fdef, (C.float)(undef)); ok == 0 {
		return errors.New("plevelhum failed")
	}
	*fDefined = (int)(fdef)
	return nil
}

// Ukjente modell-flater, gitt trykk (p):
//
//	compute=1  : temp. (Kelvin) og spes. fukt. -> rel. fuktighet (%)
//	compute=2  : pot. temp. og spesifikk fukt. -> rel. fuktighet (%)
//	compute=3  : temp. (Kelvin) og rel. fukt.  -> spesifikk fukt. (kg/kg)
//	compute=4  : pot. temp. og  relativ fukt.  -> spesifikk fukt. (kg/kg)
//	compute=5  : temp. (Kelvin) og spes. fukt. -> duggpunkt, Td (Celsius)
//	compute=6  : pot. temp. og spesifikk fukt. -> duggpunkt, Td (Celsius)
//	compute=7  : temp. (Kelvin) og rel. fukt.  -> duggpunkt, Td (Celsius)
//	compute=8  : pot. temp. og  relativ fukt.  -> duggpunkt, Td (Celsius)
//	compute=9  : temp. (Kelvin) og spes. fukt. -> duggpunkt, Td (Kelvin)
//	compute=10 : pot. temp. og spesifikk fukt. -> duggpunkt, Td (Kelvin)
//	compute=11 : temp. (Kelvin) og rel. fukt.  -> duggpunkt, Td (Kelvin)
//	compute=12 : pot. temp. og  relativ fukt.  -> duggpunkt, Td (Kelvin)
func Alevelhum(nx int, ny int, t []float32, huminp []float32, p []float32, compute int, humout []float32, fDefined *int, undef float32) error {
	size := nx * ny
	if len(t) != size || len(huminp) != size || len(humout) != size || len(p) != size {
		return errors.New("slice size does not match nx*ny")
	}

	fdef := (C.int)(*fDefined)
	if ok := C.mifieldcalc_alevelhum((C.int)(nx), (C.int)(ny), (*C.float)(&t[0]), (*C.float)(&huminp[0]), (*C.float)(&p[0]), (C.int)(compute), (*C.float)(&humout[0]), &fdef, (C.float)(undef)); ok == 0 {
		return errors.New("alevelhum failed")
	}
	*fDefined = (int)(fdef)
	return nil
}
