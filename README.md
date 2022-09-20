# simple HTTP Server

# routes: 
```
/ => returns a static page
/Hello =>  returns Simple Message
/Form => post request, returns a message with name, age that are submitted
```

# tests:
```bash
python3.7 tests/test.py
test_main_route: passed
test_form_route: passed
test_hello_route: passed
```

# build

```bash
go build main.go
```