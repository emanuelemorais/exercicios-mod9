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
Publicado: {
    "sensor": "MiCS-6814",
    "latitude": 93.76226478697556,
    "longitude": 16.492146879375806,
    "unit": "ppm",
    "transmission_rate": 0,
    "current_time": "2024-02-18T23:20:39.542474-03:00",
    "values": {
        "carbon_monoxide": 167.4,
        "nitrogen_dioxide": 9.04,
        "ethanol": 38.49,
        "hydrogen": 969.25,
        "ammonia": 439.16,
        "methane": 4834.37,
        "propane": 3767.95,
        "iso_butane": 5399.89
    }
}
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
Esse teste passara por duas funções, a TestRandomValues e a TestReceivingMessage. A primeira rirá verificar os tipos de dados gerados para simualar a latitude e logitude enviadas. Já a segunda funciona como um subscriber que analisa o qos das mensagens, recebimento e taxa de recebimento.

## Video do funcionamento


[Clique aqui](https://www.youtube.com/watch?v=FKGViEZvCag)