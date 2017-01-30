import React, {Component, PropTypes} from 'react'
import Dropdown from 'react-dropdown'
import Table from 'rc-table'

const options = ['Day', 'Week', 'Month']
const months = ['January', 'February', 'March', 'April', 'May', 'June', 'August', 'September', 'November', 'December']

class SessionsSection extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selected: options[0],
            columns: [
                {
                    title: 'Name',
                    dataIndex: 'name',
                    key: 'name',
                    width: 250
                }, {
                    title: "Duration",
                    dataIndex: 'duration',
                    key: 'duration',
                    width: 100
                }, {
                    title: "Date",
                    dataIndex: 'createdAt',
                    key: 'createdAt',
                    width: 100
                }
            ]
        }
    }
    onSelect(selection) {
        this
            .props
            .setSessionsSection(selection.value)
    }
    formatDuration(sessions) {
        for (let session of sessions) {
            session.duration = this
                .props
                .formatSeconds(session.duration)
            session.createdAt = this.formatDate(session.createdAt)
        }
    }
    formatDate(date) {
        let d = new Date(date),
            month = '' + (d.getMonth() + 1),
            day = '' + d.getDate(),
            year = d.getFullYear();

        if (day.length < 2)
            day = '0' + day;

        const currentMonthIndex = month - 1
        return months[currentMonthIndex] + ', ' + day
    }
    render() {
        const className = "sessions-section " + this.props.customClassName;
        const sessions = this.formatDuration(this.props.sessions)
        return (
            <div className={className}>
                <div className="sessions-label">
                    <h2>My Sessions</h2>
                    <Dropdown
                        className="time-dropdown"
                        options={options}
                        onChange={this
                        .onSelect
                        .bind(this)}
                        value={this.props.selectedPeriod}
                        placeholder="Select an option"/>
                </div>
                <div className="sessions-list">
                    <Table
                        data={this.props.sessions}
                        columns={this.state.columns}
                        rowKey="id"
                        className="text-left electron-sessions"/>
                </div>
            </div>
        )
    }
}

SessionsSection.propTypes = {
    sessions: React.PropTypes.array.isRequired,
    formatSeconds: React.PropTypes.func.isRequired,
    setSessionsSection: React.PropTypes.func.isRequired,
    selectedPeriod: React.PropTypes.string.isRequired,
    customClassName: React.PropTypes.string
}

export default SessionsSection