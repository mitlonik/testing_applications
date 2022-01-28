useful links:

[Managing docker logs with ELK and Fluentd](https://coralogix.com/blog/managing-docker-logs-with-elk-and-fluentd/)


Manual

``` bash
$ docker stop simple_nodejs_app
$ docker rm simple_nodejs_app
$ docker run -d --name simple_nodejs_app
--log-opt mode=non-blocking
--log-opt max-buffer-size=4m
node-log-app
```
There are two log options we have specified here that will be passed to the container:

    mode=non-blocking – this should be self-explanatory. This is how we tell Docker that we want to use the asynchronous method of collecting log entries from the container.

    max-buffer-size=4m – since we have specified the non-blocking mode, Docker will use a ring buffer to temporarily collect our logs before asynchronously passing them to the log driver. In our case, we have specified a buffer size of 4 megabytes. The default is 1 megabyte if this option isn’t specified.
