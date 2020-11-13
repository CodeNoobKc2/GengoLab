package constvar

import (
	"path"
	"strconv"
)

const (
	PKG_0_PATH = "./pkg0"
	PKG_1_PATH = "./pkg1"

	LAB_PATH = "./labs"

	PACKAGE_BASE = "lab"
)

var (
	LAB_0_PATH = path.Join(LAB_PATH, "./lab0")
)

func LabPath(no int) string {
	strNo := strconv.FormatInt(int64(no), 10)
	return path.Join(LAB_PATH, "./lab"+strNo)
}

func LabResultPath(no int) string {
	return path.Join(LabPath(no), "./result")
}
