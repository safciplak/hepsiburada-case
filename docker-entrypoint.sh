#!/bin/sh

echo $GCP_CRED_JSON > ./gcp_cred.json
exec $@