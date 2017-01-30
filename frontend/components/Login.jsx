import React, {Component} from 'react'
import GoogleLogin from 'react-google-login'
import 'whatwg-fetch'

const HOST = "http://138.197.15.62:4000"
const ENDPOINT_AUTH = "/auth"
const ENDPOINT_USER = "/user"
const CLIENT_DOMAIN = "http://thetimetracker.surge.sh"

class Login extends Component {
    constructor(props) {
        super(props)
        this.state = {
            isError: false,
        }
    }
    onSuccess(responseGoogle) {
        this.authenticate(responseGoogle)
    }
    onFailure() {
        this.setState({isError: true})
    }
    authenticate(responseGoogle) {
        const code = responseGoogle.code
        fetch(HOST + ENDPOINT_AUTH + '?code=' + code, {
            headers: {
                "Content-type": "application/x-www-form-urlencoded; charset=UTF-8"
            },
        })
        .then(this.handleFetchErrors)
        .then(r => r.json())
        .then(response => response = this.fetchUserInfo(response.AccessToken))
        .catch(this.setState({isError: true}))
    }
    fetchUserInfo(token) {
        fetch('https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=' + token)
        .then(this.handleFetchErrors)
        .then(r => r.json())
        .then(response => response = this.newUser(response.email, token))
        .catch(this.setState({isError: true}))
    }
    newUser(email, token) {
        const isError = this.state.isError
        const requestBody = "email=" + email + "&accessToken=" + token
        fetch(HOST + ENDPOINT_USER, {
            method: 'POST',
            headers: {
                "Content-type": "application/x-www-form-urlencoded; charset=UTF-8"
            },
            body: requestBody,
        })
        .then(this.handleFetchErrors)
        .then(response => response = this.setAuthenticatedUser(email, token))
        .catch(this.setState({isError: true}))
    }
    setAuthenticatedUser(email, token) {
        localStorage.setItem('time_tracker_token', token)
        localStorage.setItem('time_tracker_authenticated', true)
        localStorage.setItem("time_tracker_user_email", email)
        window.open(CLIENT_DOMAIN, '_self')
    }
    handleFetchErrors(response) {
        if (!response.ok) {
            throw Error(response.statusText);
        }
        return response;
    }
    render () {
        const isError = this.state.isError
        return (
            <div>
                <div className="login"></div>
                <GoogleLogin
                    className="google-login"
                    clientId="612927008159-ftu0ijkhk41a8coiil2psvcksei1r49h.apps.googleusercontent.com"
                    buttonText="Login with Google"
                    scope="profile email https://www.googleapis.com/auth/calendar"
                    offline={true}
                    onSuccess={this.onSuccess.bind(this)}
                    onFailure={this.onFailure.bind(this)}
                />
                {isError && <div className="auth-error">
                    <span>Error while authenticating :(</span>
                </div>}
            </div>
        )
    }
}

Login.PropTypes = {
    isAuthenticated: React.PropTypes.bool.isRequired
}

export default Login
