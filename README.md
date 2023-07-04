# Gerenciador de tarefas

Um simples crud com gerenciador de tarefas

### Pré requisitos para rodar o projeto:
- Docker instalado na sua máquina

### Como rodar o projeto?
```bash
make run
```

### Rotas
- [GET] http://localhost:8081/tasks
Lista todas as tarefas salvas no json

```bash
curl -v -X GET 'http://localhost:8081/tasks' 
```

- [POST] http://localhost:8081/tasks/create
Adiciona tarefas novas ao json

```bash
curl -X POST \
  'http://localhost:8081/tasks/create' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "description": "lavar calça branca",
  "done": false
}'
```

- [PUT] http://localhost:8081/tasks/done/:id
Altera o status de uma tarefa para done

```bash
curl -X PUT 'http://localhost:8081/tasks/done/1' 
```


- [DELETE] http://localhost:8081/tasks/delete/all
Limpa todas as tarefas

```bash
curl -X DELETE 'http://localhost:8081/tasks/delete/all' 
```