package lib

import (
	"os"
)

func LogCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
}

func LogAppend(file *os.File, line string) error {
	line = line + "\n"
	buf := []byte(line)

	if _, err := file.Write(buf); err != nil {
		return err
	}
	return file.Sync()
}

func Run() {

	lines := []string{
		"Greeting from line 1",
		"Greeting from line 2",
		"Greeting from line 3",
		"Greeting from line 4",
		"Greeting from line 5",
		"Greeting from line 6",
		"Greeting from line 7",
		"Greeting from line 8",
		"Greeting from line 9",
		"Greeting from line 10",
	}

	logs, err := LogCreate("./test.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		LogAppend(logs, line)
	}
}
