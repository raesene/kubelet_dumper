# Kubelet Dumper

A very simple tool to dump the kubelet configuration from one node in a cluster, or all nodes in a cluster.

Connection is based on your currently configured kubectl context.

## Usage

### Dump a single node

```
kubelet-dumper dump <node>
```

### Dump all nodes

```
kubelet-dumper dumpAll
```