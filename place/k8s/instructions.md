# Instructions k8s

## Create secrets

Fill in the _.env_ file in _k8s/.env_ using _.env.example_, and run:

```sh
kubectl create secret generic place-secret --from-env-file k8s/.env
```

## Create a secret for docker registry

Run:

```sh
kubectl create secret docker-registry regsecret \
--docker-server=$DOCKER_REGISTRY_SERVER \
--docker-username=$DOCKER_USER \
--docker-password=$DOCKER_PASSWORD \
--docker-email=$DOCKER_EMAIL
```

Where:

- $DOCKER_REGISTRY_SERVER: is the url for the registry
- $DOCKER_USER: is the registry user
- $DOCKER_PASSWORD: is the registry password
- $DOCKER_EMAIL: is optional, any e-mail

## Deploy all

Run:

```sh
kubectl apply -f ./k8s
```
