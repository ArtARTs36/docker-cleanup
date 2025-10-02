# docker-cleanup

![Docker Image Version](https://img.shields.io/docker/v/artarts36/docker-cleanup?style=for-the-badge&logo=docker&label=Image%20Version&link=https%3A%2F%2Fhub.docker.com%2Fr%2Fartarts36%2Fdocker-cleanup)
![Docker Image Size](https://img.shields.io/docker/image-size/artarts36/docker-cleanup?style=for-the-badge&logo=docker&label=Image%20Size&link=https%3A%2F%2Fhub.docker.com%2Fr%2Fartarts36%2Fdocker-cleanup)


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
