package registry

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"sort"

	"github.com/hhkbp2/go-logging"
)

// SetupLogging configure logging.
func SetupLogging(name string) logging.Logger {
	logger := logging.GetLogger(name)
	handler := logging.NewStdoutHandler()
	format := "%(asctime)s - %(name)s - %(levelname)s - %(message)s"
	dateFormat := "%Y-%m-%d %H:%M:%S"
	formatter := logging.NewStandardFormatter(format, dateFormat)
	handler.SetFormatter(formatter)
	logger.SetLevel(logging.LevelInfo)
	logger.AddHandler(handler)
	return logger
}

// SortedMapKeys sort keys of the map where values can be of any type.
func SortedMapKeys(m interface{}) []string {
	v := reflect.ValueOf(m)
	keys := make([]string, 0, len(v.MapKeys()))
	for _, key := range v.MapKeys() {
		keys = append(keys, key.String())
	}
	sort.Strings(keys)
	return keys
}

// PrettySize format bytes in more readable units.
func PrettySize(size float64) string {
	units := []string{"B", "KB", "MB", "GB"}
	i := 0
	for size > 1024 && i < len(units) {
		size = size / 1024
		i = i + 1
	}
	return fmt.Sprintf("%.*f %s", 0, size, units[i])
}

// ItemInSlice check if item is an element of slice
func ItemInSlice(item string, slice []string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

// RunCommand run command in system
func RunCommand(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "Run command finished with error: ", err
	}
	return stdout.String(), nil
}

// PathExists if a path is exists.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
