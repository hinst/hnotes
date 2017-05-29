import React from 'react';
import "./w3.css";

class RegisterPage extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			userName: "",
			password: "",
			retypedPassword: "",
			captchaId: "",
			captcha: "",
		};
		this.requestCaptcha();
	}

	render() {
		return(
			<div className="w3-container">
				<h2>Register</h2>
				<label>Username:</label>
				<input className="w3-input w3-border" type="text" onChange={(e)=>this.handleUsernameChange(e)}/>
				<div style={{height: 4}}/>
				<label>Password:</label>
				<input 
					className="w3-input w3-border" 
					type="password"
					onChange={(event) => {
						this.setState({password: event.target.value});
					}}
				/>
				<div style={{height: 4}}/>
				<label>Retype password:</label>
				<input 
					className="w3-input w3-border" 
					type="password"
					onChange={(event) => {
						this.setState({retypedPassword: event.target.value});
					}}
				/>
				<div style={{height: 8}}/>
				<img src={
					(this.state.captchaId !== "")
					? (this.props.serverURL + "/captcha/" + this.state.captchaId + ".png")
					: null
				} alt="captcha"/>
				<div style={{height: 4}}/>
				<label>Text from image:</label>
				<input 
					className="w3-input w3-border" 
					type="text"
					onChange={(event) => {
						this.setState({captcha: event.target.value});
					}}
				/>
				<div style={{height: 8}}/>
				<button 
					className="w3-button w3-round w3-border" 
					onClick={()=>this.receiveRegisterButtonClick()}
				>
					Register
				</button>
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

	receiveRegisterButtonClick() {
		if (this.state.password === this.state.retypedPassword) {
			if (this.state.password.length > 0) {
				this.sendRegisterRequest();
			} else {
				alert("Password is empty");
			}
		} else {
			alert("Passwords do not match");
		}
	}

	handleUsernameChange(event) {
		this.setState({userName: event.target.value});
	}

	sendRegisterRequest() {
		const url = this.props.serverURL + "/registerNewUser";
		fetch(url, {
			method: "post",
			body: JSON.stringify({
				CaptchaId: this.state.captchaId,
				Captcha: this.state.captcha,
				User: this.state.userName,
				Password: this.state.password,
			}),
		}).then(
			(response) => response.json().then((data) => this.receiveRegisterResponse(data))
		);
	}

	receiveRegisterResponse(data) {
		console.log(data);
		if (data.CaptchaSuccess) {
			if (data.Success) {

			} else {
				this.requestCaptcha();
				alert("Could not register this username.");
			}
		} else {
			this.requestCaptcha();
			alert("Image verification code is incorrect.");
		}
	}

}

export default RegisterPage;