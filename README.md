## myofficestorage
Test storage for my ofiice

### Usage examples
To test storage run `make run`, and do some curl requests:

* store data in storage
  ```bash
  > curl -v -X POST -H "Content-Type: application/json" -d 'data' localhost:8080/test
  ...
  HTTP/1.1 200 OK
  ...
  ```

* read stored data in storage
  ```bash
  > curl localhost:8080/test
  data
  ```

* read non-existing key
  ```bash
  > curl -v localhost:8080/foo
  ...
  HTTP/1.1 404 Not Found
  ...
  ```
  
* delete value
  ```bash
  curl -v -X POST -H "Content-Type: application/json" -d '' localhost:8080/test
  ...
  HTTP/1.1 200 OK
  ...
  > curl -v localhost:8080/test
  ...
  HTTP/1.1 404 Not Found
  ...
  ```