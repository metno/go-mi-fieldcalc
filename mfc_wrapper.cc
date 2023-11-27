#include "mfc_wrapper.h"

#include "mi_fieldcalc/FieldCalculations.h"
#include <mi_fieldcalc/FieldDefined.h>

using miutil::ValuesDefined;

int mifieldcalc_abshum(int nx, int ny, const float *t, const float *rhum, float *abshumout,
           int *fDefined, float undef)
{
    ValuesDefined vDefined = (ValuesDefined)*fDefined;
    const bool ok = miutil::fieldcalc::abshum(nx, ny, t, rhum, abshumout, vDefined, undef);
    *fDefined = (int)vDefined;
    return ok;
}

int mifieldcalc_plevelhum(int nx, int ny, const float* t, const float* huminp, float p,
               int compute, float* humout, int* fDefined, float undef)
{
    ValuesDefined vDefined = (ValuesDefined)*fDefined;
    const bool ok = miutil::fieldcalc::plevelhum(nx, ny, t, huminp, p, "", compute, humout, vDefined, undef);
    *fDefined = (int)vDefined;
    return ok;
}

int mifieldcalc_alevelhum(int nx, int ny, const float* t, const float* huminp, const float* p,
               int compute, float* humout, int* fDefined, float undef)
{
    ValuesDefined vDefined = (ValuesDefined)*fDefined;
    const bool ok = miutil::fieldcalc::alevelhum(nx, ny, t, huminp, p, "", compute, humout, vDefined, undef);
    *fDefined = (int)vDefined;
    return ok;
}
