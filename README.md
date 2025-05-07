# docker-cleanup

docker-cleanup - is lightweight image for prune containers and images from Docker host

## Deploy to Swarm

```yaml
services:
  docker-cleanup:
    image: artarts36/docker-cleanup:0.1.0
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
    command: --containers --images
    deploy:
      mode: global
      restart_policy:
        delay: 8h
```
