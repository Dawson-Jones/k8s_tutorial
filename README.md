kubectl apply -f nginx.yaml 

kubectl get pods

kubectl port-forward nginx-pod 4000:80

curl http://127.0.0.1:4000



kubectl exec -it nginx-pod -- /bin/bash
exec 进入 namespace 中, 和下面命令基本一样, 除了能够进入远程的 ns
```bash
[dawsonjia@VM-112-138-tencentos my_project]$ sudo nsenter -a -t 2278397
mesg: ttyname failed: No such file or directory
root@nginx-pod:/# ls
bin  boot  dev	etc  home  lib	lib64  media  mnt  opt	proc  root  run  sbin  srv  sys  tmp  usr  var
root@nginx-pod:/# cd /usr/share/nginx/
root@nginx-pod:/usr/share/nginx# ls
html
root@nginx-pod:/usr/share/nginx# cd html/
root@nginx-pod:/usr/share/nginx/html# ls
50x.html  index.html
```

echo "hello kubernetes by nginx!" > /usr/share/nginx/html/index.html

[dawsonjia:~/my_project]$ curl http://127.0.0.1:4000
hello kubernetes by nginx!


kubectl logs --follow nginx-pod
kubectl exec nginx-pod -- ls

kubectl delete pod nginx-pod

kubectl delete -f nginx.yaml

kubectl rollout undo deployment hellok8s-deployment
kubectl rollout history deployment hellok8s-deployment
kubectl rollout undo deployment/hellok8s-deployment --to-revision=2