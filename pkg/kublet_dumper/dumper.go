package kubelet_dumper

import (
	"context"
	"fmt"
	"log"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Dump(nodename string) {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}

	hostPathVolume := v1.Volume{
		Name: "kubelet",
		VolumeSource: v1.VolumeSource{
			HostPath: &v1.HostPathVolumeSource{
				Path: "/var/lib/kubelet",
			},
		},
	}

	hostPathVolumeMount := v1.VolumeMount{
		Name:      "kubelet",
		MountPath: "/var/lib/kubelet",
	}
	podName := "kubelet-dumper-" + nodename
	//create a pod in the cluster
	pod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:         podName,
					Image:        "ubuntu:22.04",
					VolumeMounts: []v1.VolumeMount{hostPathVolumeMount},
					Command: []string{
						"/bin/bash", "-c",
					},
					Args: []string{
						//"ps -fC kubelet && cat /var/lib/kubelet/config.yaml",
						"cat /var/lib/kubelet/config.yaml",
					},
				},
			},
			HostPID:       true,
			Volumes:       []v1.Volume{hostPathVolume},
			RestartPolicy: v1.RestartPolicyNever,
			NodeName:      nodename,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("Created Pod %q.\n", pod.GetObjectMeta().GetName())

	// Wait for the Pod to be in the "Running" state
	err = WaitForPodRunning(clientset, "default", podName, 5*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	podLogOpts := &v1.PodLogOptions{}
	podLogs, err := clientset.CoreV1().Pods("default").GetLogs(podName, podLogOpts).Stream(context.Background())
	if err != nil {
		log.Print(err)
	}
	defer podLogs.Close()

	// Read the logs from the stream and output them
	buf := make([]byte, 1024)
	for {
		n, err := podLogs.Read(buf)
		if n == 0 && err != nil {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(buf[:n]))
	}
	//Delete the pod
	err = clientset.CoreV1().Pods("default").Delete(context.Background(), podName, metav1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

// WaitForPodRunning polls the Pod's status until it is in the "Running" state or the timeout is exceeded
func WaitForPodRunning(clientset *kubernetes.Clientset, namespace string, podName string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for Pod %q to be Running", podName)
		default:
			pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
			if err != nil {
				return err
			}
			if pod.Status.Phase == v1.PodSucceeded {
				return nil
			}
			if pod.Status.Phase == v1.PodFailed {
				return fmt.Errorf("pod %q failed", podName)
			}
			time.Sleep(5 * time.Second)
		}
	}
}
