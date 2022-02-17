# Golang Hello World

This is a proof-of-concept golang custom runtime for Kyma Srverless functions.

### Usage

- Build and push the image to your registry
```bash
docker build . -t melsayed/gohello:latest
docker push melsayed/gohello:latest
```

- Create a custom runtime function on Kyma, using the image override field `customRuntimeImage`. This will disable Kyma build process and run the image directly as a function workload.
```bash
cat <<EOF | kubectl apply -f -
---
apiVersion: serverless.kyma-project.io/v1alpha1
kind: Function
metadata:
  creationTimestamp: null
  name: gohello
  namespace: default
spec:
  customRuntimeImage: melsayed/gohello:latest
  runtime: custom
  source: "-"
EOF
```

**Note:** Since using custom runtimes completely bypasses the Function build process, it's not possible to use custom runtimes with git functions.