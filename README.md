# fundamentos-testing-tarea1
Repositorio de Código para el ramo Fundamentos de Testing y Aseguramiento de la Calidad 2024-1

El código fué originalmente ejecutado y desarrollado en una máquina linux (Fedora 39) amd/64 con go 1.21.8. Además la librería para Unit testing (testing) es parte de la librería estándar de Go.


## Instrucciones de Ejecución

El primer paso será instalar el lenguaje de programación Go, seguir los pasos de [este link](https://go.dev/doc/install).

Una vez esté instalado Go, se debe compilar y ejecutar el programa ohce.go. Go ofrece un comando que combina la compilación y ejecución del programa.

```
$ go run ohce.go <name>
```

## Instrucciones de Testeo

Para correr los tests unitarios dentro de `/utils/utils_test.go` se debe ejecutar el siguiente comando.

```
$ go test -v ./utils
```
