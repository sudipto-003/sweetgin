
## Demo REST API project using Go-Gin

Demonstration of a **CURD** REST API using go-gin, mongodb and Redis. This project demonstrates the CURD applications for parcel courier delivery system, creating new parcel delivery, retriving parcels based on unique parcel id and dates, updating parcel delivery status with receiver OTP verification. 

## Architecture Overview

![Project Architecture Overview](images/arch-overview.png)

### Description

Go-gin router handles the http requests, dispatches the request to appropiate handlers that invokes the mongodb repository to create, retreive and update documents and returns the response as JSON object with proper http status code and error message(if any). A secondary Redis server is used to store the OTP for parcel receiver verification.

## API Endpoints

