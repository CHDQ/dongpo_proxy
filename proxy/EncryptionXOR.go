package proxy

type Xor struct{}

func (encryption *Xor) Encode(input []byte) (int, []byte) {
	for i := 0; i < len(input); i++ {
		input[i] = input[i] ^ 255
	}
	return len(input), input
}
func (encryption *Xor) Decode(input []byte) (int, []byte) {
	for i := 0; i < len(input); i++ {
		input[i] = input[i] ^ 255
	}
	return len(input), input
}
