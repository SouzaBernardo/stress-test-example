# Etapa 1: Construção com GraalVM
FROM ghcr.io/graalvm/jdk:21 AS build

# Instalar dependências necessárias
RUN microdnf install -y findutils

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de construção
COPY build.gradle.kts settings.gradle.kts gradlew ./
COPY gradle gradle/
RUN ./gradlew --no-daemon build

# Copiar código-fonte
COPY src src/

# Compilar imagem nativa
RUN ./gradlew nativeCompile --no-daemon

# Etapa 2: Imagem mínima com a aplicação nativa
FROM scratch

# Copiar binário nativo da etapa anterior
COPY --from=build /app/build/native/nativeCompile/demo /demo

# Definir ponto de entrada
ENTRYPOINT ["/demo"]