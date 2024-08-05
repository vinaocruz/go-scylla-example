## Aplicativo exemplo com Go + Scylla

Este é um exemplo de aplicativo utilizando Golang e ScyllaDB para persistência de dados.

### Setup

Para executar localmente, suba o banco através do docker-compose:

```
docker-compose up
```

Em seguida, acesse cqlsh para criar keyspace e a tabela:

```
docker exec -it scylla cqlsh
```

Execute os comandos abaixo:

```
cqlsh> CREATE KEYSPACE my_company WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': 1} ;

cqlsh> CREATE TABLE my_company.drivers (id uuid, cnh text, license_plate text, name text, model text, createdat timestamp, PRIMARY KEY (cnh, license_plate));
```

Para este momento, não há persistência do volume do banco. Assim, caso o container seja apagado, todos os dados eventualmente salvos também serão perdidos.

Para rodar a API, execute:
```
go run main.go
```

### Documentação

Este exemplo permite que seja definido motoristas e seus carros. Foi definido uma modelagem para o banco Scylla desnomalizado, visando sua otimização. Assim, para verificar os endpoints disponíveis nessa aplicação, acesse:

http://localhost:8080/docs/index.html

A documentação foi gerada através do recurso [Swaggo](https://github.com/swaggo/swag)

### TODO

Devido ao tempo, mantive como objetivo para evoluir este exemplo:

* Adicionar migrate
* Persistir volume do banco de dados
* Validação de dados de entrada
* Rodar aplicação com docker-compose
* Implementar testes com K6 para avaliar desempenho
* Implementar testes unitários