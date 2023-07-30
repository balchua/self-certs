# Create a personal CA

Use Smallstep, [step-ca](https://smallstep.com/docs/step-ca/#features)

## Install - Kubernetes

``` shell
kubectl create ns step-certs
helm repo add smallstep https://smallstep.github.io/helm-charts/
helm repo update
helm upgrade --namespace step-certs --install step-certificates smallstep/step-certificates



Release "step-certificates" does not exist. Installing it now.
NAME: step-certificates
LAST DEPLOYED: Sun Jul 30 10:19:29 2023
NAMESPACE: step-certs
STATUS: deployed
REVISION: 1
NOTES:
Thanks for installing Step CA.

1. Get the PKI and Provisioner secrets running these commands:
   kubectl get -n step-certs -o jsonpath='{.data.password}' secret/step-certificates-ca-password | base64 --decode
   kubectl get -n step-certs -o jsonpath='{.data.password}' secret/step-certificates-provisioner-password | base64 --decode
2. Get the CA URL and the root certificate fingerprint running this command:
   kubectl -n step-certs logs job.batch/step-certificates

3. Delete the configuration job running this command:
   kubectl -n step-certs delete job.batch/step-certificates

```

Sample output

``` shell
Welcome to Step Certificates configuration.

Configuring kubctl with service account...
Cluster "cfc" set.
User "bootstrap" set.
Context "cfc" created.
Switched to context "cfc".

Checking cluster permissions...
Checking for permission to create configmaps in step-certs namespace: yes
Checking for permission to create secrets in step-certs namespace: yes

Initializating the CA...

Generating root certificate... done!
Generating intermediate certificate... done!

✔ Root certificate: /home/step/certs/root_ca.crt
✔ Root private key: /home/step/secrets/root_ca_key
✔ Root fingerprint: 62eea09f574c3610042df0f00211bf60e45adb054ec14c1d09f12b3980b846d2
✔ Intermediate certificate: /home/step/certs/intermediate_ca.crt
✔ Intermediate private key: /home/step/secrets/intermediate_ca_key
✔ Database folder: /home/step/db
✔ Default configuration: /home/step/config/defaults.json
✔ Certificate Authority configuration: /home/step/config/ca.json

Your PKI is ready to go. To generate certificates for individual services see 'step help ca'.

FEEDBACK 😍 🍻
  The step utility is not instrumented for usage statistics. It does not phone
  home. But your feedback is extremely valuable. Any information you can provide
  regarding how you’re using `step` helps. Please send us a sentence or two,
  good or bad at feedback@smallstep.com or join GitHub Discussions
  https://github.com/smallstep/certificates/discussions and our Discord 
  https://u.step.sm/discord.

Creating configmaps and secrets in step-certs namespace ...
configmap/step-certificates-config replaced
configmap/step-certificates-certs replaced
configmap/step-certificates-secrets replaced
secret/step-certificates-ca-password replaced
secret/step-certificates-provisioner-password replaced
configmap/step-certificates-config labeled
configmap/step-certificates-certs labeled
configmap/step-certificates-secrets labeled
secret/step-certificates-ca-password labeled
secret/step-certificates-provisioner-password labeled

Step Certificates installed!

CA URL: https://step-certificates.step-certs.svc.cluster.local
CA Fingerprint: 62eea09f574c3610042df0f00211bf60e45adb054ec14c1d09f12b3980b846d2

```

## Generate certs

### Get root ca

``` shell
step ca bootstrap --ca-url https://10.152.183.226/ --fingerprint 62eea09f574c3610042df0f00211bf60e45adb054ec14c1d09f12b3980b846d2
```

### Generate certs

This must be done inside the ste-ca pod.

``` shell
step ca certificate --not-after 5m localhost srv.crt srv.key --force
```

Copy the certs to your local machine.

## Run you app server

``` shell
go run github.com/balchua/ca-test
```

## Get root ca

``` shell
$ step ca root root.crt

The root certificate has been saved in root.crt.

```
## Inspect the certificate

``` shell
step certificate inspect srv.crt --short
```

### Auto renewal

``` shell
step ca renew --daemon --force srv.crt srv.key
```