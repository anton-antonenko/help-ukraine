# help-ukraine

This is a simple wrapper for Docker `alpine/bombardier` image with simple REST Api.   
APIs:   
- `POST localhost:8080/job {"values": ["utl1", "url2"]}` - starts a load on specified urls  
- `GET localhost:8080/job` - shows current status  
- `DELETE localhost:8080/job` - deletes current load  
