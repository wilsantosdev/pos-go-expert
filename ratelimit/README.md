# Rate Limiter

A proposta desse projeto é desenvolver uma mecânica de rate limit com baixo acoplamento de infra, onde seja possível, reutilizar a biblioteca criada assim como substituir o uso repositório de registros ( inicialmete projetado para usar REDIS )

# API
Com a necessidade apenas de servir como base de teste para o rate limit a API possui apenas 1 endpoints 

- v1/status
    - Response 200
        >code: **200** <br>
         mensagem **OK**
    - Response 429 
        >code: **200** <br>
        mensagem **OK**

