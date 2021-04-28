# Typer

Typing test in your terminal

![Typer Banner](../assets/banner.png)

### Installation
```
go get github.com/maaslalani/typer
```

### Usage
To begin a typing test simply type `typer`. This will generate random words for you to type and show you your WPM score.
```
typer
```

To change the length of the typing test, use the `--length` flag.
```
typer -l 20
```

If you want to provide your own text, you can pass in a file name with the `--file` flag. The typing test will use the contents of the specified file.
```
typer -f filename.txt
```

### Demo
![typer](../assets/typer.gif?raw=true)

### Development
```
make
```
