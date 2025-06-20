#!/bin/bash

echo "Testando API Go (porta 8080)"
wrk -t4 -c100 -d30s http://localhost:8080/payments -H "Status: success"


echo "Testando API Java (porta 8081)"
wrk -t4 -c100 -d30s http://localhost:8081/users