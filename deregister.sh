#!/bin/bash

SERVICE_NAME=cloud-connector



curl -k https://consul.dev-lamassu.zpd.ikerlan.es/v1/agent/health/service/name/${SERVICE_NAME} | jq .[].Service.ID -r -c | while read item; do
    # do stuff with $item
    echo $item
    curl -k -X PUT https://consul.dev-lamassu.zpd.ikerlan.es/v1/agent/service/deregister/$item
done
