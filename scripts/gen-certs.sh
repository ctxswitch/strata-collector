#!/bin/bash

openssl genrsa -out ca.key 2048

openssl req -new -x509 -days 365 -key ca.key \
  -subj "/C=US/CN=strata-ctx-sh-webhook"\
  -out ca.crt

openssl req -newkey rsa:2048 -nodes -keyout server.key \
  -subj "/C=US/CN=strata-ctx-sh-webhook" \
  -out server.csr

openssl x509 -req \
  -extfile <(printf "subjectAltName=DNS:strata-ctx-sh-webhook.strata-collector.svc") \
  -days 365 \
  -in server.csr \
  -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out server.crt

cat ca.crt | base64 | fold > cabundle.crt

cat > config/overlays/dev/webhook.yaml << EOF
---
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  name: mutating-webhook-configuration
webhooks:
- name: mstrata.ctx.sh
  clientConfig:
    caBundle: "$(awk '{printf "%s\\n", $0}' cabundle.crt)"
    service:
      name: strata-ctx-sh-webhook
      namespace: strata-collector
      port: 9443
---
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: validating-webhook-configuration
webhooks:
- name: vstrata.ctx.sh
  clientConfig:
    caBundle: "$(awk '{printf "%s\\n", $0}' cabundle.crt)"
    service:
      name: strata-ctx-sh-webhook
      namespace: strata-collector
      port: 9443
EOF

mv server.crt ./config/overlays/dev/tls.crt
mv server.key ./config/overlays/dev/tls.key

rm ca.crt ca.key ca.srl server.csr cabundle.crt
