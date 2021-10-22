package controllers

import (
	httpv1alpha1 "github.com/raihankhan/httpApiServer-controller-kubebuilder/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func newDeployment(apiServer httpv1alpha1.Apiserver) *appsv1.Deployment {

	labels := map[string]string{
		"httpv1alpha1/apiserver": apiServer.ObjectMeta.Name,
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiServer.Spec.Name + "-depl",
			Namespace: apiServer.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: apiServer.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "apiserver",
							Image: "docker.io/raihankhanraka/ecommerce-api:v1.1",
							Ports: []corev1.ContainerPort{
								{
									Name:          "ecom",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}
}

func newService(apiServer httpv1alpha1.Apiserver) *corev1.Service {

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiServer.Spec.Name + "-np",
			Namespace: apiServer.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.IntOrString{IntVal: 8080},
					Port:       8080,
				},
			},
			Selector: map[string]string{
				"httpv1alpha1/apiserver": apiServer.ObjectMeta.Name,
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
}
