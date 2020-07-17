// +build wireinject

/*
Copyright 2020 Google LLC.

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

package main

import (
	"context"

	"github.com/google/knative-gcp/pkg/broker/config/volume"
	"github.com/google/knative-gcp/pkg/broker/handler"
	"github.com/google/knative-gcp/pkg/metrics"
	"github.com/google/knative-gcp/pkg/utils/clients"
	"github.com/google/wire"
)

// InitializeSyncPool initializes the retry sync pool. Uses the given projectID to initialize the
// retry pool's pubsub client and uses targetsVolumeOpts to initialize the targets volume watcher.
func InitializeSyncPool(
	ctx context.Context,
	projectID clients.ProjectID,
	podName metrics.PodName,
	containerName metrics.ContainerName,
	targetsVolumeOpts []volume.Option,
	opts ...handler.Option) (*handler.RetryPool, error) {
	// Implementation generated by wire. Providers for required RetryPool dependencies should be
	// added here.
	panic(wire.Build(handler.ProviderSet, volume.NewTargetsFromFile, metrics.NewDeliveryReporter))
}
