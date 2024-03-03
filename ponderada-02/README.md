# Ponderada 2 - Simulador de dispositivos IoT 
### Essa atividade é uma continuação da ponderada 1

## Como rodar o código inteiro

1. Inicie o broker local, para isso entre no diretório `ponderada-01/config` e rode o seguinte comando:
```
mosquitto -c mosquitto.conf
```

2. Em seguida, rode o arquivo de simulação de sensores, para isso entre no diretório `ponderada-01/cmd` e execute o comando abaixo:
```
go run simulation.go
```

Se tiver dado certo você verá o seguindo tipo de dado sendo publicado:

```
Published message: {"identifier":1,"latitude":85.46761635662278,"longitude":73.82166621077455,"current_time":"2024-03-02T22:38:43.393336-03:00","gases-values":{"sensor":"MiCS-6814","unit":"ppm","gases-values":{"carbon_monoxide":230.44,"nitrogen_dioxide":1.11,"ethanol":100.6,"hydrogen":465.36,"ammonia":152.33,"methane":3173.74,"propane":8349.14,"iso_butane":7476.35}},"radiation-values":{"sensor":"RXWLIB900","unit":"W/m2","radiation-values":{"radiation":712.8}}}
```

## Como rodar os testes
### Teste da geraçao de dados - mics6814

1. Entre no seguinte diretório `ponderada-01/internal/mics6814` e rode o comando abaixo:
```
go test -v
```
Esse teste irá verificar se os dados criados estão no range correto sem aterações de valores


### Teste da geraçao de dados - rxwlib900

1. Entre no seguinte diretório `ponderada-01/internal/rxwlib900` e rode o comando abaixo:
```
go test -v
```
Esse teste irá verificar se os dados criados estão no range correto sem aterações de valores


### Teste de conexão com o broker

1. Entre no seguinte diretório `ponderada-01/pkg/common` e rode o comando abaixo:
```
go test -v
```
Esse teste irá verificar se a conexão com o broker é feito corretamente

### Teste de geração de valores, de QOS, recebimento das mensagens e taxa de transmissão dos dados

1. Entre no seguinte diretório `ponderada-01/pkg/controller` e rode o comando abaixo:
```
go test -v
```

Este teste atua como um subscriber, analisando o QoS das mensagens recebidas, validando os dados conforme o padrão esperado e confirmando o recebimento e a taxa de disparo das mensagens.

## Video do funcionamento

https://www.loom.com/share/f57d52daae1d4c8288ba1fe6badf384e