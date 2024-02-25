#!/bin/sh

DEPLOYMENT="go-app"
NAMESPACE="default"

BLUE=$(kubectl get deployments $DEPLOYMENT-blue -n $NAMESPACE -o jsonpath='{@.spec.selector.matchLabels.slot}')
GREEN=$(kubectl get deployments $DEPLOYMENT-green -n $NAMESPACE -o jsonpath='{@.spec.selector.matchLabels.slot}')

swap(){
    env_prd=$1
    env_stage=$2
    kubectl patch deployments $DEPLOYMENT-$env_prd --patch  "{\"spec\": {\"selector\": {\"matchLabels\": {\"slot\": \"$env_stage\"}}}}"
}

if [ $BLUE == 'prd' ]; then
    echo "Production Slot is BLUE"
    swap green blue
elif [ $GREEN == 'prd' ]; then
    echo "Production Slot is GREEN"
    swap blue green
else
    echo "*** OPS... ***"
fi