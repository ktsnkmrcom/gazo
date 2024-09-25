# gazo

## Create a temporary PNG image.

Flags are easy.  
Because they are similar to html img tags.

### Installing

```
go install github.com/ktsnkmrcom/gazo@latest
```

Use "gazo --help" and "gazo img --help" for more information about a command.

### Examples:

```
gazo img --src example --alt Word --width 1080 --height 1080 --hex e0e0e0 --path $HOME
```

### Flags

--src : File name of png image. (default "temporary")  
--alt : Word to insert into image. Alphabetic characters only. default is no insertion. (default " ")  
--width : Width in pixels (default 1080)  
--height : Height in pixels (default 1080)  
--hex : Specify the background color. "random" will choose a random color. (default "e0e0e0")  
--path : Output Directory (default "$HOME")  
