errors

A simple errors package

go get github.com/0chain/errors

we introduce a new applicaiton error which has errorCode and errorMsg. 

```
type Error struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg"`
}
```

## New Error

Then `errors.New` function returns a new error given the code (optional) and msg

two arguments can be passed!
1. code
2. message
if only one argument is passed its considered as message
if two arguments are passed then
	first argument is considered for code and
	second argument is considered for message

```
 applicationError := errors.New("401", "Unauthorized")
 simpleError = errors.New("validation failed")
```


## Standard error interface implementation

This is what is printed when you do `.Error()` for the above example

```
fmt.Println(auth("username", "password"))

401: Unauthorized
validation failed
password mismatch
```

## Error propagation

The `errors.Wrap` function returns a new error that adds context to the original error. You can wrap using a msg or error. For example
```
var ErrPasswordMismatch = errors.New("password mismatch") // "invalid argument"
var ErrUnAuthorized = errors.New("401", "Unauthorized")

func auth(username, password string) error {
    err := validate(username, password)
    if err != nil {
        return errors.Wrap(err, ErrUnAuthorized)
    }
}


func validate(username, password string) error {
    err := passwordValidation(password)
    if err != nil {
        return errors.Wrap(err, "validation failed")
    }
}

func passwordValidation(password string) error{
    // on invalid password
    return ErrPasswordMismatch
}

```

The `errors.UnWrap` function returns the current error and the previous error

```
current, previous := errors.UnWrap(auth("username", "password"))

fmt.Println(current) => 401: Unauthorized
fmt.Println(previous) => validation failed
                         password mismatch

// futher more

current, previous := errors.UnWrap(previous)

fmt.Println(current) => validation failed
fmt.Println(previous) => password mismatch

// further more

current, previous := errors.UnWrap(previous)

fmt.Println(current) => password mismatch
fmt.Println(previous) => nil
```

Retrieving the cause of an error, The `errors.Cause` function is the way to go

```
cause = auth("username", "password")

fmt.Println(cause) => password mismatch

```



