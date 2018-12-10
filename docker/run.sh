#!/bin/bash
echo "********************************************************"
echo "Running job-service"
echo "********************************************************"

echo "********************************************************"
echo "Waiting for the Consul to start on port $CONSUL_PORT"
echo "********************************************************"
while ! `nc -z consul $CONSUL_PORT `; do sleep 3; done
echo ">>>>>>>>>>>> Consul has started"

./job-service
