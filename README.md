# Go Todo

## Introduction

A simple todolist application written in Go 

## Requirements
* Docker
* Go

## Installation

* Clone this repo 

```bash
git clone https://github.com/cheorges/GoTodoWebbox.git
```

* Change Directory

```bash
cd GoTodoWebbox
```

```bash
docker-compse up -d
```

* Initiate `.env` file

```bash
cp .env.example .env
```

* Modify `.env` file with your correct database credentials and desired Port

## Usage

To run this application, execute:

```bash
go run main.go
```

You should be able to access this application at `http://127.0.0.1:4000`

## Note

Mongo Express is an interactive lightweight Web-Based Administrative Tool to effectively manage MongoDB Databases.
You should be able to access this application at `http://127.0.0.1:8081`

The data are stored in docker volume: `go-todo_mongodb_data_container`.
You can inspect the volume with: `docker volume inspect go-todo_mongodb_data_container`. 
You can delete the volume with: `docker volume rm mongodb_data_container`