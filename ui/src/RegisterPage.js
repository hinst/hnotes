import React from 'react';
import "./w3.css";

class RegisterPage extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			captchaId: "",
		};
		this.requestCaptcha();
	}

	render() {
		return(
			<div className="w3-container">
				<h2>Register</h2>
				<label>Username:</label>
				<input className="w3-input w3-border" type="text"/>
				<div style={{height: 4}}/>
				<label>Password:</label>
				<input className="w3-input w3-border" type="password"/>
				<div style={{height: 4}}/>
				<label>Retype password:</label>
				<input className="w3-input w3-border" type="password"/>
				<div style={{height: 8}}/>
				<img src={
					(this.state.captchaId !== "")
					? (this.props.serverURL + "/captcha/" + this.state.captchaId + ".png")
					: null
				} alt="captcha"/>
				<div style={{height: 4}}/>
				<label>Text from image:</label>
				<input className="w3-input w3-border" type="text"/>
				<div style={{height: 8}}/>
				<button className="w3-button w3-round w3-border">Register</button>
			</div>
		);
	}

	requestCaptcha() {
		fetch(this.props.serverURL + "/getCaptcha").then((response) => this.receiveCaptchaResponse(response));
	}

	receiveCaptchaResponse(response) {
		response.text().then((text) => this.receiveCaptchaText(text));
	}

	receiveCaptchaText(text) {
		console.log("'" + text + "'");
		this.setState({
			captchaId: text,
		});
	}

}

export default RegisterPage;