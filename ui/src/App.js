import React, { Component } from 'react';
import "./w3.css";
import NoteList from "./NoteList";
import RegisterPage from "./RegisterPage";

class App extends Component {

	constructor() {
		super();
		this.state = {
			notes: [],
			activePage: "register",
		}
	}

	serverURL = "http://localhost:9001/hnotes";

	renderNotes() {
		return <NoteList notes={this.state.notes}/>
	}

	renderActivePage() {
		if (this.activePage == "")
			return (<div></div>);
		else if (this.activePage = "register")
			return (<RegisterPage/>);
	}

	render() {
		return (
			<div className="w3-container">
				<div className="w3-container w3-brown">
					<h1><a href="#" onClick={() => this.showNotes()}>hnotes</a></h1>
				</div>
				{this.renderActivePage()}
			</div>
		);
	}

	showNotes() {
		let url = this.serverURL + "/notes";
		fetch(url).then((response) => {this.receiveNotes(response);});
	}

	receiveNotes(response) {
		response.json().then(
			(data) => {
				console.log(data);
				this.setState({
					notes: data,
				});
			}
		);
	}

}

export default App;
