#ifndef mfc_wrapper_h
#define mfc_wrapper_h

#ifdef __cplusplus
extern "C" {
#endif

int mifieldcalc_abshum(int nx, int ny, const float* t, const float* rhum, float* abshumout, int* fDefined, float undef);

int mifieldcalc_plevelhum(int nx, int ny, const float* t, const float* huminp, float p,
               int compute, float* humout, int* fDefined, float undef);

int mifieldcalc_alevelhum(int nx, int ny, const float* t, const float* huminp, const float* p,
               int compute, float* humout, int* fDefined, float undef);

#ifdef __cplusplus
}
#endif

#endif // mfc_wrapper_h
