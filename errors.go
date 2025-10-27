package somedata

type sErr string

func (e sErr) Error() string { return string(e) }

// Ring buffer errors
const (
	ErrRingBuffAboveZero sErr = "ring buffer: size must be above 0"
	ErrBUfferOverflow    sErr = "ring buffer: overflow"
)

// Matrix errors
const ()
