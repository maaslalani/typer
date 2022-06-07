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

You can use Monkeytype as a source of words, just pass `-m, --monkeytype` flag,
by default it'll use `english` dictionary, you can change that by adding `--monkeytype-language string` additionally.
```
typer -m --monkeytype-language english_1k
```

If you want to provide your own text, you can pass in a file name with the `--file` flag. The typing test will use the contents of the specified file.
```
typer -f filename.txt
```

You can also pipe data by `stdin`.
```
echo 'Text from stdin!' | typer
```

### Themes

There is pretty basic theme support, theme should be saved in config file (default `$HOME/.typer.yaml`) and should look similar to this (almost default) theme:

```yaml
theme:
  bar:
    color: '#4776E6' # basic color of the progressbar
    #gradient: '#ff0000' # if passed, will generate a gradient from previous color to this one
  graph:
    color: blue # does not use rgb values
    height: 3   # height of the graph
  text:
    error: # color when misspelled
      background: '#f66'
      foreground: '#fff'
    typed: # color when character have been typed
      foreground: '#fff'
      #background: '#000' # optional, default theme does not add background
    untyped: # color when still haven't been typed
      foreground: '#ddd'
      #background: '#000' # optional, default theme does not add background

```

### Demo
![typer](../assets/typer.png?raw=true)
