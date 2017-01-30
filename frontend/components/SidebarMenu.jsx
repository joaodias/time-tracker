import React, {Component} from 'react';

class SidebarMenu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            activeMenu: "timer",
        }
    }
    onClick(id) {
        let activeMenu;
        if(id === "timer") {
            activeMenu = "timer";
        } else {
            activeMenu = "tasks";
        }
        this.setState({activeMenu: activeMenu});
        this.props.setVisibleMainSection(activeMenu);
    }
    render() {
        const timerMenuState = this.state.activeMenu === "timer" ? "active-menu" : "inactive-menu";
        const tasksMenuState = this.state.activeMenu === "tasks" ? "active-menu" : "inactive-menu";
        return (
            <div className="sidebar-menu">
                <ul className="buttons-list">
                    <li className="menu-buttons">
                        <button className={timerMenuState} onClick={this.onClick.bind(this, "timer")}>Timer</button>
                    </li>
                    <li className="menu-buttons">
                        <button className={tasksMenuState} onClick={this.onClick.bind(this, "tasks")}>Tasks</button>
                    </li>
                    <li className="image-button"></li>
                </ul>
            </div>
        );
    }
}

SidebarMenu.propTypes = {
    setVisibleMainSection: React.PropTypes.func.isRequired
}

export default SidebarMenu;