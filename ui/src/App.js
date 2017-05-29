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
			serverURL: "http://localhost:9001/hnotes",
		}
	}

	renderNotes() {
		return <NoteList notes={this.state.notes}/>
	}

	/*
	setActivePage(activePage) {
		if (this.state.activePage != activePage) {
			this.setState({
				activePage: activePage,
			});
			if (activePage == "register") {
			}
		}
	}
	*/

	renderActivePage() {
		if (this.state.activePage === "")
			return (<div></div>);
		else if (this.state.activePage === "register")
			return (<RegisterPage serverURL={this.state.serverURL}/>);
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
		const url = this.state.serverURL + "/notes";
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
