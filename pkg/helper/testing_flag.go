package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

var (
	// default flag path
	flagPath = "./pkg/helper/testing_flag.json"
)

type TestingFlag struct {
	LocalTestingFlag       bool
	ProvisionTestingFlag   bool
	DeprovisionTestingFlag bool
	LocalTestStitchingIP   string

	FakeTransactionUtil         float64
	FakeNodeComputeResourceUtil float64
	FakePodComputeResourceUtil  float64
	FakeApplicationCpuUsed      float64
	FakeApplicationMemUsed      float64
}

func SetPath(path string) {
	if _, err := os.Stat(path); err == nil {
		flagPath = path
		glog.V(2).Infof("VMT testing flag is load from file %s", path)
	} else {
		glog.V(3).Infof("%s does not exist.", path)
	}
}

func LoadTestingFlag() (*TestingFlag, error) {
	file, err := ioutil.ReadFile(flagPath)
	if err != nil {
		glog.V(4).Infof("ERROR! : %v\n", err)
		return nil, err
	}
	var flags TestingFlag
	json.Unmarshal(file, &flags)
	glog.V(5).Infof("Results: %v\n", flags)
	return &flags, nil
}

func IsActionTesting() (bool, error) {
	flag, err := LoadTestingFlag()
	if err != nil {
		return false, err
	}
	return flag.ProvisionTestingFlag, nil
}

func IsDeprovisionTesting() (bool, error) {
	flag, err := LoadTestingFlag()
	if err != nil {
		return false, err
	}
	return flag.DeprovisionTestingFlag, nil
}
