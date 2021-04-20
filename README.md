# srvcmd

Invoke a command on the host machine via HTTP and see the results.

**DISCLAIMER:** This tool was created with the intention of using on a local machine. Beware of exposing the HTTP server to the public.

## Start
```
go run . echo 'hello,' 'world'
```

## Invoke
```
curl -i http://localhost:5050/
```