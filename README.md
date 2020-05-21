# chainer
## Introduction
Chainer generates a list of service chains with the following configuration:

| Property          | Value        |
|:----------------- |:------------:|
| Length            | \[4, 7\)     |
| Per Instance Cost | 100\$        |
| Bandwidth         | 250 bps      |

Please note that customers pay the instance cost so it would be profit in the eye of datacenter owner.

Following VNF types are availabe to customers for placing in their service chains.

```yml
---
types:
  - name: ingress
    cores: 0
    ram: 0
    ingress: true
    manageable: false
  - name: egress
    cores: 0
    ram: 0
    egress: true
    manageable: false
  - name: vFW
    cores: 2
    ram: 2
    manageable: true
  - name: vNAT
    cores: 2
    ram: 4
    manageable: true
  - name: vIDS
    cores: 2
    ram: 2
    manageable: true
  - name: vDPI
    cores: 2
    ram: 4
    manageable: true
```
