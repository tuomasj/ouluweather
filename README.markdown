# Ouluweather

This is a really simple Go command-line application, that fetches current weather
information from VTT weather station at http://weather.willab.fi/weather.html and displays it on command-line.


## Example output

```
~ $ ouluweather
2015-09-21 00:17 EEST (60 seconds ago)
Oulu Linnanmaa temperature: +10.30C
```


## How to compile the program

### Compile to file

Compile code to binary file:

```
~ $ go build
```

Now you have a binary file called `ouluweather` in the same directory as source code.

It can now be run with:

```
~ $ ./ouluweather
2015-11-21 18:36 EET (41 seconds ago)
Oulu Linnanmaa temperature: -2.10C
```


### Compile on fly

Another way to try the source code is to run it directly with the Go compiler.
Probably more slow way, but useful approach while developing improvements to the tool.

```
~ $ go run ouluweather.go
2015-11-21 18:38 EET (71 seconds ago)
Oulu Linnanmaa temperature: -2.20C
```
