####Steps to build and run the co2-calculator

1. Go to the root directory of the project
2. Build and containerized the service, run 
```
docker build . -t co2-calculator
```
3. Run the service, use 
```
docker run -e ORS_TOKEN=<token> co2-calculator:latest --start <start_city> -end <end_city> --transportation-method=<transport-type>
```


####Remarks:
1. if the start and the end city is not set, it is default to `munich`
2. if the transportation method is not set, it is default to `small-diesel-car`


####Example
1. from beijing to shanghai
```
docker run -e ORS_TOKEN=<token> co2-calculator:latest --start beijing -end shanghai --transportation-method=large-petrol-car

Outputs:
start: beijing
end: shanghai
method: large-petrol-car
Your trip caused 340.99kg of CO2-equivalent

```

2. from munich to frankfurt (do not need to specify munich as starting point, do not need to specify car type)
```
docker run -e ORS_TOKEN=<token> co2-calculator:latest -end frankfurt 

Outputs:
start: munich
end: frankfurt
method: small-diesel-car
Your trip caused 56.42kg of CO2-equivalent
```

3. unknown transportation type
```
docker run -e ORS_TOKEN=<token> co2-calculator:latest -end frankfurt --transportation-method=unkown 

Outputs:
start: munich
end: frankfurt
method: unkown
panic: transportation method unkown does not exist

```

4. ORS_TOKEN Not provided
```
 docker run  co2-calculator:latest -end frankfurt

Outputs:
panic: Env var ORS_TOKEN Not set

```

5. wrong token
```$xslt
docker run -e ORS_TOKEN=<wrong-token> co2-calculator:latest --end frankfurt

Outputs:
start: munich
end: frankfurt
method: small-diesel-car
panic: Cannot get coordinates of munich
```

####Run unit tests 

detailed test report
```
docker run --volume=$(pwd):/ws  --workdir=/ws golang:1.12.1 /bin/bash -c "go test  -v ./..."
```

test with coverage info
```
docker run --volume=$(pwd):/ws  --workdir=/ws golang:1.12.1 /bin/bash -c "go test  -cover ./..."
```

Please be aware that we are running test in an empty golang container, so at beginning it will pull dependencies, which will take a bit time (20s)
