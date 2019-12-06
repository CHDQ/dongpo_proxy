package proxy

type Xor struct{}

func (encryption *Xor) Encode(input []byte) (int, []byte) {
	return len(input), input
}
func (encryption *Xor) Decode(input []byte) (int, []byte) {
	return len(input), input
}