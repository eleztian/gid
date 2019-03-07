# gid
golang goroutine id

[![Build Status][1]][2]
[![codecov][3]][4]
[![goreportcard for eleztian/gid][5]][6]
[![godoc for eleztian/gid][7]][8]
[![MIT Licence][9]][10]

# Use

```bash
go get -u github.com/eleztian/gid
```

# Example 

```go
id := gid.GoID()
```

```go
var id int64
var wg sync.WaitGroup
go gid.WithGoID(&id, func (){
	defer wg.Done()
	fmt.Println(id)
})
wg.Wait()

```

> Thanks https://github.com/huandu/go-tls for original idea


[1]: https://travis-ci.com/eleztian/gid.svg?branch=master
[2]: https://travis-ci.com/eleztian/gid
[3]: https://codecov.io/gh/eleztian/gid/branch/master/graph/badge.svg
[4]: https://codecov.io/gh/eleztian/gid
[5]: https://goreportcard.com/badge/github.com/eleztian/gid
[6]: https://goreportcard.com/report/github.com/eleztian/gid
[7]: https://godoc.org/github.com/eleztian/gid?status.svg
[8]: https://godoc.org/github.com/eleztian/gid
[9]: https://badges.frapsoft.com/os/gpl/gpl.svg?v=103
[10]: https://opensource.org/licenses/gpl-license.php