import React, {Component} from 'react';
import '../css/login.css'
import logo from '../images/logo.png';
import logo2 from '../images/logo2.png';
import {doLogin} from "../API/api";


class Login extends Component{

    constructor(){
        super();
        this.state={
            Username:"",
            Password:"",
            error:""
        };
        this.handleChange =this.handleChange.bind(this);
        this.handleUserLogin =this.handleUserLogin.bind(this);
        this.errorAlert = this.errorAlert.bind(this);

    }

    handleUserLogin(e){
        e.preventDefault();
        console.log(this.state);
        doLogin(this.state)
            .then((data)=>{
                localStorage.setItem("UserId",data.UserId);
                this.history.push('/menu')
            })
            .catch((err)=>{
                this.setState({error:"There is some error!"})
                console.log(err);
            })
        // this.props.userLogin(this.state);
    }

    handleChange(e){

        this.setState({[e.target.name]:e.target.value});
    }

    errorAlert(){
        if(this.state.error && this.state.error.length>0){
            return (
                <div className="alert alert-danger">{this.state.error}</div>
        )
        }
    }

    render(){
        return(

            <div className="back">

                <div className="div-center">

                    <div className="content">

                        <div>
                            <img className="center-block" src={logo2}/>
                            <img src={logo}/>
                        </div>
                        <br/>
                        <hr />
                        <br/>
                        <form onSubmit={this.handleUserLogin}>

                            <div className="form-group">
                                <input type="text" name="Username" onChange={this.handleChange}  value={this.state.Username} className="form-control" id="exampleInputEmail1" placeholder="Email or Username"/>
                            </div>

                            <div className="form-group">
                                <input type="password" name="Password" onChange={this.handleChange} value={this.state.Password} className="form-control" id="exampleInputPassword1" placeholder="Password"/>
                            </div>

                            <button type="submit" className="btn btn-block btn-primary">Login</button>

                            <hr />
                            <span>Don't have an account? </span>
                            <a href="/register">Signup</a>
                            <p>{this.errorAlert()}</p>
                        </form>
                    </div>
                </div>
            </div>
        )
    }
}

export default Login;
