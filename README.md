# ORG
## Usar
```go
go build
./org.exe -frm <string> -to <string> -regex <string>
```
## TODO
- [ ] Compatibilidad con Linux (estructura de carpetas)
- [ ] Buscar dentro de otras carpetas
- [X] Separar el código
- [X] Reescribir `func check(frm * string, to *string)`
- [X] Tomar regex como un input
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