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

# Variáveis de ambiente
- **WEB_SERVER_PORT** : indica porta a qual o servidor será iniciado ( padrão: 8000 )
- **REDIS_HOST** : Endereço host do red ( padrão: redis )
- **REDIS_PORT** : Porta utilizada pelo redis ( padrão: 6379 )
- **REDIS_DB** : Database redis ultilizado ( padrão: 0 )
- **REDIS_PASSWORD** : Senha para conexão do redis ( padrão: "" )
- **REDIS_RATE_LIMIT_TTL** : Tempo de retenção do rate limit ( padrão: 10 )
- **REQUEST_BLOCKING_TIME_IP** : Tempo de bloqueio por IP ( padrão: 60 )
- **REQUEST_BLOCKING_TIME_TOKEN** : Tempo de bloqueio por TOKEN ( padrão: 60 )
- **MAX_REQUESTS_PER_IP_PER_SECOND** : Quantidade máxima de requisições por IP ( padrão: 10 )
- **MAX_REQUESTS_PER_TOKEN_PER_SECOND** : Quantidade máxima de requisições por Token ( padrão: 100 )
- **RATE_LIMITER_IP_ENABLED** : Habilita rate limiter por IP ( padrão: true )
- **RATE_LIMITER_TOKEN_ENABLED** : Habilita rate limiter por TOken ( padrão: true )

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
docker compose up -d redis app
```

Para acessar o conteúdo da API basta acessar a url `http://localhost:8000/v1/status` 

Para executar chamadas usando restClient basta usar o [arquivo](./api/status.http) onde pode realizar chamadas Get com e sem token

# Executando teste de carga
Para executar o teste de carga, basta iniciar o K6 pelo comando no docker
K6 é uma ferramenta de teste de carga open source, para mais informações acessar o [site](https://k6.io/)
```
docker compose up k6
```

# Alterando implementação Redis
Para facilitar a alteração da implementação do ratelimit a seguinte [interface](./internal/entity/limiter.go) deve ser respeitada
```
type Limiter interface {
	IsRateLimitExceeded(ip string, token string) bool
	RegisterAccess(ip string, token string) error
	BlockAccess(ip string, token string) error
}
```

A implementação com redis foi feita no serviço [redisRateLimiter](./internal/service/redis_limiter.go)