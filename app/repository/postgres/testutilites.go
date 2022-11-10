package postgres

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Requests []struct {
	Resuest string `yaml:"request"`
}

type TestData struct {
	Conf struct {
		Setup struct {
			Requests Requests
		}
		Teardown struct {
			Requests Requests
		}
	}
}

func loadTestDataFromYaml(file string) (TestData, error) {
	var data TestData
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return data, fmt.Errorf("repository.loadTestDataFromYaml: %w", err)
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return data, fmt.Errorf("repository.loadTestDataFromYaml: %w", err)
	}
	return data, nil
}
