# Simple CherryPy REST Server
A small CherryPy server to simulate basic REST calls. Hopefully, the tools here can be extended to automate other things in the Hero project.

Prerequisites: openstack clients, account/project/user in Neutrino

To run, modify openrc and flavor.sh to fit your environment, then run following commands
```
    # ./run-init.sh
    # ./flavor.sh <neutrino_vip>
    # ./run-cherrypy.sh
```

List floating IPs of the CherryPy servers
```
    # source ./openrc
    # heat --insecure output-show hero-cherrypy fip
```

Wait for the VMs to start. It may take a while for the VMs to boot up.

To test CherryPy server

GET:
```
    # curl -s http://<floating_ip_of_cherrypy_server>/api/heros
    # curl -s http://<floating_ip_of_cherrypy_server>/api/heros/<id>
```

POST:
```
    # curl -s -d '{"name":"david", "title":"tiger"}' -H "Content-Type: application/json" http://<floating_ip_of_cherrypy_server>/api/heros
```


PUT:
```
    # curl -s  -d '{"title":"lion"}'-H "Content-Type: application/json" -X PUT http://<floating_ip_of_cherrypy_server>/api/heros/<id>
```

DELETE:
```
    curl -s -X DELETE http://<floating_ip_of_cherrypy_server>/api/heros/<id>
```

