# Kubelet Dumper

A very simple tool to dump the kubelet configuration from one node in a cluster, or all nodes in a cluster.

Connection is based on your currently configured kubectl context.

At the moment this is just a PoC which assumes a hard coded path to the kubelet configuration file of `/var/lib/kubelet/config.yaml`. Ideally we'd change this to get it from the Process list on the node.

## Usage

### Dump a single node

```
kubelet-dumper dump <node>
```

### Dump all nodes

```
kubelet-dumper dumpAll
```