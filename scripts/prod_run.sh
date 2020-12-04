#!/bin/bash

echo "############### Sourcing Prod Configs ###############"
set -a
if ! . env/prod.env; then
  set +a
  echo "Failed to parse configs. Exiting..."
  exit 1
fi
set +a
echo "Configs sourced."
echo "#####################################################"

echo "########### Building Application Binary #############"
if ! CGO_ENABLED=0 GOOS=linux go build -o bin/application; then
  echo "Failed to build application binary."
  exit 1
fi
echo "Application binary built."
echo "#####################################################"

echo "############ Running Application Binary #############"
if ! bin/application; then
  echo "Application exited with non-zero status code."
  exit 1
fi
echo "#####################################################"