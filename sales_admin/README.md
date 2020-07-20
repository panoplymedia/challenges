# Sales Admin

## Installation

### Dependencies
To run the application the following dependencies need to be installed:
| Depdendency    | Version     | Installation Instructions                |
| :------------- | :---------- | :------------------------                |
| Docker         | 19.03.12+   | https://docs.docker.com/engine/install/  |
| Docker Compose | 1.17.1+     | https://docs.docker.com/compose/install/ |


## Running
To start the application and it's dependencies run:
```shell
make up
```
and navigate to http://localhost:8888/sales

### Viewing the logs
If you would like to view the logs run:
```shell
    make tail-logs
```

Press **CTRL + C** to stop viewing the logs.

## Stopping
To stop the service and it's dependencies run:
```shell
make down
```

### Testing
To run the tests run:
```shell
make test
```

