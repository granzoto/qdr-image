package router_broker

import (
	"github.com/interconnectedcloud/qdr-image/test/k8s/integration/router-broker/common"
	"github.com/skupperproject/skupper/test/utils/base"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	common.TestRunner.Initialize(m)
}

func Setup(t *testing.T, needs base.ClusterNeeds) {
	var err error
	t.Logf("Building ClusterTestRunner for %s", needs.NamespaceId)
	common.TestRunner.Build(t, needs, nil)

	t.Logf("%s - starting topology setup", time.Now().String())

	//
	// - Creating the namespace
	//
	err = common.TestRunner.GetPublicContext(1).CreateNamespace()
	assert.Assert(t, err)

	//
	// - Deploying the Broker
	//


	t.Logf("%s - setup is complete", time.Now().String())
}

func Teardown(t *testing.T) {
	t.Logf("%s - starting topology teardown", time.Now().String())
	t.Logf("%s - teardown is complete", time.Now().String())
}
