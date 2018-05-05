import React, {Component} from 'react';
import './App.css';
import {BrowserRouter,Route} from 'react-router-dom';
import {Navbar, NavItem, Nav} from 'react-bootstrap';
import {LinkContainer} from 'react-router-bootstrap';
import Login from './components/Login';
import logo from './images/logo.png';
import image from './images/image.png';
import Register from "./components/Register";
import Menu from "./components/Menu";
import ViewCart from "./components/ViewCart";
import Checkout from "./components/Checkout";
import ViewTransaction from "./components/ViewTransaction";

class App extends Component {

  render() {
    return (

            <div>
                <Navbar inverse collapseOnSelect>

                    <Navbar.Header>
                        <Navbar.Brand>
                            <a href='/home' className="img-background">
                                <img src={image}/>
                            </a>
                        </Navbar.Brand>
                    </Navbar.Header>
                    <Navbar.Collapse>

                        {this.checkLoggedIn() && <Nav pullLeft>

                            <LinkContainer to="/home">
                                <NavItem eventKey={3}>
                                    Home
                                </NavItem>
                            </LinkContainer>
                            <LinkContainer to="/userProfile">
                                <NavItem eventKey={4}>
                                    Profile
                                </NavItem>
                            </LinkContainer>
                            <LinkContainer to="/dashboard">
                                <NavItem eventKey={5}>
                                    Dashboard
                                </NavItem>
                            </LinkContainer>

                        </Nav>}
                        {this.checkLoggedIn() && <Nav pullRight>

                            <span className="navbar-btn"><a className="btn post-a-project-btn" href="/postProject">Post a Project</a></span>
                                <span className="navbar-btn"><a className="btn btn-primary logout-btn" onClick={this.logOut.bind(this)}>Log Out</a></span>


                        </Nav>}
                        {this.checkLoggedIn() && <Nav pullRight>

                        </Nav>}


                        {!this.checkLoggedIn() && <Nav pullRight>

                            <LinkContainer to="/login">
                                <NavItem eventKey={1}>
                                    Login
                                </NavItem>
                            </LinkContainer>
                            <LinkContainer to="/register">
                                <NavItem eventKey={2}>
                                    Signup
                                </NavItem>
                            </LinkContainer>
                        </Nav>}
                    </Navbar.Collapse>
                </Navbar>
                <div>
                    <Route path="/login" render={()=><Login/>} />
                    <Route path="/register" render={()=><Register />}/>
                    <Route path="/menu" render={()=><Menu />}/>
                    <Route path="/viewcart" render={()=><ViewCart />}/>
                    <Route path="/checkout" render={()=><Checkout />}/>
                    <Route path="/viewtransaction" render={()=><ViewTransaction />}/>
                    {/*<Route path="/home" render={()=>{*/}
                        {/*if(this.checkLoggedIn())*/}
                            {/*return <Home /> ;*/}
                        {/*else{*/}
                            {/*this.props.history.push('/login');*/}
                            {/*return null;*/}
                        {/*}*/}
                    {/*}}/>*/}
                    {/*<Route path="/userProfile" render={()=><UserProfile loggedIn={this.checkLoggedIn.bind(this)}/>}/>*/}
                    {/*<Route path="/postProject" render={()=>{*/}
                        {/*if(this.checkLoggedIn())*/}
                            {/*return <PostProject /> ;*/}
                        {/*else{*/}
                            {/*this.props.history.push('/login');*/}
                            {/*return null;*/}
                        {/*}*/}
                    {/*}}/>*/}
                    {/*<Route path="/projectDetails/:pid" render={()=>{*/}
                        {/*if(this.checkLoggedIn())*/}
                            {/*return <ProjectDetails /> ;*/}
                        {/*else{*/}
                            {/*this.props.history.push('/login');*/}
                            {/*return null;*/}
                        {/*}*/}
                    {/*}}/>*/}
                    {/*<Route path="/dashboard" render={()=>{*/}
                        {/*if(this.checkLoggedIn())*/}
                            {/*return <Dashboard /> ;*/}
                        {/*else{*/}
                            {/*this.props.history.push('/login');*/}
                            {/*return null;*/}
                        {/*}*/}
                    {/*}}/>*/}
                </div>
            </div>
    );
  }

    checkLoggedIn(){
      console.log(document.cookie);
      if(localStorage.getItem('token')){
          return true;
      }
      return false;
    }
    logOut(){
        this.deleteAllCookies();
        localStorage.removeItem('token');
        this.props.history.push('/login');
    }

}

export default App;