# Ponderada 3 - Conexão com cluster HiveMQ
### Essa atividade é uma continuação da ponderada 1, 2 e 3

## Como rodar o código inteiro

1. Crie um arquivo `.env` o coloque em dois diretórios, um deve ser inserido no diretório `exercicios-mod9/ponderada-03`e outro em `exercicios-mod9/ponderada-03/pkg`. As informações arquivo devem seguir a seguinte estrutura:
```
BROKER_ADDR = <endereço-cluster-hivemq>
HIVE_USER = <usuario>
HIVE_PSWD = <senha>
```
Importante: É importante ter criado um cluster no HiveMQ para que funcione.

2. Em seguida, rode o arquivo de simulação de sensores, para isso entre no diretório `ponderada-03/cmd` e execute o comando abaixo:
```
go run simulation.go
```

Se tiver dado certo você verá o seguindo tipo de dado sendo publicado:

```
Published message: {"identifier":1,"latitude":85.46761635662278,"longitude":73.82166621077455,"current_time":"2024-03-02T22:38:43.393336-03:00","gases-values":{"sensor":"MiCS-6814","unit":"ppm","gases-values":{"carbon_monoxide":230.44,"nitrogen_dioxide":1.11,"ethanol":100.6,"hydrogen":465.36,"ammonia":152.33,"methane":3173.74,"propane":8349.14,"iso_butane":7476.35}},"radiation-values":{"sensor":"RXWLIB900","unit":"W/m2","radiation-values":{"radiation":712.8}}}
```

### Teste de integração
Para comprovar o funcionamento foi feito um teste que funciona como um subscriber para receber as mensagens publicadas. Para rodar o teste entre em `ponderada-03/tests` e rode o seguinte comando:

```
go test -v
```

## Video do funcionamento

https://www.loom.com/share/e99380c66cda426ca15441a93290f7fd