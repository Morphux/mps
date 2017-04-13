<p align="center">
<img src="https://raw.githubusercontent.com/Morphux/Graphic/master/logo/single_penguin.png" /><br />
</p>
<p align="center">
<img src="https://img.shields.io/badge/language-go-blue.svg" /> &nbsp;
<img src="https://img.shields.io/badge/license-Apache--2.0-yellow.svg" /> &nbsp;
<a href="https://travis-ci.org/Morphux/mps"><img src="https://travis-ci.org/Morphux/mps.svg?branch=master"/></a> &nbsp;
<!--<a href="https://scan.coverity.com/projects/morphux-libmpm">
  <img alt="Coverity Scan Build Status"
       src="https://scan.coverity.com/projects/11577/badge.svg"/>
</a>&nbsp;
<a href="https://codecov.io/gh/Morphux/libmpm">
  <img src="https://codecov.io/gh/Morphux/libmpm/branch/master/graph/badge.svg" alt="Codecov" />
</a>-->
<br />
<h1 align="center" style="border:none">Morphux/mps</h1>
<h6 align="center">Morphux Package Server</h6>
</p>
<p align="center">
<a href="#install">Install</a> • <a href="#test">Test</a> • <a href="#use">Use</a> • <a href="#documentation">Documentation</a>
</p>

# Clone
```
git clone https://github.com/Morphux/mps --recursive
```

# Install

You need to have a correct [GOPATH](https://golang.org/doc/code.html#GOPATH)

```
go get -u github.com/Morphux/mps
```

# Test

```
make test
```

# Use

Run as follow :


```
mps -db data.db 127.0.0.1:9090
```



options:
```C
mps usage : mps host:port
    -db <database file> : import sqlite3 db
```

# Documentation
