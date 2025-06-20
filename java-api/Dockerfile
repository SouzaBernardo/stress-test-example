# Etapa de build
FROM gradle:8.4-jdk21 AS build
WORKDIR /app

# Copia os arquivos de build e dependências primeiro (melhora o cache)
COPY build.gradle* settings.gradle* ./
COPY gradle ./gradle
RUN gradle --no-daemon build || return 0  # tenta build para baixar dependências

# Copia o restante da aplicação
COPY . .

# Compila o projeto
RUN gradle clean build -x test --no-daemon

# Etapa de runtime
FROM eclipse-temurin:21-jre-alpine
WORKDIR /app

# Copia o JAR gerado
COPY --from=build /app/build/libs/*.jar app.jar

# Expõe a porta (ajuste conforme sua aplicação)
EXPOSE 8080

# Comando para iniciar a aplicação
ENTRYPOINT ["java", "-jar", "app.jar"]
