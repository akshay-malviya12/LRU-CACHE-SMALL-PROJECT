// App.js
import React, { useState } from "react";
import axios from "axios";
import "./App.css";

const LRUcache = () => {
  const [key, setKey] = useState("");
  const [value, setValue] = useState("");
  const [message, setMessage] = useState("");
  const [capacity, setCapacity] = useState("");
  const [data, setData] = useState([]);
  const [searchKey, setSearchKey] = useState("");

  const handleSet = async () => {
    try {
      await axios.post(
        "http://localhost:8080/setcache",
        JSON.stringify({ key, value }),
      );

      setMessage("Item added to cache successfully");
    } catch (error) {
      setMessage("Error adding item to cache");
    }
  };

  const handleGet = async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/getcache?key=${searchKey}`,
      );
      console.log(response.data);
      json.Data = await response.data;
      setData(jsonData);
    } catch (error) {
      console.log(error.response.data);
    }
  };

  const handleDelete = async () => {
    try {
      await axios.delete(`http://localhost:8080/deletecache?key=${searchKey}`);
      setMessage("Item deleted from cache successfully");
    } catch (error) {
      setMessage("Error deleting item from cache");
    }
  };

  const handleSetCapacity = async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/setcapacity?capacity=${capacity}`,
      );
      console.log(response.data);
    } catch (error) {
      console.log(error.response.data);
    }
  };
  return (
    <div className="Box">
      <h2>LRU Cache</h2>
      <div className="setcapacity">
        <input
          className="searchcapacity"
          type="number"
          value={capacity}
          onChange={(e) => setCapacity(e.target.value)}
          placeholder="Capacity"
        />
        <button className="button" onClick={handleSetCapacity}>
          Set Capacity
        </button>
      </div>
      <div className="setkeyvalue">
        <input
          className="search"
          type="text"
          value={key}
          onChange={(e) => setKey(e.target.value)}
          placeholder="Key"
        />
        <input
          className="search"
          type="text"
          value={value}
          onChange={(e) => setValue(e.target.value)}
          placeholder="Value"
        />
        <br />
        <br />
        <button className="button btn" onClick={handleSet}>
          Set
        </button>

        <p>{message}</p>
      </div>
      <div className="Getkeyvalues">
        <input
          className="searchcapacity"
          type="text"
          value={searchKey}
          onChange={(e) => setSearchKey(e.target.value)}
          placeholder="Enter Key"
        />
        <button className="button" onClick={handleGet}>
          Get
        </button>
        <button className="button" onClick={handleDelete}>
          Delete
        </button>
        <table className="table">
          <thead>
            <tr>
              <th>Key</th>
              <th>Value</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>1</td>
              <td>2</td>
            </tr>
            {data.map((item, index) => {
              // return(
              <tr key={index}>
                <td>{item.key}</td>
                <td>{item.value}</td>
              </tr>;
              // )
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default LRUcache;




     