# ORG
Small program to organize pictures with names similar to IMG20230112.png --> year, month, day.
## Use
If you want to build from source
```
go build -ldflags="-w -s"
```
This is how you use the program
```
./org.exe -from <string> -to <string> -regex <string> -recursive <string>
```
The `-from` flag is the folder you want to read from, all the files in here will be moved to the folder you used in the `-to` flag. The program will make new folders to store all the images that follow the regex string into the `-to` folder you gave as input, each image will be stored in they own folder.

The `-recursive` and the `-regex` are optional. If you don't use the `-recursive` flag then the program won't try to read any images that are in subdirectories, and if you don't use the `-regex` flag then the regex used will be `2023(\w\w)`. Be careful with the regex string you give as an input because the code was made with the intent to find filenames like this IMG20230112.png, submatching the month.

## Examples
Let's say you have a folder in your current directory named Images and you want to move all files within that folder to a folder named Backups. The folder named Backup doesn't have any files nor folders inside of it.
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
Doing the following command
```
org -from Images -to Backup 
```
Will do this 
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
If you have something like this:
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
You can do this to search all the folders inside Images, doing this will have the same result as the previous example
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
