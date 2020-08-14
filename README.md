 # Install
 ` go get github.com/voltento/nomd`
 # Uninstall 
 ` go clean -i github.com/voltento/nomd`
 
 # Usage
 `md <text>`
 
 # Example
 Call 
 ``` 
    nomd \
    "# Hello world
    [https://stackoverflow.com/questions/](https://stackoverflow.com/questions/)
    **bold text underlined text**"
```
output
```
    Hello world
    https://stackoverflow.com/questions/
    bold text underlined text
```
 
 