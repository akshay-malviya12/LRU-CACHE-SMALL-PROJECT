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


### ReactJS for Front-end.
### LRU Cache Capacity Configuration:
In the LRU cache settings, you can specify the capacity by entering a number (integer). For instance, if you wish to set the capacity to 10,
simply enter "10" in the designated section. You can set the capacity to any desired number. Once a capacity is set, you should not exceed 
the specified limit when adding key-value pairs to the LRU cache. If you attempt to add more than the specified capacity, an error message 
indicating that the capacity is full will be displayed.

### Default Expiry Time:
By default, the expiry time for each key-value pair is set to 5 seconds. After this time period elapses, the key-value pair is automatically 
removed from the cache. 

### Functionality Overview:
1) Set Key-Value Pair:
Enter the key and value in the input boxes provided. After adding the key-value pair, you will receive a success message confirming that it 
has been added to the cache.
2)Get Value for a Key:
Enter the key in the input box designated for retrieving key-value pairs. Once the key is entered, the corresponding value will be displayed 
in the table.
3)Delete Key-Value Pair:
Enter the key in the input box provided for deletion of the key-value pair. Upon deletion, you will receive a confirmation message.

### GoLang backend 
### LRU Cache Implementation in Go
### The LRU Cache implementation in Go provides a simple and efficient way to store key-value pairs with an expiration time limit and a maximum capacity.

### LRU Cache is a data structure that stores a fixed number of items, removing the least recently used item when the cache is full (after 5 seconds, as I have set the time limit).

### LRU Cache is useful for caching frequently accessed data, particularly data that is frequently accessed but not frequently updated.

### Firstly, I've implemented a function to set data in the LRU Cache with a key, value, and expiration time. This function adds data to the LRU Cache dynamically, and I've tested it using Curl requests, which is working perfectly.

### Next, I've implemented a GetCache function to retrieve a value from the cache by its key, with the key obtained dynamically from the URL. This function is also working correctly.

### Then, I've created a DeleteCache function to delete cache entries by their key, with the key obtained dynamically from the URL. This function is functioning properly as well.

###Following that, I've implemented a GetAllKeys function to retrieve all keys and values from the LRU cache. This function provides a way to check the total number of keys available in the LRU Cache at a given time, ensuring the correctness of the code logic. This function is also working correctly.

### Additionally, I've created a frontend using React.js and a backend using Golang. Requests are sent to the Golang backend, and responses are received in JSON format, providing a fully functional API.

### This setup allows for efficient communication between the frontend and backend, ensuring smooth data retrieval and manipulation.
