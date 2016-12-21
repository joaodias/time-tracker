import React, {Component} from 'react'
import SessionsSection from './SessionsSection/SessionsSection.jsx'
import TimerSection from './TimerSection/TimerSection.jsx'
import 'whatwg-fetch'

// This should become environment variables
const HOST = 'http://127.0.0.1:4000'
const ENDPOINT = '/timesession'

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            currentSession: {
                name: '',
                duration: 0
            },
            sessions: [],
            selectedPeriod: 'Day'
        }
    }

    componentDidMount() {
        this.listSessions()
    }

    saveSession(session) {
        fetch(HOST + ENDPOINT, {
            method: 'POST',
            headers: {
                "Content-type": "application/x-www-form-urlencoded; charset=UTF-8"
            },
            body: 'name=' + session.name + '&duration=' + session.duration
        })
            .then(this.listSessions())
            .catch(e => alert("Error saving time session"))
    }

    setSessionsSection(selectedPeriod) {
        this.setState({selectedPeriod}, () => {
            this.listSessions();
        });
    }

    listSessions() {
        let sessions = [];
        fetch(HOST + ENDPOINT + '?period=' + this.state.selectedPeriod)
            .then(r => r.json())
            .then(sessions => sessions = this.setState({sessions}))
            .catch(e => alert("Error listing time sessions"))
    }

    formatSeconds(sec) {
        let hours   = Math.floor(sec / 3600);
        let minutes = Math.floor((sec - (hours * 3600)) / 60);
        let seconds = sec - (hours * 3600) - (minutes * 60);
        return hours + ':' + minutes + ':' + seconds
    }

    render() {
        return (
            <div>
                <SessionsSection
                sessions={this.state.sessions}
                formatSeconds={this.formatSeconds.bind(this)}
                setSessionsSection={this.setSessionsSection.bind(this)}
                selectedPeriod={this.state.selectedPeriod}/>
                <TimerSection
                saveSession={this.saveSession.bind(this)}
                formatSeconds={this.formatSeconds.bind(this)}/>
            </div>
        )
    }
}

export default App
