import React from 'react';
import "./w3.css";

class RegisterPage extends React.Component {

	constructor() {
		super();
		this.state = {
			captchaImageURL = "",
		};
		requestCaptcha();
	}

	render() {
		return(
			<div className="w3-container">
				<h2>Register</h2>
				<label>Username</label>
				<input className="w3-input w3-border" type="text"/>
				<div style={{height: 4}}/>
				<label>Password</label>
				<input className="w3-input w3-border" type="password"/>
				<div style={{height: 4}}/>
				<label>Retype password</label>
				<input className="w3-input w3-border" type="password"/>
				<div style={{height: 8}}/>
				<button className="w3-button w3-round w3-border">Register</button>
			</div>
		);
	}

	requestCaptcha() {
		fetch(this.serverURL + "/getCaptcha");
	}

}

export default RegisterPage;