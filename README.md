# Typer

Typing test in your terminal

![Typer Banner](../assets/banner.png)

### Installation

To just install `typer` simply use this command:
```
go install github.com/maaslalani/typer@latest
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
// You can just run `typer` without installing it
go run .

// You can also build an executable
go build -o typer .

// You can also install the package, so you can run it as any other command
go install

// Instead of any of the commands above you can also use `make`
make
make build
make install
```
