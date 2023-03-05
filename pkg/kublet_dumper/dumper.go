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

func Dump() {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}
	//create a pod in the cluster
	pod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kubelet-dumper",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "kubelet-dumper",
					Image: "busybox",
					Command: []string{
						"whoami",
					},
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("Created Pod %q.\n", pod.GetObjectMeta().GetName())

	// Wait for the Pod to be in the "Running" state
	err = WaitForPodRunning(clientset, "default", "kubelet-dumper", 5*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	podLogOpts := &v1.PodLogOptions{}
	podLogs, err := clientset.CoreV1().Pods("default").GetLogs("kubelet-dumper", podLogOpts).Stream(context.Background())
	if err != nil {
		log.Fatal(err)
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
			time.Sleep(5 * time.Second)
		}
	}
}
