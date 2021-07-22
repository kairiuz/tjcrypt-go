package tjcrypt

func checkHeader(buf []byte) bool {
	return buf[0] == 't' && buf[1] == 'j' && buf[2] == '!'
}
