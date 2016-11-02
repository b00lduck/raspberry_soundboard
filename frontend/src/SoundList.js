import React from 'react';
import Websocket from 'react-websocket';
import './SoundList.css';

export default class SoundList extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            Sounds: []
        };
    }

    handleData(data) {
        this.setState(JSON.parse(data));
    }

    render() {
        return (
            <div className="SoundList">
                {this.state.Sounds.map(item => (
                    <div className="Sound" key={item.SoundFile}>
                        <img alt="mp3" src={"http://localhost:8080/api/image/" + item.ImageFile} />
                        <div>{item.Count}x</div>
                    </div>
                ))}
                <Websocket url='ws://localhost:8080/api/websocket' onMessage={this.handleData.bind(this)}/>
            </div>
        );
    }
}
