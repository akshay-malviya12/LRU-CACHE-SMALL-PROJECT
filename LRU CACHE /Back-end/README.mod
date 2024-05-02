### ReactJs For Frontend and backend Golang.


### if you want to testing backend for curl request you should be use this curls :-
(Step by steps)

### Set for LRU Cache capacity (if you do not set capacity show error, so first of all set capacity. )
1) curl -X POST -H "Content-Type: application/json" http://localhost:8080/setcapacity?capacity=5


### Set Key-value Pair in LRU Cache.
2)curl -X POST -H "Content-Type: application/json" -d '{"key":"1","value":"AB"}' http://localhost:8080/setcache

### Get Key-Value pair Provided by specific Key.
3)curl -X POST -H "Content-Type: application/json" http://localhost:8080/getcache?key=1

### Delete Key-value pair itself use this curl ,provide any key if you want to delete.
4)curl -X POST -H "Content-Type: application/json" http://localhost:8080/deletecache?key=1   


### This is extra work (out of the requirement).
### Get All key-value pair .
5)curl -X POST -H "Content-Type: application/json" http://localhost:8080/GetAllLRUCases

### ReactJS For Front-end.
### ReactJs - I have create set capacity for LRU cache set value in number(Integer).If you want to set capacity in LRU cache then you have to set capacity in set capacity section. For Example if you want to set capacity as 10 then you have to set capacity as 10 in set capacity section.you can set capacity as any number.if you set capacity you don't add more than 10 key-value pair in LRU cache.if you add more than 10 key-value pair show error message capacity is full. I have set default expire time as 5 seconds.after 5 seconds auto delete key ,than you enter new key-value pair and you can see message added to LRU cache successfully.

### I have create set key and value in input box for set key-value pair.than you enter key and value than you can see message added to cache successfully.
### I have create get key in input box for get key pair.than you enter key ,key related value show in the table.
### I have create user delete key in input box for delete key-value pair.

GoLang BackEnd 
### This LRU Cache implmentation in Go provide  a simple and efficient way to store a LRU cache in key- value pair and expire time limit with maximum capacity.
### The LRU Cache is a data structure that stores a fixed number of items, with the most recently used item being removed when the cache is full(after 5 seconds because i have set time limit).

### The LRUCache is useful for caching frequently accessed data, such as data that is frequently accessed but not frequently updated.

### First of all,i have set data in LRU Cache with key and value and expire time.This data added in LRU Cache dynamically and i tested this set function use of Curl request.This is completly working fine.


### Now,i have created a GetCache function which is used to get the value from the cache by key and key get dynamically from url.This is also working fine.


### Now,i have created a DelelteCache function which is used to delete the cache by key and key get dynamically from url.This is also working fine.


### Now,i have created a GetAllKeys function which is used to get all keys and values from the LRU cache. only for satisfaction(checking count Total LRU Cache availble on time accoding condition and code logic) how much key enter in LRU Cache.This is also working fine.

### than ,i have creaeted front-end side Reactjs and back-end side Golang.
### i have sent requets to back-end side Golang and get response in json format.This is completly working API Formate.
