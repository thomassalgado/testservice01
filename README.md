# Go lang API Sample

## Start Service

To start the service use the following command:

go run main.go

The dependencies are organized with govendor, so there is no need to execute the go get command

The server will respond to the https://localhost:8080/numbers URL

## Using the service

To use the service, do a http GET request to the https://localhost:8080/numbers without parameters
or passing:

begin: the first number of the sequence
end: the last number of the sequence

By default, the values of begin and end are 0 and 100 so there is no need to pass
both parameters

The return will be a JSON with the sequence of numbers between begin and end, fallowing the rule:

For multiples of 3, print the word 'Pé'instead of the number.
For multiples of 5, print the word 'Do' instead of the number.
For multiples of both 3 and 5, print 'PéDo'.
Else print the number itself

## Constraints

Begin and end parameters must be unsigned integers lower than 100

## Errors

If invalid parameters were passed, the server will return a JSON with a errro error message
and a 400 http status
