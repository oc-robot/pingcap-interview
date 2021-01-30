# PingCAP 面试小作业

实现在 Kubernetes 环境下，当创建 TiDB 集群时候给 TiDB 集群的每一个 Pod 动态注入一个 agent 容器， 这个 agent 容器中运行一个 agent 进程，并且这个 agent 支持设置网络延迟。

## Agent 使用方式

`curl -X GET  pod-ip:2332/latency/20ms`

## Agent 延迟验证方式

ping pod-ip 
```
PING 10.244.1.10 (10.244.1.10) 56(84) bytes of data.
64 bytes from 10.244.1.10: icmp_seq=1 ttl=62 time=20 ms
64 bytes from 10.244.1.10: icmp_seq=2 ttl=62 time=20 ms
```
