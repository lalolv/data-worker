# Overview

[![alt travis](https://travis-ci.org/lalolv/data-worker.svg?branch=master)](https://travis-ci.org/lalolv/data-worker)
[![alt report](https://goreportcard.com/badge/github.com/lalolv/data-worker)](https://goreportcard.com/report/github.com/lalolv/data-worker)

Data worker, data generator, generate a lot data by config.

[中文说明](./README_ZH.md)

## Application scenario

- Data test
- Demonstration
- Unit Test for developing

## Features

- [x] Define some specified fields and parameters by config.
- [x] Load external dict to use more data.
- [x] Custom formatted output data

## Usage

### Create a json file for config，and save to config folder。for example

```json
{
  "dict_path": "dict",
  "build": {
    "count": 5,
    "format": "csv",
    "name": "demo2",
    "path": "out"
  },
  "fields": [
    {
      "name": "order_id",
      "value": "po-{uuid}"
    },
    {
      "name": "email",
      "value": "{username}@{email}"
    },
    {
      "name": "creat_dt",
      "value": "{datetime}"
    }
  ]
}
```

- dict_path: It's a dict folder includes some dict files
- build: Build command
- fields: Setup output a few fields
- Use curly braces to customize the output variables. Variables come from data dictionaries and built-in types.

### Assign -c parameter，setup config

```shell
data-worker build -c ./config/demo.json
```

## Built-in types

| Type     |         Desc          |
| -------- | :-------------------: |
| string   |        String         |
| uuid     | Unique identification |
| mobile   |     Mobile phone      |
| idno     |       ID Number       |
| datetime |  Full date and time   |

## Plan

- [ ] Custom variables, increase filters.
- [ ] Support for more built-in types
