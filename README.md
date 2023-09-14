# Packaging Solution Readme

This application is designed to calculate the number of packs required to fulfill an order based on specific pack sizes. It follows the following rules:

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfill the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfill each order.

To interact with the API, a user interface (UI) has been developed.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage Server](#usage-server)
- [Usage UI](#usage-ui)
- [Testing](#testing)
---

## Getting Started

### Prerequisites

- [Golang](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository:

```bash
git clone git@github.com:istomin10593/gimshark-test.git
cd gimshark-test
```
2. Build Docker images and launch both the server and UI in a Docker environment:
```
make up
```

3. Stop the server and UI:

```
make down
```

4. Run the server on local machine:

```
make server-run
```

5. Run UI on local machine:

```
make ui-run
```

## Usage Server
The application provides an HTTP API to calculate the number of packs needed for a given order. You can interact with the API by sending a POST request to http://localhost:40999/packs with the items quantity in the request body.

Example using cURL:
```bash
curl -X POST http://localhost:40999/packs -d '{"items": 251}'
```
The response will be a JSON object containing the calculated packs:

```bash
'{"500":1}'
```
You can modify the pack sizes without changing the code. Response can be customized by adjusting the pack sizes in `server/conf.yaml`.

## Usage UI
1. Open your web browser and navigate to http://localhost:40998/packs. 
2. Input the "Quantity of Items" and click the "Calculate Packs" button.
3. You will be redirected to http://localhost:40998/packs/calculate with the resulting pack sizes displayed.

## Testing
Tests are provided for server logic with pack size counting, and including both integration and unit tests.
To run the tests for the server, use the following command:
```
make test
```