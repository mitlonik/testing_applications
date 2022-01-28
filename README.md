# testing_applications
Applications in different programming languages for testing various tasks


1. A simple on golang application to send data to 'stdout' stream.

## For Mac OS X Users

You may get an error:
``` bash
failed to create network testing_applications_net: Error response from daemon: This node is not a swarm manager. Use "docker swarm init" or "docker swarm join" to connect this node to swarm and try again.
```
The thing is that docker on MacOSX does not know  networks driver: 'overlay'. Comment it out.
