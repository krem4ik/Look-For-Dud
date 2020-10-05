# Look-For-Dud server-side

## Description

Look4dood is a web service that allows to detect, explore, report on combusting locations

## Project setup
go get -u -d ./...

## Compiles and hot-reloads for development
go run app/main.go

## Compiles and minifies for production
go build app/main.go

The project is not for production. At the moment, there is an API for registering users and uploading the location of the fire and the smoke from it.
The next step is to use a machine learning tool for processing satellite data and cloud detection.
