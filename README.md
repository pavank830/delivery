# delivery application
## steps to build application
```bash
# commands shown below run at the root of the project directory
1. To build the delivery application code:
      make all
```
## steps to run application
```bash
# commands shown below run at the root of the project directory
1. ./bin/delivery -port 60010

2.to see the flags ,run it as below: 
  ./bin/delivery -h
   Usage of ./bin/delivery:
       -port string
            HTTP server port,default port 8080 (default "8080")
```

# brief about the application
 to find the best route with shortest possible time for delivery executive who has been tasked with 2 orders from 2 different restaurants at the same time to 2 different customers.

# overview on the approach
The delivery exec can not go to any customer without picking the customer order from the restaurant and an assumption is that delivery exec can pick order from another restaurant while carrying another order.So initially delivery exec has to go to a restaurant from there we calculate all possible valid paths and then find path which takes less time.

application runs as HTTP server with endpoint /api/bestroute [to find the best route] 


sample curl [/api/bestroute]
```bash
# request for /api/bestroute
curl --location --request POST 'http://127.0.0.1:60010/api/bestroute' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deliver_data": {
        "delivery_exec_info": {
            "id": "D",
            "speed": 20, //avg speed in km/hr
            "location": {
                "latitude": 12.848000,
                "longitude": 77.671002
            }
        },
        "delivery_info": [
            {
                "consumer": {
                    "id": "C1",
                    "location": {
                        "latitude": 12.910363,
                        "longitude": 77.627086
                    }
                },
                "restaurant": {
                    "id": "R1",
                    "location": {
                        "latitude": 12.855796,
                        "longitude": 77.664828
                    },
                    "prep_time": 30 // in mins
                }
            },
            {
                "consumer": {
                    "id": "C2",
                    "location": {
                        "latitude": 12.873284,
                        "longitude": 77.650928
                    }
                },
                "restaurant": {
                    "id": "R2",
                    "location": {
                        "latitude": 12.900526,
                        "longitude": 77.633151
                    },
                    "prep_time": 20 // in mins
                }
            }
        ]
    }
}'

# response for /api/bestroute
{"path":["D","R2","C2","R1","C1"],"total_time":1.02529432958654,"response_code":0,"response_description":""} 
// response --> total_time in hrs
// best path D -> R2 -> C2 -> R1 -> C1
```
