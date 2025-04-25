#!/bin/bash

echo "Testando Go (porta 8081)"
wrk -t4 -c100 -d30s http://localhost:8081/hello

echo ""
echo "Testando Kotlin (porta 8082)"
wrk -t4 -c100 -d30s http://localhost:8082/hello