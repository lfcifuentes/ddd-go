# Go DDD Example

## Notes

### Domain
Se refiere a la área de conocimiento o actividad principal sobre la cual tu aplicación está construida.


### Entity
Una Entidad es un objeto del dominio que tiene una identidad única a lo largo del tiempo y entre diferentes instancias. Su identidad es lo que la distingue de otras entidades, incluso si tienen los mismos atributos. Las entidades tienen las siguientes características:

- **Identidad Única:** Cada entidad debe tener un identificador único (por ejemplo, ID).
- **Estado Mutable:** Las entidades pueden cambiar su estado a lo largo del tiempo.
- **Igualdad por Identidad:** Dos entidades son iguales si tienen el mismo identificador, sin importar los otros atributos.

### Value Object
Un Value Object es un objeto que no tiene identidad propia. En lugar de ser definido por un identificador único, un Value Object se define por sus atributos. Los Value Objects son inmutables y su igualdad se basa en la comparación de sus atributos.
Características Clave de los Value Objects:

- **Inmutabilidad:** Una vez creado, el estado de un Value Object no puede cambiar. Si necesitas un Value Object diferente, debes crear uno nuevo.
- **Sin Identidad:** Los Value Objects no tienen identidad propia. Dos Value Objects son iguales si todos sus atributos son iguales.
- **Comparación por Valor:** La igualdad se determina comparando los valores de sus atributos, no por una identidad única.
- **Encapsulación de Conceptos:** Los Value Objects encapsulan conceptos o atributos del dominio que no necesitan identidad propia.

Los Value Objects son frecuentemente utilizados dentro de los Agregados para encapsular y modelar atributos o conceptos del dominio que no necesitan identidad propia.

### Aggregate
Un Agregado es un patrón que define un grupo de Entidades que están relacionadas entre sí y que deben ser tratadas como una única unidad de consistencia para propósitos de datos. Un agregado tiene las siguientes características:

Raíz del Agregado (Aggregate Root): Es la entidad principal del agregado. Todos los accesos externos al agregado deben hacerse a través de esta entidad.
Consistencia Interna: Las reglas de negocio y las invariantes se aplican a nivel del agregado para garantizar la consistencia interna.
Encapsulamiento: Las entidades dentro del agregado solo pueden ser modificadas a través de la raíz del agregado.

### Factories
El patrón de fábrica es un patrón de diseño que se utiliza para encapsular lógica compleja en funciones que crean la instancia deseada, sin que la persona que llama sepa nada sobre los detalles de implementación.

### Repository
Es un patrón que se basa en ocultar la implementación de la solución de almacenamiento/base de datos detrás de una interfaz. Esto nos permite definir un conjunto de métodos que deben estar presentes y, si están presentes, están calificados para usarse como repositorio.

### Services
Un servicio vinculará todos los repositorios débilmente acoplados en una lógica empresarial que satisfaga las necesidades de un determinado dominio.


## CMD

#### Testing

Correr los test
```bash
go test -cover ./...
```
Limpiar la cache
```bash
go clean -testcache 
```
Ver el coverage de los test
```bash
go test  ./... -coverprofile=coverage.out
``` 
```bash
go tool cover -html=coverage.out
```

#### Run Server

Run development server 
```bash
go run main.go serve
```