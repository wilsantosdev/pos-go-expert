# Stress tester
Objetivo desse projeto é conseguir realizar chamadas para uma URL e verificar o comportamento em determinados cenários

- Chamadas a uma **URL**
- Definir o **total de chamadas** que será feito
- Definir o **número de chamdas simultâneas** 

para isso alguns parametros foram criados para executar o projeto

# Como utilizar 
```
docker compose run --rm app stress --url https://google.com --concurrency 1 --requests 1
```
Executar o comando acima, especificando os seguintes parametros:
- **stress**: Esse é o comando que executa a rotina de testes
- **--url**: Esse argumento obrigatório diz qual o endereço que receberá a requisição
- **--concurrency**: Esse argumento obrigatório define quantas chamadas simultâneas serão feitas
- **--requests**: Essse argumento obrigatório define quantas requisições serão feitas no total