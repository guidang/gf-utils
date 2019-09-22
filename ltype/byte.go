package ltype

// ByteToDword 无符号四字节整型
func ByteToDword(byte0, byte1, byte2, byte3 int) int {
	result := ((byte3 & 0xFF) << 24) | ((byte2 & 0xFF) << 16) | ((byte1 & 0xFF) << 8) | (byte0 & 0xFF)
	return result
}

// ByteToWord 无符号双字节整型
func ByteToWord(byte0, byte1 int) int {
	result := ((byte1 & 0xFF) << 8) | (byte0 & 0xFF)
	return result
}
