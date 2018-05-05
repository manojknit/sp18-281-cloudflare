import React, {Component} from 'react';
import '../css/login.css'
import logo from '../images/logo.png';
import logo2 from '../images/logo2.png';


class Login extends Component{

    constructor(){
        super();
        this.state={
            username:"",
            password:"",
            error:""
        };
        this.handleChange =this.handleChange.bind(this);
        this.handleUserLogin =this.handleUserLogin.bind(this);
        this.errorAlert = this.errorAlert.bind(this);

    }

    handleUserLogin(e){
        e.preventDefault();
        console.log(this.state);
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
                                <input type="text" name="username" onChange={this.handleChange}  value={this.state.username} className="form-control" id="exampleInputEmail1" placeholder="Email or Username"/>
                            </div>

                            <div className="form-group">
                                <input type="password" name="password" onChange={this.handleChange} value={this.state.password} className="form-control" id="exampleInputPassword1" placeholder="Password"/>
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
