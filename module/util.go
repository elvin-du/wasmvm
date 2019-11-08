package module

import "io"

func readBytes(r io.Reader, n uint32) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := io.ReadFull(r, bytes)
	if err != nil {
		return bytes, err
	}

	return bytes, nil
}