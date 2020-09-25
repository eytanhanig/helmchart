package helmtest

import (
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestDashImageTag(t *testing.T) {

	helmChartPath := "../pachyderm"

	options := &helm.Options{
		SetValues: map[string]string{"dash.image.tag": "0.5.5"},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "deployment", []string{"templates/dash/deployment.yaml"})

	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)

	expectedContainerImage := "pachyderm/dash:0.5.5"
	deploymentContainers := deployment.Spec.Template.Spec.Containers
	if deploymentContainers[0].Image != expectedContainerImage {
		t.Fatalf("Rendered container image (%s) is not expected (%s)", deploymentContainers[0].Image, expectedContainerImage)
	}
}

func TestEtcdImageTag(t *testing.T) {

	helmChartPath := "../pachyderm"

	options := &helm.Options{
		SetValues: map[string]string{"etcd.image.tag": "blah"},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "statefulset", []string{"templates/etcd/statefulset.yaml"})

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(t, output, &statefulSet)

	expectedContainerImage := "pachyderm/etcd:blah"
	deploymentContainers := statefulSet.Spec.Template.Spec.Containers
	if deploymentContainers[0].Image != expectedContainerImage {
		t.Fatalf("Rendered container image (%s) is not expected (%s)", deploymentContainers[0].Image, expectedContainerImage)
	}
}

func TestPachdImageTag(t *testing.T) {
	helmChartPath := "../pachyderm"

	options := &helm.Options{
		SetValues: map[string]string{"pachd.image.tag": "1234"},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "deployment", []string{"templates/pachd/deployment.yaml"})

	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)

	expectedContainerImage := "pachyderm/pachd:1234"
	deploymentContainers := deployment.Spec.Template.Spec.Containers
	if deploymentContainers[0].Image != expectedContainerImage {
		t.Fatalf("Rendered container image (%s) is not expected (%s)", deploymentContainers[0].Image, expectedContainerImage)
	}
}

func TestSetNamespaceWorkerRoleBinding(t *testing.T) {

	helmChartPath := "../pachyderm"

	expectedNamespace := "kspace"

	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", expectedNamespace),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "rolebinding", []string{"templates/pachd/rbac/worker-rolebinding.yaml"})

	var rolebinding rbacv1.RoleBinding
	helm.UnmarshalK8SYaml(t, output, &rolebinding)

	subjects := rolebinding.Subjects
	if subjects[0].Namespace != expectedNamespace {
		t.Fatalf("Rendered namespace (%s) is not expected (%s)", subjects[0].Namespace, expectedNamespace)
	}

}

func TestSetNamespaceServiceAccount(t *testing.T) {

	helmChartPath := "../pachyderm"

	expectedNamespace := "kspace1234"

	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", expectedNamespace),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "serviceaccount", []string{"templates/pachd/rbac/serviceaccount.yaml"})

	var serviceAccount v1.ServiceAccount
	helm.UnmarshalK8SYaml(t, output, &serviceAccount)

	metadata := serviceAccount.ObjectMeta

	if metadata.Namespace != expectedNamespace {
		t.Fatalf("Rendered namespace (%s) is not expected (%s)", metadata.Namespace, expectedNamespace)
	}
}

func TestSetNamespaceWorkerServiceAccount(t *testing.T) {

	helmChartPath := "../pachyderm"

	expectedNamespace := "kspaceworker"

	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", expectedNamespace),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "workerserviceaccount", []string{"templates/pachd/rbac/worker-serviceaccount.yaml"})

	var serviceAccount v1.ServiceAccount
	helm.UnmarshalK8SYaml(t, output, &serviceAccount)

	metadata := serviceAccount.ObjectMeta

	if metadata.Namespace != expectedNamespace {
		t.Fatalf("Rendered namespace (%s) is not expected (%s)", metadata.Namespace, expectedNamespace)
	}
}

func TestSetNamespaceClusterRoleBinding(t *testing.T) {

	helmChartPath := "../pachyderm"

	expectedNamespace := "kspace123"

	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", expectedNamespace),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, "clusterrolebinding", []string{"templates/pachd/rbac/clusterrolebinding.yaml"})

	var rolebinding rbacv1.ClusterRoleBinding
	helm.UnmarshalK8SYaml(t, output, &rolebinding)

	subjects := rolebinding.Subjects
	if subjects[0].Namespace != expectedNamespace {
		t.Fatalf("Rendered namespace (%s) is not expected (%s)", subjects[0].Namespace, expectedNamespace)
	}

}