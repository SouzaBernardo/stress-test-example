# Etapa 1: build nativo
FROM ghcr.io/graalvm/native-image-community:latest AS build
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y findutils && ./gradlew nativeCompile

# Etapa 2: imagem mínima
FROM scratch
COPY --from=build /app/build/native/nativeCompile/demo .
ENTRYPOINT ["./demo"]