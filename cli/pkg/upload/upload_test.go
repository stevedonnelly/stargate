// Copyright DataStax, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package upload

import (
	"testing"
	"time"

	"github.com/datastax/stargate/cli/pkg/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UploadSuite struct {
	suite.Suite
	client docker.Client
}

func (suite *UploadSuite) SetupSuite() {
	client, err := docker.NewClient()
	if err != nil {
		panic(1)
	}
	err = client.StartCassandra(&docker.StartCassandraOptions{
		DockerImageHost: "docker.io/library/",
		ImageName:       "cassandra",
	})
	if err != nil {
		panic(1)
	}
	err = client.StartService(&docker.StartServiceOptions{
		CassandraURL:    "stargate-cassandra",
		ExposedPorts:    []string{"8080"},
		DockerImageHost: "docker.io/",
		ImageName:       "service",
	})
	if err != nil {
		panic(1)
	}
	suite.client = client
	time.Sleep(20 * time.Second)
}

func (suite *UploadSuite) TearDownSuite() {
	suite.client.Remove("service")
	suite.client.Remove("cassandra")
}

const validHost = "http://localhost:8080/test"
const validPath = "../../../src/main/resources/schema.conf"

func (suite *UploadSuite) TestInvalidPathInput() {
	err := Upload("", validHost)
	assert.NotNil(suite.T(), err)
}

func (suite *UploadSuite) TestMalformedHostInput() {
	err := Upload(validPath, "\n")
	assert.NotNil(suite.T(), err)
}

func (suite *UploadSuite) TestBadHostInput() {
	err := Upload(validPath, "")
	assert.NotNil(suite.T(), err)
}

func (suite *UploadSuite) TestBadPathInput() {
	err := Upload("./upload_test.go", validHost)
	assert.NotNil(suite.T(), err)
}

func (suite *UploadSuite) TestValidInput() {
	err := Upload(validPath, validHost)
	assert.Nil(suite.T(), err)
}

func TestUpload(t *testing.T) {
	suite.Run(t, new(UploadSuite))
}