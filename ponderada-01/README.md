# Ponderada 1 - Simulador de dispositivos IoT 
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
### Teste da geraçao de dados
- Entre no diretório `ponderada-01/internal/mics` e rode o comando abaixo:
```
go test -v
```

Caso o teste tenha dado certo, resultado será como o bloco abaixo:
```
=== RUN   TestCreateGasesValues
--- PASS: TestCreateGasesValues (0.00s)
PASS
ok      github.com/emanuelemorais/exercicios-mod9/ponderada-01/internal/mics    0.489s
```

### Teste do controller para envio MQTT
- Entre no diretório `ponderada-01/pgk/controller`e rode o comando abaixo:
```
go test -v
```
Caso o teste tenha dado certo, resultado será como o bloco abaixo:
```
=== RUN   TestConnectBroker
--- PASS: TestConnectBroker (0.00s)
=== RUN   TestRandomValues
--- PASS: TestRandomValues (0.00s)
=== RUN   TestCotroller
Publicado: {
    "sensor": "MiCS-6814",
    "latitude": 79.24691802739905,
    "longitude": 8.137942016438766,
    "unit": "ppm",
    "transmission_rate": 0,
    "current_time": "2024-02-18T23:19:59.99274-03:00",
    "values": {
        "carbon_monoxide": 896.8,
        "nitrogen_dioxide": 6.14,
        "ethanol": 164.08,
        "hydrogen": 26.6,
        "ammonia": 419.38,
        "methane": 5939.32,
        "propane": 3267.38,
        "iso_butane": 6185.97
    }
}
Publicado: {
    "sensor": "MiCS-6814",
    "latitude": 79.24691802739905,
    "longitude": 8.137942016438766,
    "unit": "ppm",
    "transmission_rate": 0,
    "current_time": "2024-02-18T23:20:00.994682-03:00",
    "values": {
        "carbon_monoxide": 748.67,
        "nitrogen_dioxide": 2.71,
        "ethanol": 81.13,
        "hydrogen": 53.29,
        "ammonia": 104.65,
        "methane": 8080.66,
        "propane": 6928.85,
        "iso_butane": 3139.45
    }
}
    controller_test.go:61: Recebido: {
            "sensor": "MiCS-6814",
            "latitude": 79.24691802739905,
            "longitude": 8.137942016438766,
            "unit": "ppm",
            "transmission_rate": 0,
            "current_time": "2024-02-18T23:19:59.99274-03:00",
            "values": {
                "carbon_monoxide": 896.8,
                "nitrogen_dioxide": 6.14,
                "ethanol": 164.08,
                "hydrogen": 26.6,
                "ammonia": 419.38,
                "methane": 5939.32,
                "propane": 3267.38,
                "iso_butane": 6185.97
            }
        } do tópico: mics6814
        
    controller_test.go:61: Recebido: {
            "sensor": "MiCS-6814",
            "latitude": 79.24691802739905,
            "longitude": 8.137942016438766,
            "unit": "ppm",
            "transmission_rate": 0,
            "current_time": "2024-02-18T23:20:00.994682-03:00",
            "values": {
                "carbon_monoxide": 748.67,
                "nitrogen_dioxide": 2.71,
                "ethanol": 81.13,
                "hydrogen": 53.29,
                "ammonia": 104.65,
                "methane": 8080.66,
                "propane": 6928.85,
                "iso_butane": 3139.45
            }
        } do tópico: mics6814
        
--- PASS: TestCotroller (2.00s)
PASS
ok      github.com/emanuelemorais/exercicios-mod9/ponderada-01/pkg/controller   2.460s
```



## Como funciona o cógigo

O sensor simulado nesta atividade é o MiCS-6814, que mede diversos aspectos atmosféricos. Para simular as informações captadas por ele, foi criada uma função que randomiza valores de acordo com o intervalo que o sensor é capaz de detectar. Além disso, foi desenvolvido um código de teste que verifica se os valores gerados estão dentro do ideal. O código relacionado aos dados do sensor pode ser encontrado em `ponderada-01/internal/mics`.

Adicionalmente, foi criado um controller responsável pelo envio das informações via MQTT. Ele utiliza os dados gerados anteriormente para enviar informações dinâmicas para o broker. Também foram realizados testes relacionados a esse código, testando a conexão ao broker, a geração de valores randômicos para latitude e longitude, e, por fim, um subscriber para assegurar que as informações estão sendo transmitidas corretamente. Esse código pode ser encontrado em `ponderada-01/pkg/controller`.

Finalmente, foi realizada uma simulação de vários sensores, que pode ser encontrada em `ponderada-01/cmd`. O código cria 5 threads diferentes para representar 5 sensores diferentes publicando informações.


## Video do funcionamento


[Clique aqui](https://www.youtube.com/watch?v=v_4snGECx5s)