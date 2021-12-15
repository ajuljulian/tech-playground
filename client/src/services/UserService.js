import axios from 'axios';

const USER_API_BASE_URL = "/api/users";

class UserService {

    getUsers() {
        return axios.get(USER_API_BASE_URL);
    }

    createUser(user) {
        return axios.post(USER_API_BASE_URL, `name=${user.name}&email=${user.email}`);
    }

    getUserById(userId) {
        return axios.get(USER_API_BASE_URL + '/' + userId);
    }

    updateUser(user, userId) {
        return axios.put(USER_API_BASE_URL + '/' + userId, `name=${user.name}&email=${user.email}`);
    }

    deleteUser(userId) {
        return axios.delete(USER_API_BASE_URL + '/' + userId);
    }
}

export default new UserService();