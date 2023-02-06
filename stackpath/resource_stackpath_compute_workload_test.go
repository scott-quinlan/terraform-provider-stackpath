package stackpath

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/workload/workload_client/workloads"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/workload/workload_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	IPv4IPFamilies      = []string{"IPv4"}
	DualStackIPFamilies = []string{"IPv4", "IPv6"}
)

func TestComputeWorkloadContainers(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	// By design, the StackPath API does not return image pull secrets to the
	// user when retrieving a workload. Expect to see an empty secret in the
	// result and suppress the diff error.
	//emptyImagePullSecrets := regexp.MustCompile("(.*)image_pull_credentials.0.docker_registry.0.password:(\\s*)\"\" => \"secret registry password\"(.*)")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerBasic(nameSuffix, nil, IPv4IPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddPorts(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "https", "TCP", 443, true),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "some value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerRemoveEnvVar(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddProbes(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigContainerImagePullCredentials(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckImagePullCredentials(workload, "docker.io", "my-registry-user", "developers@stackpath.com"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigAutoScalingConfiguration(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckTargetAutoScaling(workload, "us", "cpu", 2, 4, 50),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			// TODO: there's a ordering issue where the order of the containers is shuffled when being read in from the API
			//   Need to ensure consistent ordering of containers when reading in state.
			//
			// {
			// 	Config: testComputeWorkloadConfigContainerAddContainer(),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app-2", "nginx:v1.15.0"),
			// 	),
			// },
		},
	})
}

func TestComputeWorkloadContainersWithOneToOneNATEnabled(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "nat-" + strconv.Itoa(int(time.Now().Unix()))
	oneToOneNAT := true

	// By design, the StackPath API does not return image pull secrets to the
	// user when retrieving a workload. Expect to see an empty secret in the
	// result and suppress the diff error.
	//emptyImagePullSecrets := regexp.MustCompile("(.*)image_pull_credentials.0.docker_registry.0.password:(\\s*)\"\" => \"secret registry password\"(.*)")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerBasic(nameSuffix, &oneToOneNAT, IPv4IPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddPorts(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "https", "TCP", 443, true),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "some value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerRemoveEnvVar(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddProbes(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigContainerImagePullCredentials(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckImagePullCredentials(workload, "docker.io", "my-registry-user", "developers@stackpath.com"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigAutoScalingConfiguration(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckTargetAutoScaling(workload, "us", "cpu", 2, 4, 50),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
			// TODO: there's a ordering issue where the order of the containers is shuffled when being read in from the API
			//   Need to ensure consistent ordering of containers when reading in state.
			//
			// {
			// 	Config: testComputeWorkloadConfigContainerAddContainer(),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app-2", "nginx:v1.15.0"),
			// 	),
			// },
		},
	})
}

func TestComputeWorkloadContainersAdditionalVolume(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerAddVolumeMounts(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo-volume", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadAdditionalVolume(workload, "volume", "10Gi"),
					testAccComputeWorkloadContainerVolumeMount(workload, "app", "volume", "/var/log"),
				),
			},
		},
	})
}

func TestComputeWorkloadContainersIPv6(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "ipv6-" + strconv.Itoa(int(time.Now().Unix()))
	oneToOneNAT := true

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerBasic(nameSuffix, &oneToOneNAT, DualStackIPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, DualStackIPFamilies),
				),
			},
		},
	})
}

func TestComputeWorkloadVirtualMachines(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "ipv6-" + strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigVirtualMachineBasic(nameSuffix, nil, IPv4IPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.bar", workload),
					testAccComputeWorkloadCheckVirtualMachineImage(workload, "app", "stackpath-edge/centos-7:v201905012051"),
					testAccComputeWorkloadCheckVirtualMachinePort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, IPv4IPFamilies),
				),
			},
		},
	})
}

func TestComputeWorkloadVirtualMachinesIPv6(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigVirtualMachineBasic(nameSuffix, nil, DualStackIPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.bar", workload),
					testAccComputeWorkloadCheckVirtualMachineImage(workload, "app", "stackpath-edge/centos-7:v201905012051"),
					testAccComputeWorkloadCheckVirtualMachinePort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, DualStackIPFamilies),
				),
			},
		},
	})
}

func testAccComputeWorkloadCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_workload" {
				continue
			}

			resp, err := config.edgeCompute.Workloads.GetWorkload(&workloads.GetWorkloadParams{
				StackID:    config.StackID,
				WorkloadID: rs.Primary.ID,
				Context:    context.Background(),
			}, nil)
			// Since compute workloads are deleted asynchronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && resp.Payload.Workload.Metadata.DeleteRequestedAt == nil {
				return fmt.Errorf("compute workload still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccComputeWorkloadCheckInterface(
	workload *workload_models.V1Workload,
	interfaceIndex int,
	networkName string,
	enableOneToOneNAT bool,
	ipFamilies []string,
) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		interfaces := workload.Spec.NetworkInterfaces
		if interfaceIndex < 0 {
			return fmt.Errorf("invalid interface index to check: %d", interfaceIndex)
		}
		if interfaceIndex >= len(interfaces) {
			return fmt.Errorf("could not find the interface index %d/%d", interfaceIndex, len(interfaces))
		}
		inter := interfaces[interfaceIndex]
		if inter.Network != networkName {
			return fmt.Errorf("invalid network on interface %d. got=%s want=%s", interfaceIndex, inter.Network, networkName)
		}
		if inter.EnableOneToOneNat != enableOneToOneNAT {
			return fmt.Errorf("invalid enableOneToOneNat on interface %d. got=%v want=%v", interfaceIndex, inter.EnableOneToOneNat, enableOneToOneNAT)
		}
		if len(inter.IPFamilies) > 0 {
			ipFamiliesStrList := make([]string, len(inter.IPFamilies))
			for i, ipFamily := range inter.IPFamilies {
				ipFamiliesStrList[i] = string(*ipFamily)
			}
			if !reflect.DeepEqual(ipFamiliesStrList, ipFamilies) {
				return fmt.Errorf("invalid ipFamilies on interface %d. got=%v want=%v", interfaceIndex, ipFamiliesStrList, ipFamilies)
			}
		}
		return nil
	}
}

func testAccComputeWorkloadContainerVolumeMount(workload *workload_models.V1Workload, containerName, volumeSlug, mountPath string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		container, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		}
		var mount *workload_models.V1InstanceVolumeMount
		for _, v := range container.VolumeMounts {
			if v.Slug == volumeSlug {
				mount = v
				break
			}
		}
		if mount == nil {
			return fmt.Errorf("container volume mount not found: %s", volumeSlug)
		} else if mount.MountPath != mountPath {
			return fmt.Errorf("mount path '%s' does not match expected '%s'", mount.MountPath, mountPath)
		}
		return nil
	}
}

func testAccComputeWorkloadAdditionalVolume(workload *workload_models.V1Workload, volumeName, size string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		var volume *workload_models.V1VolumeClaim
		for _, v := range workload.Spec.VolumeClaimTemplates {
			if v.Name == volumeName {
				volume = v
				break
			}
		}
		if volume == nil {
			return fmt.Errorf("could not find volume: %s", volumeName)
		} else if volume.Spec.Resources.Requests["storage"] != size {
			return fmt.Errorf("volume size '%s' does not match expected '%s'", volume.Spec.Resources.Requests["storage"], size)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckVirtualMachinePort(workload *workload_models.V1Workload, vmName, portName, protocol string, portNum int32) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if vm, found := workload.Spec.VirtualMachines[vmName]; !found {
			return fmt.Errorf("virtual machine was not found: %s", vmName)
		} else if port, found := vm.Ports[portName]; !found {
			return fmt.Errorf("virtual machine port not found: %s", portName)
		} else if port.Protocol != protocol {
			return fmt.Errorf("virtual machine port protocol '%s' does not match expected '%s'", port.Protocol, protocol)
		} else if port.Port != portNum {
			return fmt.Errorf("virtual machine port '%d' does not match expected '%d'", port.Port, portNum)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckVirtualMachineImage(workload *workload_models.V1Workload, name, image string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if vm, found := workload.Spec.VirtualMachines[name]; !found {
			return fmt.Errorf("virtual machine was not found: %s", name)
		} else if vm.Image != image {
			return fmt.Errorf("virtual machine image '%s' does not match expected '%s'", vm.Image, image)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckNoImagePullCredentials(workload *workload_models.V1Workload) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if len(workload.Spec.ImagePullCredentials) != 0 {
			return fmt.Errorf("unexpected image pull credentials set on the workload")
		}
		return nil
	}
}

func testAccComputeWorkloadCheckImagePullCredentials(workload *workload_models.V1Workload, server, username, email string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if len(workload.Spec.ImagePullCredentials) == 0 {
			return fmt.Errorf("no image pull credentials set on the workload")
		} else if creds := workload.Spec.ImagePullCredentials[0]; creds.DockerRegistry.Server != server {
			return fmt.Errorf("image pull server '%s' does not match expected value '%s'", creds.DockerRegistry.Server, server)
		} else if creds.DockerRegistry.Email != email {
			return fmt.Errorf("image pull email '%s' does not match expected value '%s'", creds.DockerRegistry.Email, email)
		} else if creds.DockerRegistry.Username != username {
			return fmt.Errorf("image pull username '%s' does not match expected value '%s'", creds.DockerRegistry.Username, username)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckTarget(workload *workload_models.V1Workload, targetName, scope, operator string, minReplicas int32, values ...string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		sort.Strings(values)
		if target, found := workload.Targets[targetName]; !found {
			return fmt.Errorf("target not found: %s", targetName)
		} else if target.Spec.DeploymentScope != scope {
			return fmt.Errorf("target scope '%s' does not match expected scope '%s'", target.Spec.DeploymentScope, scope)
		} else if deployment := target.Spec.Deployments; deployment.MinReplicas != minReplicas {
			return fmt.Errorf("target min replicas '%d' does not match expected '%d'", target.Spec.Deployments.MinReplicas, minReplicas)
		} else if selector := target.Spec.Deployments.Selectors[0]; selector.Operator != operator {
			return fmt.Errorf("target selector operator '%s' does not match expected operator '%s'", selector.Operator, operator)
		} else if sort.Strings(selector.Values); !reflect.DeepEqual(values, selector.Values) {
			return fmt.Errorf("target selector values %v do not match expected values %v", selector.Values, values)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerEnvVar(workload *workload_models.V1Workload, containerName, envVar, value string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if envVarSpec, found := containerSpec.Env[envVar]; !found {
			return fmt.Errorf("environment variable not found: %s", envVar)
		} else if envVarSpec.Value != value {
			return fmt.Errorf(`environment variable '%s="%s"' does not match expected value '%s'`, envVar, envVarSpec.Value, value)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerEnvVarNotExist(workload *workload_models.V1Workload, containerName, envVar string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if _, found := containerSpec.Env[envVar]; found {
			return fmt.Errorf("unexpected environment variable found: %s", envVar)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerPort(workload *workload_models.V1Workload, containerName, portName, protocol string, port int32, enableImplicitNetworkPolicy bool) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if portSpec, found := containerSpec.Ports[portName]; !found {
			return fmt.Errorf("port not found: %s", portName)
		} else if portSpec.Port != port {
			return fmt.Errorf("port number '%d' does not match expected port '%d'", portSpec.Port, port)
		} else if portSpec.Protocol != protocol {
			return fmt.Errorf("port protocol '%s' does not match expected protocol '%s'", portSpec.Protocol, protocol)
		} else if portSpec.EnableImplicitNetworkPolicy != enableImplicitNetworkPolicy {
			return fmt.Errorf("port enable implicit network policy '%t' does not match expected enable implicit network policy '%t'", portSpec.EnableImplicitNetworkPolicy, enableImplicitNetworkPolicy)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerPortNotExist(workload *workload_models.V1Workload, containerName, portName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if _, found := containerSpec.Ports[portName]; found {
			return fmt.Errorf("unexpected port found: %s", portName)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerImage(workload *workload_models.V1Workload, containerName, image string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if containerSpec, found := workload.Spec.Containers[containerName]; !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if containerSpec.Image != image {
			return fmt.Errorf("container image '%s' does not match expected '%s'", containerSpec.Image, image)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckExists(name string, workload *workload_models.V1Workload) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.edgeCompute.Workloads.GetWorkload(&workloads.GetWorkloadParams{
			WorkloadID: rs.Primary.ID,
			StackID:    config.StackID,
			Context:    context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve workload: %v", err)
		}

		*workload = *found.Payload.Workload

		return nil
	}
}

func testAccComputeWorkloadCheckTargetAutoScaling(workload *workload_models.V1Workload, targetName, metric string, minReplicas, maxReplicas, averageUtilization int32) resource.TestCheckFunc {
	return func(*terraform.State) error {
		target, found := workload.Targets[targetName]
		if !found {
			return fmt.Errorf("target not found: %s", targetName)
		} else if target.Spec.Deployments.MinReplicas != minReplicas {
			return fmt.Errorf("expected %d min replicas, got %d", minReplicas, target.Spec.Deployments.MinReplicas)
		} else if target.Spec.Deployments.MaxReplicas != maxReplicas {
			return fmt.Errorf("expected %d max replicas, got %d", maxReplicas, target.Spec.Deployments.MaxReplicas)
		} else if target.Spec.Deployments.ScaleSettings == nil {
			return fmt.Errorf("expected scale settings to be non-nil")
		}

		for _, m := range target.Spec.Deployments.ScaleSettings.Metrics {
			if m.Metric != metric {
				continue
			} else if m.AverageUtilization != averageUtilization {
				return fmt.Errorf("expected average utilization to be %d, got %d", averageUtilization, m.AverageUtilization)
			}
		}
		return nil
	}
}

func printSlice(a []string) string {
	q := make([]string, len(a))
	for i, s := range a {
		q[i] = fmt.Sprintf("%q", s)
	}
	return fmt.Sprintf("[%s]", strings.Join(q, ", "))
}

func getInterface(network string, enableNAT *bool, ipFamilies []string) string {
	var config string
	if enableNAT == nil {
		config = fmt.Sprintf(`
    network_interface {
      network = "%s"
	  ip_families = %s
    }`, network, printSlice(ipFamilies))
	} else {
		config = fmt.Sprintf(`
    network_interface {
      network = "%s"
      enable_one_to_one_nat = %v
	  ip_families = %s
    }`, network, *enableNAT, printSlice(ipFamilies))
	}
	return config
}

func testComputeWorkloadConfigContainerBasic(suffix string, enableNAT *bool, ipFamilies []string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, ipFamilies))
}

func testComputeWorkloadConfigContainerAddPorts(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name                           = "http"
      port                           = 80
      protocol                       = "TCP"
      enable_implicit_network_policy = false
    }
    port {
      name                           = "https"
      port                           = 443
      protocol                       = "TCP"
      enable_implicit_network_policy = true
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "some value"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, IPv4IPFamilies))
}

func testComputeWorkloadConfigContainerRemoveEnvVar(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, IPv4IPFamilies))
}

func testComputeWorkloadConfigContainerAddProbes(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    liveness_probe {
      period_seconds = 60
      success_threshold = 1
      failure_threshold = 4
      initial_delay_seconds = 60
      http_get {
        port = 80
        path = "/"
        scheme = "HTTP"
        http_headers = {
          "content-type" = "application/json"
        }
      }
    }
    readiness_probe {
      period_seconds = 60
      success_threshold = 1
      failure_threshold = 4
      initial_delay_seconds = 60
      timeout_seconds = 75
      tcp_socket {
        port = 80
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, IPv4IPFamilies))
}

func testComputeWorkloadConfigContainerImagePullCredentials(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  image_pull_credentials {
    docker_registry {
      server   = "docker.io"
      username = "my-registry-user"
      password = "secret registry password"
      email    = "developers@stackpath.com"
    }
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, IPv4IPFamilies))
}

func testComputeWorkloadConfigContainerAddContainer(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  container {
    name  = "app-2"
    image = "nginx:v1.15.0"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, IPv4IPFamilies))
}

func testComputeWorkloadConfigVirtualMachineBasic(suffix string, enableNAT *bool, ipFamilies []string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "bar" {
  name = "My Terraform Compute VM Workload - %s"
  slug = "terraform-vm-workload-%s"
  %s

  virtual_machine {
    name  = "app"
    image = "stackpath-edge/centos-7:v201905012051"
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    user_data = <<EOT
package_update: true
packages:
- nginx
EOT
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, ipFamilies))
}

func testComputeWorkloadConfigContainerAddVolumeMounts(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo-volume" {
  name = "My Compute Workload Volume - %s"
  slug = "my-compute-workload-volume-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    volume_mount {
      slug = "volume"
      mount_path = "/var/log"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }

  volume_claim {
    name = "volume"
    slug = "volume"
    resources {
      requests = {
        "storage" = "10Gi"
      }
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, IPv4IPFamilies))
}

func testComputeWorkloadConfigAutoScalingConfiguration(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  image_pull_credentials {
    docker_registry {
      server   = "docker.io"
      username = "my-registry-user"
      password = "secret registry password"
      email    = "developers@stackpath.com"
    }
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    max_replicas = 4
    scale_settings {
      metrics {
        metric = "cpu"
        average_utilization = 50
      }
    }
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", enableNAT, IPv4IPFamilies))
}
