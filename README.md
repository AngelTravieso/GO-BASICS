### GOLANG:

### Ejecutar archivo Go

Basta con ir a la terminal y ejecutar:

```
go run [archivo].go
```

ejemplo:

```
go run app.go
```

### Módulos en Go

Los módulos se crean para poder hacer un build y crear un archivo ejecutable `(.exe)`, que pueda correr en cualquier PC

## Crear un módulo en Go:

```
go mod init [name]
go mod init example.com/first-app
```

Esto creará un archivo `go.mod` con el nombre del módulo, luego con el comando `go build` se creará el archivo ejecutable con el nombre que se le dio al módulo anteriormente, es decir, tendremos un archivo `first-app.exe`
