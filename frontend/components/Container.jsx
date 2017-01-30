import React from 'react';

export class Container extends React.Component {
    render() {
        let children = null;
        if (this.props.children) {
            children = React.cloneElement(this.props.children, {})
        }

        return (
            <div>
                {children}
            </div>
        )
    }
}

export default Container;
