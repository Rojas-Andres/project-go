#  Iniciar go

## Tener en cuenta siempre donde esta nuestra carpeta de trabajo

GOPATH 

- Estructura es
    - bin
    - src
        - github.com
            - Rojas-Andres
                - project-go( Repository )
    - pkg
Recordar configurar el GOROOT Y GOPATH 
go mod init github.com/Rojas-Andres/project-go

- El paquete fmt permite enviar desde la consola

## Ejecutar archivos de go

go run main.go


# Compilar en go

- go build main.go
- Esto genera un .exe (main.exe) se podria ejecutar solo en la terminal ejecutando main.exe

## Interfaces
- Son modelos de comportamiento que se pueden asociar a una estructura


## Defer

- cuando termine una funcion ejecuta ese defer


- generalmente se usa el recover con defer