import React, { Component } from 'react'

class TimerSection extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "A default session name",
      secondsElapsed: 0,
      isStarted: false,
      isStopped: false,
      lastClearedIncrementer: null
    };
    this.incrementer = null;
  }

  handleTextChange(e) {
    e.preventDefault()
    this.setState({name: e.target.value})
  }

  handleStartClick() {
    this.setState({isStopped: false, isStarted: true})
    this.incrementer = setInterval( () =>
      this.setState({
        secondsElapsed: this.state.secondsElapsed + 1
      })
    , 1000);
  }

  handleStopClick() {
    this.setState({isStopped: true})
    clearInterval(this.incrementer);
    this.setState({
      lastClearedIncrementer: this.incrementer
    });
  }

  handleResetClick() {
    this.setState({
      isStarted: false,
      isStopped: false
    })
    clearInterval(this.incrementer);
    this.setState({
      secondsElapsed: 0
    });
  }

  handleSaveSession() {
    const { saveSession } = this.props;
    const { name, secondsElapsed } = this.state;
    let session = {
      name: name,
      duration: secondsElapsed
    }
    saveSession(session)
    this.handleResetClick()
  }

  render() {
    let {isStarted, isStopped} = this.state;
    const formattedSeconds = this.props.formatSeconds(this.state.secondsElapsed)
    return (
      <div className="timer-section">
        <div className="session-name">
          <textarea onChange={this.handleTextChange.bind(this)} placeholder="Your Session Name" defaultValue={this.state.name} wrap></textarea>
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
    formatSeconds: React.PropTypes.func.isRequired
}

export default TimerSection