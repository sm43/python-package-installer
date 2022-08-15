# Python Package Installer

Python Package Installer is an HTTP Service written in Go. This exposes an API to install python packages on the server.
The only supported command at this moment is `pip install <package names>`.

The packages are installed in a venv and then zipped and copied to location `$HOME/python-packages`. The location can be 
configurable.

Prerequisites:
* go 1.17 or more
* python3 
* pip

To run the service on locally, clone the code and execute
```shell
  make run    
```
The service will run on port `8080`, can be configurable using flag.

You can hit the API using curl as follows

```shell
  curl -X POST http://localhost:8080 \
    -d '{"command": "pip install numpy"}'
```
This will register a request to install numpy package.

You can also pass a list of packages
```shell
  curl -X POST http://localhost:8080 \
    -d '{"command": "pip install django numpy tensorflow git+https://github.com/httpie/httpie.git#egg=httpie fastapi deta"}'
```

### How this works

- The service use worker queues to handle the installation.
- When a request is received, it is queued and as a worker is idle, the installation is given to that worker.
- The number of workers by default are 2, which is also configurable by a flag.
