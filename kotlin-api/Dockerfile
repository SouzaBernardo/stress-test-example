# Etapa 1: Build nativo com GraalVM
FROM gradle:jdk-21-and-22-graal AS build
WORKDIR /app
COPY . .
RUN ./gradlew nativeCompile --no-daemon

# Etapa 2: Imagem mínima com o binário nativo
FROM scratch
COPY --from=build /app/build/native/nativeCompile/demo .
ENTRYPOINT ["./demo"]