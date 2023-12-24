package lib

import (
	"fmt"
	"math/rand"
	"os"
)

func SaveData1(data []byte, path string) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())

	file, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}

	if err = file.Sync(); err != nil {
		os.Remove(tmp)
		return err
	}

	if _, err = file.Write(data); err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}
