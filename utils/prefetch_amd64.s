//go:build amd64

#include "textflag.h"

TEXT ·Prefetch(SB), NOSPLIT, $0
    MOVQ addr+0(FP), AX
    PREFETCHT0 (AX)
    RET
