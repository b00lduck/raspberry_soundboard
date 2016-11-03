import React from 'react';
import './Image.css';

export default class Sound extends React.Component {

    constructor(props) {
        super(props);
        this.state = props.data;
    }

    componentWillReceiveProps(props) {
        this.setState(props.data);
    }

    render() {
        return (
            <div className="Image">
                <img alt="mp3" src={"http://localhost:8080/api/image/" + this.state.ImageFile} />
            </div>
        )
    }
}
