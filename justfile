region := "germanywestcentral"
rg := "web-shop"
cluster := "aks-cluster"

deploy:
    cd tf
    tofu apply -var="location={{region}}" -auto-approve
    cd ..
    kubectl create namespace argocd
    kubectl apply -n argocd --server-side --force-conflicts -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

connect:
    az aks get-credentials --resource-group {{rg}} --name {{cluster}} --overwrite-existing

sync:
    kubectl apply -f app/webshop.yml
    kubectl apply -f app/lgtm-stack.yml
    # Ne das stimmt so noch nicht

monitor:
    kubectl port-forward svc/lgtm-stack-grafana 3000:80 -n lgtm
    
monitor-password:
    kubectl get secret lgtm-stack-grafana -n lgtm -o jsonpath="{.data.admin-password}" | base64 -d 

destroy:
    tofu destroy -var="location={{region}}" -auto-approve

status:
    kubectl get pods -A 
    kubectl get pvc -n lgtm

argocd:
    kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
    kubectl get svc -n argocd | grep argocd-server

argocd-password:
    kubectl -n argocd get secred argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d 