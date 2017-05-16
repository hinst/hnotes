import React, { Component } from 'react';
import logo from './logo.svg';
import "./w3.css";

class App extends Component {

  render() {
    return (
      <div className="w3-container">
        <div className="w3-container w3-brown">
          <h1>hnotes</h1>
          <button onClick={() => this.showNotes()} >show notes</button>
        </div>
      </div>
    );
  }

  showNotes() {
    let url = process.env.PUBLIC_URL + "/notes";
    console.log(url);
    fetch(url).then(this.receiveNotes);
  }

  receiveNotes(response) {
    console.log(response.json());
  }

}

export default App;
