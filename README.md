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

If you want to provide your own text, you can pass in a file name. The typing test will use the contents of the specified file.
```
typer filename.txt
```

You can also use any arbitrary command, for example `curl`ing text from the internet to practice on.
```
typer $(curl -s https://loripsum.net/api)
```

### Demo
![typer](../assets/typer.gif?raw=true)

### Development
```
make
```
