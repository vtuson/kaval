# kaval

This is a small tool that validates kubeapps deployment. It checks for:
* pod status in kubeless, kube-system and kubeapps namespackes
* endpoint status in kubeless, kube-system and kubeapps namespackes
* reaches out to the three main api endpoints externall to validate that ingress and services are working (checks for a response code below 400)


```
./kaval -url http://192.168.99.100
kubeconfig path is: ~/.kube/config
## PASS:kubeless pod status: OK
## PASS:kubeless endpoints status: OK
## PASS:kubeapps pod status: OK
## PASS:kubeapps endpoints status: OK
## PASS:kube-system pod status: OK
## PASS:kube-system endpoints status: OK
## PASS:http://192.168.99.100/ pass with response 200
## PASS:http://192.168.99.100/api/v1/repos pass with response 200
## PASS:http://192.168.99.100/kubeless pass with response 200
----------------------------------------
		test pass: 9
		test fail: 0

		Overall: PASS
----------------------------------------
```
