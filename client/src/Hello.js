import React, { Component } from 'react';
import axios from 'axios';

class Hello extends Component {
    state = {
        message: ''
    }

    componentDidMount() {
        this.fetchMessage(); 
    }

    async fetchMessage() {
        const message = await axios.get('http://localhost:1323/');
        console.log(message);
        this.setState({
            message: message.data
        });
    }

    render() {
        return (
            <div>{this.state.message}</div>
        );
    }
}

export default Hello;