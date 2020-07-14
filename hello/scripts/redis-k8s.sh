#!/bin/bash

#install Redis statefulset in K8S

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install redis bitnami/redis

#get redis password with
#kubectl get secret --namespace default redis -o jsonpath="{.data.redis-password}" | base64 --decode

#set obtained password in deploy/redis.yaml