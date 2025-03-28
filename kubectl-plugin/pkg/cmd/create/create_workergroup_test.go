package create

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/ptr"

	"github.com/ray-project/kuberay/kubectl-plugin/pkg/util"

	rayv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
)

func TestCreateWorkerGroupSpec(t *testing.T) {
	rayStartParams := map[string]string{"dashboard-host": "0.0.0.0", "num-cpus": "2"}
	options := &CreateWorkerGroupOptions{
		groupName:         "example-group",
		image:             "DEADBEEF",
		workerReplicas:    3,
		workerMinReplicas: 1,
		workerMaxReplicas: 5,
		workerCPU:         "2",
		workerMemory:      "5Gi",
		workerGPU:         "1",
		rayStartParams:    rayStartParams,
	}

	expected := rayv1.WorkerGroupSpec{
		RayStartParams: rayStartParams,
		GroupName:      "example-group",
		Template: corev1.PodTemplateSpec{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "ray-worker",
						Image: "DEADBEEF",
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:     resource.MustParse("2"),
								corev1.ResourceMemory:  resource.MustParse("5Gi"),
								util.ResourceNvidiaGPU: resource.MustParse("1"),
							},
							Limits: corev1.ResourceList{
								corev1.ResourceCPU:     resource.MustParse("2"),
								corev1.ResourceMemory:  resource.MustParse("5Gi"),
								util.ResourceNvidiaGPU: resource.MustParse("1"),
							},
						},
					},
				},
			},
		},
		Replicas:    ptr.To[int32](3),
		MinReplicas: ptr.To[int32](1),
		MaxReplicas: ptr.To[int32](5),
	}

	assert.Equal(t, expected, createWorkerGroupSpec(options))
}
