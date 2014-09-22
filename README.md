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

+-----------+----------+--------+
| name      | age      | amount |
+-----------+----------+--------+
| Cucumbers | 1m       | 1      |
+-----------+----------+--------+
| Tomatoes  | 1m       | 1      |
+-----------+----------+--------+
| Beer      | just now | 3      |
+-----------+----------+--------+
```
**Searching food**
```
moldyfridge search matoe

+----------+-----+--------+
| name     | age | amount |
+----------+-----+--------+
| Tomatoes | 1m  | 1      |
+----------+-----+--------+
```

## TODO
- Display output in some nice tables
- Remove food when amount <= 0
- Collect stats
