import React, { Component } from 'react'
import UserService from '../services/UserService';

class CreateUserComponent extends Component {
    constructor(props) {
        super(props);

        this.state = {
            id: this.props.match.params.id,
            name: '',
            email: ''
        }

        this.changeNameHandler = this.changeNameHandler.bind(this);
        this.changeEmailHandler = this.changeEmailHandler.bind(this);
        this.saveOrUpdateUser = this.saveOrUpdateUser.bind(this);
    }

    componentDidMount() {
        if (this.state.id === '_add') {
            return
        }
        else {
            UserService.getUserById(this.state.id).then( (res) => {
                let user = res.data;
                this.setState({
                    name: user.name,
                    email: user.email
                });
            });
        }
    }

    saveOrUpdateUser = (e) => {
        e.preventDefault();
        let user = {
            name: this.state.name,
            email: this.state.email
        }

        console.log('user => ' + JSON.stringify(user));

        if (this.state.id === '_add') {
            UserService.createUser(user).then(res => {
                this.props.history.push('/users');
            });
        }
        else {
            UserService.updateUser(user, this.state.id).then(res => {
                this.props.history.push('/users');
            });
        }
    }

    changeNameHandler = (e) => {
        this.setState({name: e.target.value});
    }

    changeEmailHandler = (e) => {
        this.setState({email: e.target.value});
    }

    cancel() {
        this.props.history.push('/users');
    }

    getTitle() {
        if (this.state.id === '_add') {
            return <h3 className="text-center">Add User</h3>
        }
        else {
            return <h3 className="text-center">Update User</h3>
        }
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label>Name:</label>
                                            <input placeholder="Name" name="Name" className="form-control" 
                                                value={this.state.name} onChange={this.changeNameHandler}/>
                                        </div>
                                        <div className = "form-group">
                                            <label>Email</label>
                                            <input placeholder="Email" name="email" className="form-control" 
                                                value={this.state.email} onChange={this.changeEmailHandler}/>
                                        </div>

                                        <button className="btn btn-success" onClick={this.saveOrUpdateUser}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default CreateUserComponent;