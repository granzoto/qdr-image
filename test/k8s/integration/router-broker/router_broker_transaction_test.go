//+build router_broker transaction integration

package router_broker

import (
	"github.com/interconnectedcloud/qdr-image/test/k8s/integration/router-broker/common"
	"github.com/interconnectedcloud/qdr-image/test/k8s/utils/k8s"
	"github.com/skupperproject/skupper/test/utils/base"
	"github.com/skupperproject/skupper/test/utils/constants"
	k8s2 "github.com/skupperproject/skupper/test/utils/k8s"
	"gotest.tools/assert"
	v1 "k8s.io/api/core/v1"

	"testing"
)

func TestTransaction(t *testing.T) {
	clusterNeeds := base.ClusterNeeds{
		NamespaceId:     "router-broker-trx",
		PublicClusters:  1,
		PrivateClusters: 0,
	}

	defer Teardown(t)
	base.HandleInterruptSignal(t, func(t *testing.T) {
		Teardown(t)
	})
	Setup(t, clusterNeeds)

	// Cluster context
	ctx, err := common.TestRunner.GetPublicContext(1)
	assert.Assert(t, err)

	// Preparing the jms-amqp-tests job
	jmsAmqpTests := k8s.NewJob("jms-amqp-tests", ctx.Namespace, k8s.JobOpts{
		Image:        "docker.io/fgiorgetti/jms-amqp-tests",
		BackoffLimit: 1,
		Restart:      v1.RestartPolicyNever,
		Env: map[string]string{
			"QPID_JMS_TRANSACTION_ROUTER_URL": "amqp://router:5672",
		},
		Labels: map[string]string{
			"app": "jms-amqp-tests",
		},
	})

	// Running the job
	_, err = ctx.VanClient.KubeClient.BatchV1().Jobs(ctx.Namespace).Create(jmsAmqpTests)
	assert.Assert(t, err)

	// Waiting for job to complete
	job, err := k8s2.WaitForJob(ctx.Namespace, ctx.VanClient.KubeClient, jmsAmqpTests.Name, constants.ImagePullingAndResourceCreationTimeout)
	assert.Assert(t, err)
	k8s2.AssertJob(t, job)

}
