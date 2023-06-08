# URL Shortener
This repository contains Golang implementation of a url shortener.

## Install

To install, download or clone the repo, then:

```
npm install
```

## Usage

Firstly, we need Docker and Go version 1.19.3. To run this, on the terminal run these commands in order:

```
make postgres
```
```
make server
```
This will generate a container and database with the title `urlshortener`. Then we need to create a table on that database.
```
CREATE TABLE "urls" (
  "id" varchar(5) PRIMARY KEY,
  "long_url" varchar NOT NULL,
  "short_url" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);
```
Once this is complete, we can try to run the server to verify that it works
```
make server
```
Optionally, we can run these tests:
```
make test
```
