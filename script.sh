#!/bin/bash

# Iniciar containers
docker-compose up -d

# Aguardar APIs estarem disponíveis
for i in {1..30}; do
  if curl -s http://localhost:8081/hello && curl -s http://localhost:8082/hello; then
    echo "APIs estão prontas."
    break
  fi
  echo "Aguardando APIs ficarem prontas..."
  sleep 2
done

# Executar testes com wrk
echo "Testando API Go (porta 8081)"
wrk -t4 -c100 -d30s http://localhost:8081/hello

echo "Testando API Kotlin (porta 8082)"
wrk -t4 -c100 -d30s http://localhost:8082/hello

# Capturar logs das APIs
echo "Logs da API Go:"
curl http://localhost:8081/hello

echo "Logs da API Kotlin:"
curl http://localhost:8082/hello

# Parar containers
docker-compose down