# Overview

[![alt travis](https://travis-ci.org/lalolv/data-worker.svg?branch=master)](https://travis-ci.org/lalolv/data-worker)
[![alt report](https://goreportcard.com/badge/github.com/lalolv/data-worker)](https://goreportcard.com/report/github.com/lalolv/data-worker)

Data worker, data generator, generate a lot data by config.

[中文说明](./README_ZH.md)

## Where is for?

- Data test
- Demo
- Unit Test for developing

## Features

- [x] Define some specified fields and parameters by config.
- [x] Load external dict to use more data.

## Usage

### Create a json file for config，and save to config folder。for example：

```json
{
  "dict_path": "dict",
  "build": {
    "count": 5,
    "format": "json",
    "name": "demo",
    "path": "out"
  },
  "fields": [
    {
      "name": "id",
      "type": "uuid"
    },
    {
      "name": "name",
      "type": "string",
      "dict": "usernames"
    },
    {
      "name": "creat_dt",
      "type": "datetime"
    }
  ]
}
```

- dict_path: It's a dict folder includes some dict files
- build: Build command
- fields: Setup output a few fields

### Assign -c parameter，setup config。

```shell
data-worker build -c ./config/demo.json
```

## Plan

- Output cvs file
- Suport generate email address
