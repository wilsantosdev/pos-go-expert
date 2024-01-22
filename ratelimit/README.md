# Rate Limiter

A proposta desse projeto é desenvolver uma mecânica de rate limit com baixo acoplamento de infra, onde seja possível, reutilizar a biblioteca criada assim como substituir o uso repositório de registros ( inicialmete projetado para usar REDIS )

# API
Com a necessidade apenas de servir como base de teste para o rate limit a API possui apenas 1 endpoints 

- v1/status
    - Response 200
        >code: **200** <br/>
         mensagem **OK**
    - Response 429 
        >code: **200** <br/>
        mensagem **OK**

# Configs
Para que seja possível configurar os parametros do ratelimit algumas variáveis de ambiente foram adicionadas

- REDIS_RATE_LIMIT_TTL
    - Tempo de permanência do dado no **REDIS**
- REQUEST_BLOCKING_TIME_IP
    - Tempo de bloqueio quando baseado no **IP**
- REQUEST_BLOCKING_TIME_TOKEN
    - Tempo de bloqueio quando baseado no **TOKEN**
- RATE_LIMITER_IP_ENABLED
    - Habilita rate limit por **IP**
- RATE_LIMITER_TOKEN_ENABLED
    - Habilita rate limit por **TOKEN**

# Executando aplicação
Para executar a aplicação utilize o comando a baixo: 
```
docker compose -f deployments/local/docker-compose.yml up -d
```
