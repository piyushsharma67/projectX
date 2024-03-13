# get list of all deployments

deployments=$(kubectl get deployment -o name)

for name in deployments;do
    kubectl rollout restart "$name"
done