package kubelet_dumper

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getNodes() []string {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	var nodeList []string
	for _, node := range nodes.Items {
		nodeList = append(nodeList, node.Name)
	}
	return nodeList
}

func DumpAll() {
	nodes := getNodes()
	for _, node := range nodes {
		Dump(node)
	}
}
