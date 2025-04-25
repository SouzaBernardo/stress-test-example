# Etapa 1: Build com GraalVM
FROM ghcr.io/graalvm/native-image-community AS build
WORKDIR /app
COPY . .
RUN ./gradlew nativeCompile --no-daemon

# Etapa 2: Imagem mínima com binário nativo
FROM scratch
COPY --from=build /app/build/native/nativeCompile/demo .
EXPOSE 8080
ENTRYPOINT ["./demo"]