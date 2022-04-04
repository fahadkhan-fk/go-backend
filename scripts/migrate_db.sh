#!/bin/bash

cd migrations
goose postgres "user=postgres dbname=go-backend sslmode=disable password=postgres" up