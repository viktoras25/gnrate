### Gnrate

A small utility to generate random names, addresses, phone numbers, etc.
Nothing really fancy, just [github.com/dmgk/faker](https://github.com/dmgk/faker) + [github.com/urfave/cli/v2](https://github.com/urfave/cli/v2) glued together.

### Installation

```
go install github.com/viktoras25/gnrate@latest
```

### Usage

```
$ gnrate 5 german names
Hr. Pascal Vogelgsang
Fr. Emma Wartenberg
Maryam vom Vogelgsang
Charlotte Osch
```

You can provide 1 to 3 arguments: count (default = 1), language (default = english), type (mandatory)
