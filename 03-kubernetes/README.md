# 03-kubernetes - развёртывание приложения в Kubernetes

Развёртывание приложения `no1d3d/go-info-app:1.0.0` в локальном кластере minikube.

## Что сделано

- Кластер minikube с драйвером Docker.
- Ingress-контроллер NGINX включён через `minikube addons enable ingress`.
- Deployment с 3 репликами в namespace `go-info-app`.
- Service типа ClusterIP.
- Ingress-правило для доступа к приложению.
- Переменная `AUTHOR` передаётся через ConfigMap.
- Настроены readiness (`/ready`) и liveness (`/live`) пробы.

## Развёртывание

```bash
minikube start --driver=docker
minikube addons enable ingress

kubectl apply -f namespace.yaml
kubectl apply -f configmap.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml
```

## Проверка

```bash
kubectl get pods -n go-info-app
kubectl get ingress -n go-info-app

IP=$(kubectl get ingress -n go-info-app -o jsonpath='{.items[0].status.loadBalancer.ingress[0].ip}')
curl -H "Host: go-info.example" http://$IP/
curl -H "Host: go-info.example" http://$IP/health
```

Ожидаемо: HTML-страница с hostname, IP и Author; {"status":"ok"} на /health.

## Примечание о доступе

Приложение доступно только с хоста, где запущен minikube (через localhost
или внутренний IP 192.168.49.2). Другие устройства в локальной сети получить доступ не смогут.

Чтобы открыть доступ в локальную сеть, выполните на хосте:
```bash
kubectl port-forward -n go-info-app service/go-info-app 8080:80 --address 0.0.0.0
```

После этого приложение будет доступно по адресу http://<IP-хоста>:8080 с любого устройства в той же сети. Туннель работает, пока открыт терминал с командой. Сеть на сервере должна быть настроена как Bridged Adapter.