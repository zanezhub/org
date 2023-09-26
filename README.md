# ORG
## Usar
```go
go build
./org.exe -frm <string> -to <string>

```

## TODO
- [X] Separar el código
- [X] Reescribir `func check(frm * string, to *string)`
- [ ] Compatibilidad con Linux (estructura de carpetas)
- [ ] Tomar regex como un input
- [X] Directorios en el mismo directorio
```
(.) You are here
test <|
      |->des
      |->from
           |-> Files you want to move 
```
```
./org -frm .\test\from -to .\test\des
```

## Features
- [ ] Buscar dentro de otras carpetas
