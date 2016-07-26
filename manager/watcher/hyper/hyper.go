// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rkt implements the watcher interface for rkt
package hyper

import (
	"time"

	"github.com/google/cadvisor/container/hyper"
	"github.com/google/cadvisor/manager/watcher"

	"github.com/golang/glog"
)

const WatchInterval = 3 * time.Second

type hyperContainerWatcher struct {
	// Signal for watcher thread to stop.
	stopWatcher chan error
	isPod       bool
	client      *hyper.HyperClient
	// container watcher
	containers map[string]string
}

func NewHyperContainerWatcher() (watcher.ContainerWatcher, error) {
	watcher := &hyperContainerWatcher{
		stopWatcher: make(chan error),
		client:      hyper.NewHyperClient(),
		containers:  make(map[string]string),
	}

	return watcher, nil
}

func (self *hyperContainerWatcher) Start(events chan watcher.ContainerEvent) error {
	// Disable subcontainers watcher since we can't fetch container data now
	if self != nil {
		return nil
	}

	timer := time.NewTimer(WatchInterval)
	// TODO harry why need this?
	// if !self.isPod {
	// 	return nil
	// }

	go func(self *hyperContainerWatcher) {
		for {
			select {
			case <-self.stopWatcher:
				self.stopWatcher <- nil
				return
			case <-timer.C:
				containers, err := self.client.ListContainers()
				if err != nil {
					glog.Errorf("Error list hyper containers: %v", err)
					continue
				}

				newContainerMap := make(map[string]string)
				for _, c := range containers {
					// TODO harry why need this?
					// if c.podID != self.podID {
					// 	continue
					// }

					containerName := "/hyper/" + c.ContainerID
					newContainerMap[containerName] = containerName

					if _, ok := self.containers[containerName]; !ok {
						self.containers[containerName] = containerName
						// Deliver the event.
						events <- watcher.ContainerEvent{
							EventType:   watcher.ContainerAdd,
							Name:        containerName,
							WatchSource: watcher.Hyper,
						}
					}
				}

				for k := range self.containers {
					if _, ok := newContainerMap[k]; !ok {
						delete(self.containers, k)
						// Deliver the event.
						events <- watcher.ContainerEvent{
							EventType:   watcher.ContainerDelete,
							Name:        k,
							WatchSource: watcher.Hyper,
						}
					}
				}
			}
		}
	}(self)

	return nil
}

func (self *hyperContainerWatcher) Stop() error {
	// Rendezvous with the watcher thread.
	self.stopWatcher <- nil
	return <-self.stopWatcher
}
