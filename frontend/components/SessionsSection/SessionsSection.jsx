import React, {Component, PropTypes} from 'react'
import Dropdown from 'react-dropdown'
import ReactList from 'react-list';

const options = ['Day', 'Week', 'Month']

class SessionsSection extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selected: options[0],
        }
    }
    onSelect(selection) {
        this.props.setSessionsSection(selection.value)
    }
    renderItem(index, key) {
        const name = this.props.sessions[index].name
        const duration = this.props.formatSeconds(this.props.sessions[index].duration)
        return <ul className="sessions-list-item" key={key}>
            <li>{name}</li>
            <li>{duration}</li>
        </ul>;
    }
    render() {
        return (
            <div className="sessions-section">
                <div className="sessions-label">
                    <h2>My Sessions</h2>
                    <Dropdown
                        className="time-dropdown"
                        options={options}
                        onChange={this.onSelect.bind(this)}
                        value={this.props.selectedPeriod}
                        placeholder="Select an option"/>
                </div>
                <div className="sessions-list">
                    <ReactList
                        itemRenderer={this.renderItem.bind(this)}
                        length={this.props.sessions.length}
                        type='uniform'/>
                </div>
            </div>
        )
    }
}

SessionsSection.propTypes = {
    sessions: React.PropTypes.array.isRequired,
    formatSeconds: React.PropTypes.func.isRequired,
    setSessionsSection: React.PropTypes.func.isRequired,
    selectedPeriod: React.PropTypes.string.isRequired
}

export default SessionsSection