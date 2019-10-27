import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import 'antd/dist/antd.css';
// import { Router, Route } from 'react-router'
import {HashRouter, Route, Switch, hashHistory} from 'react-router-dom';
import Login from './pages/login'
import Users from './pages/Users'
import NoMatch from './pages/NoMatch'
// ReactDOM.render(<App />, document.getElementById('root'));
import { createHashHistory } from 'history';
const history = createHashHistory();
console.log("history", history)
ReactDOM.render((
    // <Router>
    //   <Route path="/" component={App}>
    //     <Route path="Login" component={Login}/>
    //     <Route path="users" component={Users}>
    //       {/* <Route path="/user/:userId" component={User}/> */}
    //     </Route>
    //     <Route path="*" component={NoMatch}/>
    //   </Route>
    // </Router>
    <HashRouter history={history}>
        <Switch>
            <Route exact  path="/" component={App}>
            </Route>
            <Route exact  path="/login" component={Login}></Route>
            <Route exact  path="/users" component={Users}>
            </Route>
            <Route exact  path="*" component={NoMatch}/>
        </Switch>
    </HashRouter>
), document.body)
serviceWorker.unregister();
