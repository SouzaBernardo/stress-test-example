FROM gradle:jdk23-graal AS build
WORKDIR /app
COPY . . 
RUN ./gradlew nativeCompile 

FROM scratch

# Copiar binário nativo da etapa anterior
COPY --from=build /app/build/native/nativeCompile/demo /demo

# Definir ponto de entrada
ENTRYPOINT ["/demo"]