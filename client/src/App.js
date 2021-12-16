import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import ListUserComponent from './components/ListUserComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import CreateUserComponent from './components/CreateUserComponent';
import ViewUserComponent from './components/ViewUserComponent';

function App() {
  return (
    <div>
      <Router>
        <div className="container">
          <div class="row">
            <HeaderComponent />
          </div>
          <div class="row">
            <Switch>
              <Route path="/" exact component={ListUserComponent}></Route>
              <Route path="/user" component={ListUserComponent}></Route>
              <Route path="/add-user/:id" component={CreateUserComponent}></Route>
              <Route path="/view-user/:id" component={ViewUserComponent}></Route>
            </Switch>
          </div>
          <div class="row">
            <FooterComponent />
          </div>
        </div>    
      </Router>
    </div>
  );
}

export default App;
