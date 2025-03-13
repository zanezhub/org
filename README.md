# ORG
Programa pequeño para organizar imágenes con nombres similares a IMG20230112.png --> año, mes, día.

## ¿Por qué usar ORG?
Si eres de los que tienen miles de fotos desordenadas en sus carpetas, ORG es la herramienta perfecta para ti. Con esta sencilla aplicación, puedes organizar tus imágenes rápidamente, basándote en su fecha de creación, moviéndolas a carpetas que se organizan automáticamente por año, mes y día.

Olvídate de perder tiempo buscando esas fotos que se encuentran en carpetas caóticas. ORG hace todo el trabajo por ti, permitiéndote ordenar y almacenar tus imágenes de forma eficiente y sin esfuerzo. Además, con la opción de hacerlo de manera recursiva, podrás organizar no solo las imágenes en tu carpeta principal, sino también las de subdirectorios.

**¿Por qué elegir ORG?**

- **Automatización:** Organiza tus imágenes automáticamente según la fecha, con una estructura de carpetas clara y ordenada.
- **Fácil de usar:** Con solo un par de comandos, tendrás tus imágenes organizadas en minutos.
- **Personalizable:** Si necesitas una organización más específica, puedes modificar el regex y las opciones de búsqueda.
- **Recursividad:** No solo mueve las imágenes de una carpeta, sino que también las organiza en subcarpetas, ¡sin que tengas que mover un dedo!

Haz que tu vida digital sea mucho más ordenada y eficiente con ORG. ¡Prueba ahora mismo y organiza tus fotos como un profesional!

## Uso
Si deseas compilar desde el código fuente:
```
go build -ldflags="-w -s"
```
Así es como se usa el programa:
```
./org.exe -from <string> -to <string> -regex <string> -recursive <string>
```
La opción `-from` es la carpeta de donde deseas leer, todos los archivos en esta carpeta serán movidos a la carpeta que indiques en la opción `-to`. El programa creará nuevas carpetas para almacenar todas las imágenes que coincidan con la cadena regex que proporcionaste en la opción `-to`, cada imagen se almacenará en su propia carpeta.
Las opciones `-recursive` y `-regex` son opcionales. Si no usas la opción `-recursive`, el programa no intentará leer imágenes en subdirectorios, y si no usas la opción `-regex`, la expresión regular por defecto será `2023(\w\w)`. Ten cuidado con la cadena regex que proporcionas como entrada, ya que el código está diseñado para encontrar nombres de archivos como IMG20230112.png, subcombinando el mes.

## Ejemplos
Supongamos que tienes una carpeta en tu directorio actual llamada `Images` y deseas mover todos los archivos dentro de esa carpeta a una carpeta llamada `Backups`. La carpeta `Backup` no tiene archivos ni carpetas dentro de ella.
```
$ ls Images/ IMG-2023011.png IMG20230211.png IMG20230212.png IMG20230511.png IMG20230811.png Backup/ ..
```
Al ejecutar el siguiente comando:
```
org -from Images -to Backup
```
Se hara lo siguiente
```
$ ls
Images/
      IMG-2023011.png
      IMG20230211.png
      IMG20230212.png
      IMG20230511.png
      IMG20230811.png
Backup/
      ..
```
Usando el siguiente comando:
```
org -from Images -to Backup 
```
Se obtendra esto:
```
$ ls
Images/
Backup/
      01/
          IMG-2023011.png
      02/
          IMG20230211.png
          IMG20230212.png
      05/
          IMG20230511.png
      08/
          IMG20230811.png
```
Si tienes algo similar a:
```
$ ls
Images/
      Dir1/
          IMG-2023011.png
          IMG20230211.png
          IMG20230212.png
      Dir2/
          IMG20230511.png

      IMG20230811.png
Backup/
      ..
```
Puedes hacer esto para buscar todas las carpetas dentro de `Images`, lo que tendrá el mismo resultado que el ejemplo anterior:
```
org -from Images -to Backup -recursive true
```
## TODO
- [X] Compatibilidad con Linux (estructura de carpetas)
- [X] Buscar dentro de otras carpetas
- [X] Separar el código
- [X] Reescribir `func check(frm * string, to *string)`
- [X] Tomar regex como un input
- [X] Directorios en el mismo directorio
