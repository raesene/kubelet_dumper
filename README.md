# Kubelet Dumper

A very simple tool to dump the kubelet configuration from one node in a cluster, or all nodes in a cluster.

Connection is based on your currently configured kubectl context.

The approach used at the moment is to connect to the `configz` endpoint on the Kubelet API, via the API server proxy. You'll need appropriate rights at the cluster level for this (`GET` on `node/proxy`). This should be relatively portable.


## Usage

### Dump a single node

```
kubelet-dumper dump <node>
```

### Dump all nodes

```
kubelet-dumper dumpAll
```
