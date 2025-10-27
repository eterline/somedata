//go:build arm64

#include "textflag.h"

TEXT ·prefetch(SB), NOSPLIT, $0
    MOVD addr+0(FP), R0
    PRFM PLDL1KEEP, (R0)
    RET