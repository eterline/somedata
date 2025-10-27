// +build amd64,windows

#include "textflag.h"

TEXT ·prefetch(SB), NOSPLIT, $0
    MOVQ addr+0(FP), AX
    PREFETCHT0 (AX)
    RET