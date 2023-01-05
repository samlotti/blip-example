# blip-example
Example go project to show blip template usage


# install

Clone this repository

```shell
  git clone https://github.com/samlotti/blip-example.git
  
  cd blip-example
```



Install blip if you have not already done that: [Blip] (https://github.com/samlotti/blip)

Generate the template Go file but running ./buildTemplates.sh

```shell
    ./buildTemplates.sh
    
    or 
    
    blip -dir="./html" --rebuild --watch
```

Note that watch will monitor for file changes and rebuild as needed.

Generated files will be in the folder

```shell
  ./blipped
```

Start the application using 

```shell
    air
    
    or 
    
    go run blipServer.go

```

Note air will restart the application if the templates have changed.

open:  http://localhost:8181

Enjoy

