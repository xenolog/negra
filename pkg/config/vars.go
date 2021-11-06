/*
Copyright © 2020 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
	"k8s.io/klog"
)

var (
	BotConfig      *Config
	ConfigFilePath string
)

type Topology map[int]map[int]int

type MtProtoAuth struct {
	AppID         int    `yaml:"appID,omitempty"`
	AppHash       string `yaml:"appHash,omitempty"`
	PhoneNumber   string `yaml:"phoneNumber,omitempty"`
	DeviceModel   string `yaml:"deviceModel,omitempty"`
	SystemVersion string `yaml:"systemVersion,omitempty"`
	AppVersion    string `yaml:"appVersion,omitempty"`
	Server        string `yaml:"server,omitempty"`
}

type Channel struct {
	ID         int32 `yaml:"ID,omitempty"`
	AccessHash int64 `yaml:"accessHash,omitempty"`
}

// Shared config for all cloud providers
type Config struct {
	BotAPIkey    string      `yaml:"botApiKey,omitempty"`
	MainChatHash string      `yaml:"mainChatHash,omitempty"`
	Channel      Channel     `yaml:"channel,omitempty"`
	MtProto      MtProtoAuth `yaml:"mtproto,omitempty"`
	Users        []string
	Admins       []string
	Map          Topology
	Workdir      string
}

func (c *Config) Parse(configFilePath string) error {
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("could not open config file %s: %w", configFilePath, err)
	}

	err = yaml.Unmarshal(configFile, &c)
	if err != nil {
		return fmt.Errorf("failed to parse config file %s: %w", configFilePath, err)
	}
	return nil
}

func (c *Config) String() (rv string) {
	if xxx, err := yaml.Marshal(c); err != nil {
		klog.Fatal(err)
	} else {
		rv = string(xxx)
	}
	return rv
}

func init() {
	BotConfig = &Config{}
}
