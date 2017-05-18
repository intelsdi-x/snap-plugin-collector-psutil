//
// +build small

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2015 Intel Corporation

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

package psutil

import (
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPSUtilPlugin(t *testing.T) {
	Convey("Create PSUtil Collector", t, func() {
		psCol := NewPsutilCollector()
		Convey("So psCol should not be nil", func() {
			So(psCol, ShouldNotBeNil)
		})
		Convey("So psCol should be of Psutil type", func() {
			So(psCol, ShouldHaveSameTypeAs, &Psutil{})
		})
		Convey("psCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := psCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a cpolicy.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})
}
