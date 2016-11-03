import React from 'react';

export default class Sound extends React.Component {

    constructor(props) {
        super(props);
        this.state = props.data;
    }

    componentWillReceiveProps(props) {
        this.setState(props.data);
    }

    play() {
        fetch('http://localhost:8080/api/play/' + this.state.SoundFile);
    }

    render() {
        return (
            <div className="Sound" onClick={this.play.bind(this)}>
                <img alt="mp3" src={"http://localhost:8080/api/image/" + this.state.ImageFile} />
                <div>{this.state.Count}x</div>
                <div>{Math.round(this.state.Temperature*100)/100}°</div>
            </div>
        )
    }
}
