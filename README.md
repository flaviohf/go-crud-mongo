# App feita em GO utilizando o framework Gin para aprendizado  

**Para gerar a documentação swagger**  
```swag init -d ./cmd/web,./internal/controllers,./internal/routes,./internal/domains```

**Para subir a app**  
```go run .\cmd\web\main.go```

**Para gerar (ou derrubar) o container docker**  
```docker compose up -d --build```  
```docker compose down```
