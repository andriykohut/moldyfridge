moldyfridge
===========

*Eat your food before it gets moldy*
simple fridge tracker in golang

```bash
moldyfridge -h
moldyfridge.

Usage:
  moldyfridge (add | rm) <food>...
  moldyfridge (add | rm) <food> [--amount <amount>]
  moldyfridge ls
  moldyfridge search <food>
  moldyfridge reset
  
Options:
  -h --help     Show this screen.
  --version     Show version.
```

## Installation
`go get github.com/andriykohut/moldyfridge`

## Usage
**Adding food**
```
moldyfridge add Cucumbers Tomatoes
moldyfridge add Beer --amount 3
```
**Listing food**
```
moldyfridge ls
Cucumbers: 1, age - 2m
Tomatoes: 1, age - 2m
Beer: 3, age - just now
```
**Searching food**
```
moldyfridge search bee
```

## TODO
- Display output in some nice tables
- Remove food when amount <= 0
- Collect stats
