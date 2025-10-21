package sm2

const (
	// encryptBlockSize SM2 加密后的块大小相对固定，可以根据实际情况调整
	encryptBlockSize = 4096 - 97

	// decryptBlockSize SM2 解密时也可以根据实际情况调整块大小
	decryptBlockSize = 4096
)
