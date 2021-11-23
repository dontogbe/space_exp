# space_exp
## Installation
Run
```azure
go install
```

## Testing
Using Postman or any other tool you send a post request to 
```azure
http://localhost:8080/map
```
Example of request that should be contained in the body
```azure
{
    "x":123.12,
    "y":456.56,
    "z":789.89,
    "vel":20.0
}
```