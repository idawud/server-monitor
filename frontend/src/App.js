import React from 'react';  

import { w3cwebsocket as W3CWebSocket } from "websocket";
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

const client = new W3CWebSocket('ws://127.0.0.1:8080/feedback');
 
class App extends React.Component  {  
  componentDidMount(){
    client.onopen = () => {
      console.log('WebSocket Client Connected');
      client.send("Hello server!");
    };


    client.onmessage = (message) => {
      let status = JSON.parse(message.data);

      let elem = document.createElement('div');
      let head = document.createElement('h4');
      head.textContent = status.timestamp;
      elem.appendChild(head);
  
      for(var k in status) {
          if (k !== "timestamp"){
            let entry = document.createElement('div');
            entry.textContent = k; 
            console.log(status)
            if ( status[k] === true){ 
                //entry.style.color = "green";
                entry.className = "alert alert-success";
                entry.role = "alert";
            }
            else if(status[k] === false){
              entry.className = "alert alert-danger";
              entry.role = "alert";
            }
            elem.appendChild(entry);
           // elem.appendChild(document.createElement("br"))
            
          }
      }
      document.getElementById('body').prepend(elem);
      //console.log(`[message] Data received from server: ${keys}`);
    };

    client.onclose = (event) => {
      console.log("Connecton closed: ", event);
    }
    client.onerror = (event) => {
      console.log("Connecton Error: ", event);
    }
  }  

render(){
  return (
    <div className="App"> 
      <div id="messages"></div>
        <div id="header">
          <h1>Api Enpoints Stream</h1>
        </div>
        <div id="body">
         
        </div>
    </div>
  );
  }
}

export default App;
