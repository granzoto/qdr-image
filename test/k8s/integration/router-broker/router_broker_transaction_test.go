//+build router_broker transaction

package router_broker

import (
	"github.com/skupperproject/skupper/test/utils/base"
	"testing"
)

func TestTransaction(t *testing.T) {
	clusterNeeds := base.ClusterNeeds{
		NamespaceId:     "router-broker-trx",
		PublicClusters:  1,
		PrivateClusters: 0,
	}

	defer Teardown(t)
	Setup(t, clusterNeeds)
}
