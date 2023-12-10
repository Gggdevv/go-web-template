#!/bin/bash

cd ./scripts/docker && docker compose -p go-web-template up  -d  --force-recreate && cd ../..