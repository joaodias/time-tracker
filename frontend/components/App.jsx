import React, {Component} from 'react'
import Sidebar from './Sidebar.jsx'
import MainSection from './MainSection.jsx'
import 'whatwg-fetch'

// This should become environment variables
const CLIENT_DOMAIN = "http://thetimetracker.surge.sh"
const HOST = 'http://138.197.15.62:4000'
const ENDPOINT = '/timesession'

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            sessions: [],
            selectedPeriod: 'Day',
            isTimerVisible: true,
            isTasksVisible: false,
            wantCalendar: true,
        }
    }

    componentDidMount() {
        this.listSessions()
    }

    saveSession(session) {
        const requestBody = 'name=' + session.name + '&duration=' + session.duration + '&initialTimestamp=' + session.initialTimestamp + '&gCalendar=' + this.state.wantCalendar + '&userId=' + localStorage.getItem("time_tracker_user_email")
        console.log("Save with ", requestBody)
        fetch(HOST + ENDPOINT, {
            method: 'POST',
            headers: {
                "Content-type": "application/x-www-form-urlencoded; charset=UTF-8"
            },
            body: requestBody,
        })
        .then(this.handleFetchErrors)
        .then(response => this.listSessions())
        .catch(this.signOut())
    }

    signOut() {
        // For now when logging out there is no complete clearance of the
        // localstrorage. But there should be. Every data should be wiped on
        // logout. Since refresh tokens are not yet set. Session data should be
        // preserved
        localStorage.setItem('time_tracker_authenticated', false)
        window.open(CLIENT_DOMAIN, '_self')
    }

    setSessionsSection(selectedPeriod) {
        this.setState({selectedPeriod}, () => {
            this.listSessions();
        });
    }

    listSessions() {
        let sessions = [];
        fetch(HOST + ENDPOINT + '?period=' + this.state.selectedPeriod + '&userId=' + localStorage.getItem("time_tracker_user_email"))
        .then(this.handleFetchErrors)
        .then(response => response.json())
        .then(sessions => sessions = this.setState({sessions}))
        .catch(e => alert("Error listing time sessions"))
    }

    formatSeconds(sec) {
        let hours   = Math.floor(sec / 3600);
        let minutes = Math.floor((sec - (hours * 3600)) / 60);
        let seconds = sec - (hours * 3600) - (minutes * 60);

        if(minutes === 0 && hours === 0) {
            return seconds + 's'
        } else if (hours === 0 && minutes != 0) {
            return minutes + 'm ' + seconds + 's'
        } else {
            return hours + 'h ' + minutes + 'm ' + seconds + 's'
        }
    }

    setVisibleMainSection(section) {
        if (section === "timer") {
            this.setState({isTimerVisible: true});
            this.setState({isTasksVisible: false});
        } else if(section === "tasks") {
            this.setState({isTimerVisible: false});
            this.setState({isTasksVisible: true});
        }
    }

    setWantCalendar(wantCalendar) {
        this.setState({wantCalendar})
    }

    handleFetchErrors(response) {
        if (!response.ok) {
            throw Error(response.statusText);
        }
        return response;
    }

    render() {
        return (
            <div>
                <Sidebar
                sessions={this.state.sessions}
                formatSeconds={this.formatSeconds.bind(this)}
                setSessionsSection={this.setSessionsSection.bind(this)}
                selectedPeriod={this.state.selectedPeriod}
                setVisibleMainSection={this.setVisibleMainSection.bind(this)}
                />
                <MainSection
                saveSession={this.saveSession.bind(this)}
                formatSeconds={this.formatSeconds.bind(this)}
                isTasksVisible={this.state.isTasksVisible}
                isTimerVisible={this.state.isTimerVisible}
                sessions={this.state.sessions}
                formatSeconds={this.formatSeconds.bind(this)}
                setSessionsSection={this.setSessionsSection.bind(this)}
                selectedPeriod={this.state.selectedPeriod}
                wantCalendar={this.state.wantCalendar}
                setWantCalendar={this.setWantCalendar.bind(this)}/>
            </div>
        )
    }
}

export default App
