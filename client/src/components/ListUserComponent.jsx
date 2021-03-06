import React, { Component } from 'react'
import UserService from '../services/UserService'

class ListUserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                users: []
        }
        this.addUser = this.addUser.bind(this);
        this.editUser = this.editUser.bind(this);
        this.deleteUser = this.deleteUser.bind(this);
    }

    async deleteUser(id){
        await UserService.deleteUser(id)
        this.setState({users: this.state.users.filter(user => user.ID !== id)});
    }

    viewUser(id){
        console.log('viewUser with id:' + id);
        this.props.history.push(`/view-user/${id}`);
    }
    
    editUser(id){
        this.props.history.push(`/add-user/${id}`);
    }

    componentDidMount(){
        UserService.getUsers().then((res) => {
            this.setState({ users: res.data});
        });
    }

    addUser(){
        this.props.history.push('/add-user/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">User List</h2>
                 <div className="row mt-5">
                        <table className="table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th>Name</th>
                                    <th>Email</th>
                                    <th>Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.users.map(
                                        user => 
                                        <tr key={user.ID}>
                                             <td>{user.name}</td>   
                                             <td>{user.email}</td>
                                             <td>
                                                 <button onClick={ () => this.editUser(user.ID)} className="btn btn-info">Update</button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteUser(user.ID)} className="btn btn-danger">Delete</button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewUser(user.ID)} className="btn btn-info">View</button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>
                 <div className="row">
                     <div className="d-grid gap-2 col-6 mx-auto">
                        <button className="btn btn-primary" onClick={this.addUser}>Add User</button>
                    </div>
                 </div>

            </div>
        )
    }
}

export default ListUserComponent