Silicon Pod
====================
Sample operator to expose metrics in openshift, developed using operator-sdk .
Operator SDK will create servicemonitor for exposing metrics endpoint


To build binaries, run

```
make

````

For the monitoring to work the operator need to be deployed under namespaec openshift-* with label as show below

```
apiVersion: v1
kind: Namespace
metadata:
  name: openshift-silicon-pod
  annotations:
    openshift.io/node-selector: ""
  labels:
    openshift.io/cluster-monitoring: "true"
    name: openshift-silicon-pod
    network.openshift.io/policy-group: monitoring
```

And create role and role binding to give permission for promethues to access your pod .

Role Binding
====================

```
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  # TODO this should be a clusterrole
  name: prometheus-k8s
  namespace: openshift-silicon-pod
rules:
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  - pods
  verbs:
  - get
  - list
  - watch

```

Role Binding
====================

``` 
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: prometheus-k8s
  namespace: openshift-silicon-pod
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: prometheus-k8s
subjects:
- kind: ServiceAccount
  name: prometheus-k8s
  namespace: openshift-monitoring  #important

``` 