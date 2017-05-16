import React, { Component } from 'react';
import "./w3.css";
import NoteList from "./NoteList";

class App extends Component {

  serverURL = "http://localhost:9001/hnotes";

  renderNotes() {
    return <NoteList/>
  }

  render() {
    return (
      <div className="w3-container">
        <div className="w3-container w3-brown">
          <h1>hnotes</h1>
          <button onClick={() => this.showNotes()} >show notes</button>
        </div>
        {this.renderNotes()}
      </div>
    );
  }

  showNotes() {
    let url = this.serverURL + "/notes";
    fetch(url).then(this.receiveNotes);
  }

  receiveNotes(response) {
    response.json().then((data)=>{console.log(data)});
  }

}

export default App;
