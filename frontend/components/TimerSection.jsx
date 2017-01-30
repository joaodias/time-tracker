import React, { Component } from 'react'

class TimerSection extends Component {
  constructor(props) {
    super(props);
    // FIXME: This state should be moved to the top component.
    this.state = {
      name: localStorage.getItem("time_tracker_name"),
      secondsElapsed: this.checkLocalStorageNanDuration(localStorage.getItem("time_tracker_duration")),
      initialTimestamp: localStorage.getItem("time_tracker_initial_ts"),
      isStarted: false,
      isStopped: false,
      lastClearedIncrementer: null
    };
    this.incrementer = null
  }

  checkLocalStorageNanDuration(duration) {
    const integerDuration = parseInt(duration)
    if(isNaN(integerDuration)) {
      return 0
    } else {
      return integerDuration
    }
  }

  handleTextChange(e) {
    e.preventDefault()
    this.setState({name: e.target.value}, () => {
      localStorage.setItem("time_tracker_name", this.state.name);
    })
  }

  handleStartClick() {
    if (!this.state.isStopped && !this.state.isStarted) {
      const initialTimestamp = new Date().toISOString()
      this.setState({initialTimestamp: initialTimestamp})
      localStorage.setItem("time_tracker_initial_ts", this.state.initialTimestamp);
    }
    if(isNaN(this.state.secondsElapsed)) {
      this.setState({secondsElapsed: 0})
    }
    this.setState({isStopped: false, isStarted: true})
    this.incrementer = setInterval( () => this.incrementCount(), 1000)
  }

  incrementCount() {
    console.log(this.state.secondsElapsed)
    this.setState({secondsElapsed: this.state.secondsElapsed + 1});
    localStorage.setItem("time_tracker_duration", this.state.secondsElapsed)
  }

  handleStopClick() {
    this.setState({isStopped: true})
    clearInterval(this.incrementer)
    this.setState({
      lastClearedIncrementer: this.incrementer
    });
  }

  handleResetClick() {
    this.setState({
      isStarted: false,
      isStopped: false,
      name: '',
    })
    clearInterval(this.incrementer)
    this.setState({
      secondsElapsed: 0
    });
    localStorage.setItem("time_tracker_duration", 0);
    localStorage.setItem("time_tracker_name", '');
    localStorage.setItem("time_tracker_initial_ts", 0);
  }

  setWantCalendar(wantCalendar) {
    this.props.setWantCalendar(wantCalendar)
  }

  handleSaveSession() {
    const { saveSession } = this.props;
    const { name, secondsElapsed, initialTimestamp } = this.state;
    if(name != "") {
      let session = {
        name: name,
        duration: secondsElapsed,
        initialTimestamp: initialTimestamp,
      }
      saveSession(session)
      this.handleResetClick()
    } else {
      alert("Give them tasks a name. They diserves it!")
    }
  }

  render() {
    let {isStarted, isStopped} = this.state;
    const formattedSeconds = this.props.formatSeconds(this.state.secondsElapsed)
    const wantCalendar = this.props.wantCalendar
    return (
      <div className="timer-section">
        <div className="gcalendar" onClick={this.setWantCalendar.bind(this, !wantCalendar)}>
          {wantCalendar ? <div>
              <img src="../images/google-active.svg"/>
              <span>Save to google calendar</span>
            </div>
             : <div>
              <img src="../images/google-inactive.svg"/>
              <span>Do not save to google calendar</span>
             </div>}
        </div>
        <div className="session-name">
          <textarea onChange={this.handleTextChange.bind(this)} placeholder="Task name goes here" defaultValue={this.state.name} wrap></textarea>
        </div>
        <div className="stopwatch">
          <h1 className="stopwatch-timer">{formattedSeconds}</h1>
        </div>
        <div className="controls">
            {(!isStarted && !isStopped) || isStarted && isStopped ? <button type="button" onClick={this.handleStartClick.bind(this)} className={"btn " + "start-btn"}>start</button> : <button type="button" disabled onClick={this.handleStartClick.bind(this)} className={"btn " + "start-btn"}>start</button>}

            {isStarted && !isStopped ? <button type="button" onClick={this.handleStopClick.bind(this)} className={"btn " + "stop-btn"}>stop</button> : <button type="button" disabled onClick={this.handleStopClick.bind(this)} className={"btn " + "stop-btn"}>stop</button>}

            {isStarted && isStopped ? <button type="button" onClick={this.handleResetClick.bind(this)} className={"btn " + "reset-btn"}>reset</button> : <button type="button" disabled onClick={this.handleResetClick.bind(this)} className={"btn " + "reset-btn"}>reset</button>}

            {isStarted && isStopped ? <button type="button" onClick={this.handleSaveSession.bind(this)} className={"btn " + "save-btn"}>save</button> : <button type="button" disabled onClick={this.handleSaveSession.bind(this)} className={"btn " + "save-btn"}>save</button>}
        </div>
      </div>
    );
  }
}

TimerSection.propTypes = {
    saveSession: React.PropTypes.func.isRequired,
    formatSeconds: React.PropTypes.func.isRequired,
    setWantCalendar: React.PropTypes.func.isRequired,
    wantCalendar: React.PropTypes.bool.isRequired
}

export default TimerSection
