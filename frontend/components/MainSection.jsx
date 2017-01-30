import React, {Component} from 'react';
import TimerSection from './TimerSection.jsx';
import SessionSection from './SessionsSection.jsx';

class MainSection extends Component {
    render() {
        return (
            <div className="main-section">
                {this.props.isTimerVisible ? <TimerSection {...this.props}></TimerSection> : null}
                {this.props.isTasksVisible ? <SessionSection {...this.props} customClassName="sessions-section-main"></SessionSection> : null}
            </div>
        );
    }
}

MainSection.propTypes = {
    sessions: React.PropTypes.array.isRequired,
    setSessionsSection: React.PropTypes.func.isRequired,
    selectedPeriod: React.PropTypes.string.isRequired,
    saveSession: React.PropTypes.func.isRequired,
    formatSeconds: React.PropTypes.func.isRequired,
    isTimerVisible: React.PropTypes.bool.isRequired,
    isTasksVisible: React.PropTypes.bool.isRequired,
    setWantCalendar: React.PropTypes.func.isRequired,
    wantCalendar: React.PropTypes.bool.isRequired
}

export default MainSection;
