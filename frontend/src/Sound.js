import React from 'react';

export default class Sound extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            lastPlayed: new Date(),
            numPlayed: 42
        };
    }

    render() {
        return (
            <div className="Sound">
                <h1>This is Sound "{this.props.name}"</h1>
                <h2>Last played: {this.state.lastPlayed.toLocaleString()}</h2>
                <h3>Num played: {this.state.numPlayed}</h3>
            </div>
        );
    }
}
