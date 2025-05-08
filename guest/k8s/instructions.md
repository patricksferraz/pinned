# Kubernetes Deployment Instructions - Guest Service

This document provides step-by-step instructions for deploying the Guest Service to a Kubernetes cluster.

## Prerequisites

- Kubernetes cluster up and running
- `kubectl` configured to communicate with your cluster
- Docker registry credentials
- Environment variables for the service

## Deployment Steps

### 1. Configure Environment Variables

1. Navigate to the service's k8s directory:
   ```bash
   cd guest/k8s
   ```

2. Create your environment file:
   ```bash
   cp .env.example .env
   ```

3. Edit the `.env` file with your specific configuration values.

### 2. Create Kubernetes Secrets

#### Service Secrets
Create a secret containing the service's environment variables:
```bash
kubectl create secret generic guest-secret --from-env-file k8s/.env
```

#### Docker Registry Secrets
Create a secret for accessing the Docker registry:
```bash
kubectl create secret docker-registry regsecret \
--docker-server=$DOCKER_REGISTRY_SERVER \
--docker-username=$DOCKER_USER \
--docker-password=$DOCKER_PASSWORD \
--docker-email=$DOCKER_EMAIL
```

Required environment variables:
- `DOCKER_REGISTRY_SERVER`: Docker registry URL
- `DOCKER_USER`: Registry username
- `DOCKER_PASSWORD`: Registry password
- `DOCKER_EMAIL`: (Optional) Any valid email address

### 3. Deploy the Service

Deploy all Kubernetes resources:
```bash
kubectl apply -f ./k8s
```

### 4. Verify Deployment

Check the deployment status:
```bash
kubectl get pods -l app=guest
kubectl get services -l app=guest
```

## Troubleshooting

### Common Issues

1. **Pod fails to start**
   - Check pod logs: `kubectl logs <pod-name>`
   - Verify secrets are properly created
   - Ensure environment variables are correctly set

2. **Image pull errors**
   - Verify Docker registry credentials
   - Check network connectivity to registry
   - Ensure image exists in registry

### Useful Commands

```bash
# View pod logs
kubectl logs -f <pod-name>

# Describe pod for detailed information
kubectl describe pod <pod-name>

# Check service endpoints
kubectl get endpoints guest-service

# View deployment status
kubectl rollout status deployment/guest
```

## Cleanup

To remove the deployment:
```bash
kubectl delete -f ./k8s
kubectl delete secret guest-secret
kubectl delete secret regsecret
```

## Additional Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Docker Registry Authentication](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/)
- [Kubernetes Secrets Management](https://kubernetes.io/docs/concepts/configuration/secret/)
