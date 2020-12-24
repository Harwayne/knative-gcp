// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"

	"cloud.google.com/go/iam/admin/apiv1"
	"github.com/google/knative-gcp/pkg/apis/configs/brokerdelivery"
	"github.com/google/knative-gcp/pkg/apis/configs/dataresidency"
	"github.com/google/knative-gcp/pkg/apis/configs/gcpauth"
	"github.com/google/knative-gcp/pkg/reconciler/broker"
	"github.com/google/knative-gcp/pkg/reconciler/brokercell"
	"github.com/google/knative-gcp/pkg/reconciler/deployment"
	"github.com/google/knative-gcp/pkg/reconciler/events/auditlogs"
	"github.com/google/knative-gcp/pkg/reconciler/events/build"
	"github.com/google/knative-gcp/pkg/reconciler/events/pubsub"
	"github.com/google/knative-gcp/pkg/reconciler/events/scheduler"
	"github.com/google/knative-gcp/pkg/reconciler/events/storage"
	"github.com/google/knative-gcp/pkg/reconciler/identity/iam"
	"github.com/google/knative-gcp/pkg/reconciler/intevents/pullsubscription/keda"
	"github.com/google/knative-gcp/pkg/reconciler/intevents/pullsubscription/static"
	"github.com/google/knative-gcp/pkg/reconciler/intevents/topic"
	"github.com/google/knative-gcp/pkg/reconciler/messaging/channel"
	"github.com/google/knative-gcp/pkg/reconciler/trigger"
	"knative.dev/pkg/injection"
)

import (
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// Injectors from wire.go:

func InitializeControllers(ctx context.Context) ([]injection.ControllerConstructor, error) {
	v := ClientOptions()
	iamClient, err := admin.NewIamClient(ctx, v...)
	if err != nil {
		return nil, err
	}
	iamPolicyManager, err := iam.NewIAMPolicyManager(ctx, iamClient)
	if err != nil {
		return nil, err
	}
	storeSingleton := &gcpauth.StoreSingleton{}
	constructor := auditlogs.NewConstructor(iamPolicyManager, storeSingleton)
	storageConstructor := storage.NewConstructor(iamPolicyManager, storeSingleton)
	schedulerConstructor := scheduler.NewConstructor(iamPolicyManager, storeSingleton)
	pubsubConstructor := pubsub.NewConstructor(iamPolicyManager, storeSingleton)
	buildConstructor := build.NewConstructor(iamPolicyManager, storeSingleton)
	staticConstructor := static.NewConstructor(iamPolicyManager, storeSingleton)
	kedaConstructor := keda.NewConstructor(iamPolicyManager, storeSingleton)
	dataresidencyStoreSingleton := &dataresidency.StoreSingleton{}
	topicConstructor := topic.NewConstructor(iamPolicyManager, storeSingleton, dataresidencyStoreSingleton)
	channelConstructor := channel.NewConstructor(dataresidencyStoreSingleton)
	triggerConstructor := trigger.NewConstructor(dataresidencyStoreSingleton)
	brokerdeliveryStoreSingleton := &brokerdelivery.StoreSingleton{}
	brokerConstructor := broker.NewConstructor(brokerdeliveryStoreSingleton, dataresidencyStoreSingleton)
	deploymentConstructor := deployment.NewConstructor()
	brokercellConstructor := brokercell.NewConstructor()
	v2 := Controllers(constructor, storageConstructor, schedulerConstructor, pubsubConstructor, buildConstructor, staticConstructor, kedaConstructor, topicConstructor, channelConstructor, triggerConstructor, brokerConstructor, deploymentConstructor, brokercellConstructor)
	return v2, nil
}
