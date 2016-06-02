# tagoæl

Enhanced Hello World written in Go with advanced configuration mechanism provided by Flæg &amp; Stært

## Why tagoæl ?

The goal is to understand how to use Flæg &amp; Stært. It's written in Go and french pronunciation of tagoæl means more or less 'shut up': what you will want to say to the program once you will test it with unmodified default values.

## Examples

Display automatically formatted help message with `-h` short flag:
```shell
$ ./tagoael.exe -h
tagoæl is an enhanced Hello World program to display messages with
an advanced configuration mechanism provided by flæg & stært.

flæg:   https://github.com/containous/flaeg
stært:  https://github.com/containous/staert
tagoæl: https://github.com/debovema/tagoael


Usage: tagoael [--flag=flag_argument] [-f[flag_argument]] ...     set flag_argument to flag(s)
   or: tagoael [--flag[=true|false| ]] [-f[true|false| ]] ...     set true/false to boolean flag(s)

Flags:
        --configfile                                       Configuration file to use (TOML). (default "tagoael")
        -i, --displayindex                                 Whether to display index of each message (default "false")
        -m, --messagetodisplay                             Message to display (default "Hello World")
        -n, --numbertodisplay                              Number of message to display (default "5")
        -h, --help                                         Print Help (this message) and exit
```

Revert to the simplest Hello World:
```shell
$ ./tagoael.exe --messageToDisplay='Hello world' --displayIndex=false --numberToDisplay=1
```

It is possible to combine short flags "-in10" = "--displayindex  --numbertodisplay=10":
```shell
$ ./tagoael.exe -in10 --messageToDisplay='Gooooooo!'
1: Gooooooo!
2: Gooooooo!
3: Gooooooo!
4: Gooooooo!
5: Gooooooo!
6: Gooooooo!
7: Gooooooo!
8: Gooooooo!
9: Gooooooo!
10: Gooooooo!
```

Last it is possible to provide configuration values from a TOML file
```shell
$ ./tagoael.exe --configFile=tagoael.toml.sample
1: Hello from TOML
2: Hello from TOML
3: Hello from TOML
4: Hello from TOML
5: Hello from TOML
6: Hello from TOML
7: Hello from TOML
8: Hello from TOML
9: Hello from TOML
10: Hello from TOML
```
