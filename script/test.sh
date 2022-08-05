#!/bin/bash

curl http://localhost:8888/
curl http://localhost:8888/v0/
curl http://localhost:8888/v0/test
curl http://localhost:8888/v0/date
curl http://localhost:8888/index
curl http://localhost:8888/panic
curl http://localhost:8888/v1/
curl http://localhost:8888/v1/hello
curl http://localhost:8888/v2/hello/i0Ek3
curl http://localhost:8888/v2/login
ab -c 1000 -n 1000 http://localhost:8888/v2/login
