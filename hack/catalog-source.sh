#!/bin/sh

if [[ -z ${1} ]]; then
    CATALOG_NS="openshift-marketplace"
else
    CATALOG_NS=${1}
fi

CSV=`cat  deploy/olm-catalog/silicon-pod-operator/0.0.1/silicon-pod-operator.v0.0.1.clusterserviceversion.yaml | sed -e 's/^/          /' | sed '0,/ /{s/          /        - /}'`
CRD=`cat deploy/crds/app_v1alpha1_siliconpod_crd.yaml  | sed -e 's/^/          /' | sed '0,/ /{s/          /        - /}'`
PKG=`cat deploy/olm-catalog/silicon-pod-operator/silicon-pod-operator.package.yaml | sed -e 's/^/          /' | sed '0,/ /{s/          /        - /}'`

cat << EOF > deploy/olm-catalog/silicon-pod-operator/catalog-source.yaml
apiVersion: v1
kind: List
items:
  - apiVersion: v1
    kind: ConfigMap
    metadata:
      name: silicon-pod-operator-resources
      namespace: ${CATALOG_NS}
    data:
      clusterServiceVersions: |
${CSV}
      customResourceDefinitions: |
${CRD}
      packages: >
${PKG}

  - apiVersion: operators.coreos.com/v1alpha1
    kind: CatalogSource
    metadata:
      name: silicon-pod-operator-resources
      namespace: ${CATALOG_NS}
    spec:
      configMap: silicon-pod-operator-resources
      displayName: Silicon Pod Operators
      publisher: Red Hat
      sourceType: internal
    status:
      configMapReference:
        name: silicon-pod-operator-resources
        namespace: ${CATALOG_NS}
EOF