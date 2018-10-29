package main

import (
	"fmt"

	"github.com/Masterminds/semver"

	kubermaticapiv1 "github.com/kubermatic/kubermatic/api/pkg/api/v1"
	"github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Returns a matrix of (version x operating system)
func getAWSScenarios() []testScenario {
	var scenarios []testScenario
	for _, v := range supportedVersions {
		// Ubuntu
		scenarios = append(scenarios, &awsScenario{
			version: v,
			nodeOsSpec: kubermaticapiv1.OperatingSystemSpec{
				Ubuntu: &kubermaticapiv1.UbuntuSpec{},
			},
		})
		// CoreOS
		scenarios = append(scenarios, &awsScenario{
			version: v,
			nodeOsSpec: kubermaticapiv1.OperatingSystemSpec{
				ContainerLinux: &kubermaticapiv1.ContainerLinuxSpec{
					// Otherwise the nodes restart directly after creation - bad for tests
					DisableAutoUpdate: true,
				},
			},
		})
		// CentOS
		//TODO: Fix
		//scenarios = append(scenarios, &awsScenario{
		//	version: v,
		//	nodeOsSpec: kubermaticapiv2.OperatingSystemSpec{
		//		CentOS: &kubermaticapiv2.CentOSSpec{},
		//	},
		//})
	}
	return scenarios
}

type awsScenario struct {
	version    *semver.Version
	nodeOsSpec kubermaticapiv1.OperatingSystemSpec
}

func (s *awsScenario) Name() string {
	return fmt.Sprintf("aws-%s-%s", getOSNameFromSpec(s.nodeOsSpec), s.version.String())
}

func (s *awsScenario) Cluster(secrets secrets) *v1.Cluster {
	return &v1.Cluster{
		ObjectMeta: metav1.ObjectMeta{},
		Spec: v1.ClusterSpec{
			Version:           s.version.String(),
			HumanReadableName: s.Name(),
			ClusterNetwork: v1.ClusterNetworkingConfig{
				Services: v1.NetworkRanges{
					CIDRBlocks: []string{"10.10.10.0/24"},
				},
				Pods: v1.NetworkRanges{
					CIDRBlocks: []string{"172.25.0.0/16"},
				},
				DNSDomain: "cluster.local",
			},
			Cloud: v1.CloudSpec{
				DatacenterName: "aws-eu-central-1a",
				AWS: &v1.AWSCloudSpec{
					SecretAccessKey: secrets.AWS.SecretAccessKey,
					AccessKeyID:     secrets.AWS.AccessKeyID,
				},
			},
		},
	}
}

func (s *awsScenario) Nodes(num int) []*kubermaticapiv1.Node {
	var nodes []*kubermaticapiv1.Node
	for i := 0; i < num; i++ {
		node := &kubermaticapiv1.Node{
			ObjectMeta: kubermaticapiv1.ObjectMeta{},
			Spec: kubermaticapiv1.NodeSpec{
				Cloud: kubermaticapiv1.NodeCloudSpec{
					AWS: &kubermaticapiv1.AWSNodeSpec{
						InstanceType: "t2.medium",
						VolumeType:   "gp2",
						VolumeSize:   50,
					},
				},
				Versions: kubermaticapiv1.NodeVersionInfo{
					Kubelet: s.version.String(),
				},
				OperatingSystem: s.nodeOsSpec,
			},
		}

		nodes = append(nodes, node)
	}

	return nodes
}
