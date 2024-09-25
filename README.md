# gazo

## Create a temporary PNG image.

Flags are easy.  
Because they are similar to html img tags.

Use "gazo img --help" for more information about a command.

### Examples:

```
gazo img --src example --alt Word --width 1080 --height 1080 --hex e0e0e0 --path $HOME
```

### Flags

--src string   File name of png image. (default "temporary")  
--alt string   Word to insert into image. default is no insertion.  
--width int    Width in pixels (default 1080)  
--height int   Height in pixels (default 1080)  
--hex string   Specify the background color. "random" will choose a random color. (default "e0e0e0")  
--path string   Output Directory (default "$HOME")  
-h, --help         help for img
