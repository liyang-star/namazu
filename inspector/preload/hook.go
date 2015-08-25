// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"time"
)

func tryInjectSleep() {
	maxSleep := config.GetDuration("max_sleep")
	if maxSleep > 0 {
		determinedSleep := time.Duration(rand.Int63n(int64(maxSleep)))
		l := log.WithFields(log.Fields{
			"max_sleep":        maxSleep,
			"determined_sleep": determinedSleep})
		l.Debug("Sleeping")
		time.Sleep(determinedSleep)
		l.Debug("Slept")
	}
}

func commonPrehook(name string, args ...interface{}) (bool, int, error) {
	tryInjectSleep()
	// TODO: try inject errno
	return false, 0, nil
}

func commonPosthook(lowerRet int, lowerError error, name string, args ...interface{}) (bool, int, error) {
	return false, lowerRet, lowerError
}
