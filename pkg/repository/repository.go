// kibosh
//
// Copyright (c) 2017-Present Pivotal Software, Inc. All Rights Reserved.
//
// This program and the accompanying materials are made available under the terms of the under the Apache License,
// Version 2.0 (the "License”); you may not use this file except in compliance with the License. You may
// obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the
// License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing permissions and
// limitations under the License.

package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/lager"
	"github.com/cf-platform-eng/kibosh/pkg/helm"
	"github.com/pkg/errors"
	"k8s.io/helm/pkg/chartutil"
)

//go:generate counterfeiter ./ Repository
type Repository interface {
	LoadCharts() ([]*helm.MyChart, error)
	SaveChart(path string) error
	DeleteChart(name string) error
}

type repository struct {
	helmChartDir          string
	privateRegistryServer string
	logger                lager.Logger
}

func NewRepository(chartPath string, privateRegistryServer string, logger lager.Logger) Repository {
	return &repository{
		helmChartDir:          chartPath,
		privateRegistryServer: privateRegistryServer,
		logger:                logger,
	}
}

func (r *repository) LoadCharts() ([]*helm.MyChart, error) {
	charts := []*helm.MyChart{}

	chartExists, err := fileExists(filepath.Join(r.helmChartDir, "Chart.yaml"))
	if err != nil {
		return charts, err
	}

	if chartExists {
		myChart, err := helm.NewChart(r.helmChartDir, r.privateRegistryServer)
		if err != nil {
			return charts, err
		}
		charts = append(charts, myChart)
	} else {
		helmDirFiles, err := ioutil.ReadDir(r.helmChartDir)
		if err != nil {
			return charts, err
		}
		for _, fileInfo := range helmDirFiles {
			if fileInfo.Name() == "workspace_tmp" {
				//rename doesn't support moving things across disks, so we're expanding to a working dir
				continue
			}
			if fileInfo.IsDir() {
				subChartPath := filepath.Join(r.helmChartDir, fileInfo.Name())
				subdirChartExists, err := fileExists(filepath.Join(subChartPath, "Chart.yaml"))
				if err != nil {
					return charts, err
				}
				if subdirChartExists {
					myChart, err := helm.NewChart(filepath.Join(subChartPath), r.privateRegistryServer)
					if err != nil {
						return charts, err
					}
					charts = append(charts, myChart)
				} else {
					r.logger.Info(fmt.Sprintf("[%s] does not contain Chart.yml, skipping", subChartPath))
				}
			}
		}
	}

	return charts, nil
}

func (r *repository) SaveChart(path string) error {
	expandedTarPath := filepath.Join(r.helmChartDir, "workspace_tmp")
	err := os.RemoveAll(expandedTarPath)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	err = os.Mkdir(expandedTarPath, 0700)
	if err != nil {
		return err
	}

	err = chartutil.ExpandFile(expandedTarPath, path)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(expandedTarPath)
	var chartPathInfo os.FileInfo
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			if chartPathInfo != nil {
				return errors.New("Multiple directories found in uploaded archive")
			} else {
				chartPathInfo = file
			}
		}
	}

	chartPath := filepath.Join(expandedTarPath, chartPathInfo.Name())
	chart, err := helm.NewChart(chartPath, r.privateRegistryServer)
	if err != nil {
		return err
	}

	destinationPath := filepath.Join(r.helmChartDir, chartPathInfo.Name())
	info, _ := os.Stat(destinationPath)
	if info != nil {
		os.RemoveAll(destinationPath)
	}

	if chartPathInfo.Name() != chart.Metadata.Name {
		return errors.New("Chart metadata name and top level directory in archive for chart does not match")
	}

	err = os.Rename(chartPath, filepath.Join(r.helmChartDir, chartPathInfo.Name()))
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteChart(name string) error {

	deletePath := filepath.Join(r.helmChartDir, name)

	_, err := os.Stat(deletePath)

	if os.IsNotExist(err) {
		r.logger.Info(fmt.Sprintf("[%s] does not exist, skipping", deletePath))
		return nil
	} else if err != nil {
		r.logger.Info(fmt.Sprintf("[%s] error reading at path, skipping", deletePath))
		return err
	}
	os.RemoveAll(deletePath)

	return nil

}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
}
