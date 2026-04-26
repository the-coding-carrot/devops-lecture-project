region := "germanywestcentral"
rg := "web-shop"
cluster := "aks-cluster"

deploy:
    cd tf && tofu apply -var="location={{region}}" -auto-approve
    just connect
    kubectl create namespace argocd || true
    kubectl apply -n argocd --server-side --force-conflicts -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
    @echo "Warte 60 Sekunden auf ArgoCD..."
    just sync

connect:
    az aks get-credentials --resource-group {{rg}} --name {{cluster}} --overwrite-existing

sync:
    kubectl apply -f app/webshop.yml
    kubectl apply -f app/lgtm-stack.yml
    kubectl apply -f app/alloy-stack.yml

monitor:
    (sleep 3 && open http://localhost:3000) &
    kubectl port-forward svc/lgtm-stack-grafana 3000:80 -n lgtm
    
monitor-password:
    kubectl get secret lgtm-stack-grafana -n lgtm -o jsonpath="{.data.admin-password}" | base64 -d && echo

destroy:
    cd tf && tofu destroy -var="location={{region}}" -auto-approve

status:
    kubectl get pods -A 
    kubectl get pvc -n lgtm

argocd-ui:
    kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
    @echo "Warte auf die externe IP"
    kubectl wait --for=jsonpath='{.status.loadBalancer.ingress[0].ip}' service/argocd-server -n argocd --timeout=300s
    kubectl get svc argocd-server -n argocd -o jsonpath='{.status.loadBalancer.ingress[0].ip}'
    @echo ""

argocd-password:
    kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d && echo

webshop:
    just products &
    just auth &
    just checkout & 

products:
    kubectl port-forward svc/product-service 8081:80 -n webshop 

auth:
    kubectl port-forward svc/auth-service 8080:80 -n webshop

checkout:
    kubectl port-forward svc/checkout-service 8082:80 -n webshop

webshop-stop:
    -pkill -f "kubectl port-forward.*-n webshop"
