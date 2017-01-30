import React, {Component} from 'react';
import SessionsSection from './SessionsSection.jsx';
import SidebarMenu from './SidebarMenu.jsx'

class Sidebar extends Component {
    render() {
        return (
            <div className="sidebar">
                <SessionsSection {...this.props} customClassName="sessions-section-sidebar"></SessionsSection>
                <SidebarMenu
                setVisibleMainSection={this.props.setVisibleMainSection}></SidebarMenu>
            </div>
        );
    }
}

Sidebar.propTypes = {
    sessions: React.PropTypes.array.isRequired,
    formatSeconds: React.PropTypes.func.isRequired,
    setSessionsSection: React.PropTypes.func.isRequired,
    selectedPeriod: React.PropTypes.string.isRequired,
    setVisibleMainSection: React.PropTypes.func.isRequired
}

export default Sidebar;