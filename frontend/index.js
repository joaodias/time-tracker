import React from 'react'
import ReactDOM from 'react-dom'
import ReactRouter, { Router, Route, hashHistory, IndexRedirect } from 'react-router';
import App from './components/App.jsx'
import Login from './components/Login.jsx'
import Container from './components/Container.jsx'

let isAuthenticated = localStorage.getItem('time_tracker_authenticated')

const _requireAuth = (nextState, replace) => {
    console.log("CHECK REQUIRE", isAuthenticated)
    if (isAuthenticated === 'false') {
        replace({ pathname: '/login' });
    }
}

ReactDOM.render(
    <Router history={hashHistory}>
        <Route path="/" component={Container}>
            <IndexRedirect to="/app" />
            <Route path="/app" component={App} onEnter={_requireAuth} />
            <Route path="/login" component={Login} />
        </Route>
    </Router>,
    document.getElementById('app')
);
