import React, {Component} from 'react';
import '../css/login.css'
import logo from '../images/logo.png';
import logo2 from '../images/logo2.png';

class Register extends Component{

    constructor(){
        super();
        this.state={
            username:"",
            password:"",
            confirmPassword:"",
            error:"",
        }
        this.handleChange =this.handleChange.bind(this);
        this.handleRegisterUser =this.handleRegisterUser.bind(this);
        this.errorAlert = this.errorAlert.bind(this);

    }

    handleRegisterUser(e){
        e.preventDefault();
        console.log(this.state);
        const {password,confirmPassword} = this.state;
        if(password && confirmPassword && password!==confirmPassword){
            this.setState({error:"Password and confirm password do not match"});
        }
        else{
            // this.props.userRegistration(this.state);
            this.setState({error:""});
        }

    }

    handleChange(e){

        this.setState({[e.target.name]:e.target.value});
    }

    errorAlert(){
        if(this.state.error && this.state.error.length>0){
            return(
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
                        <hr />
                        <form onSubmit={this.handleRegisterUser}>

                            <div className="form-group">
                                <input type="text" name="username" required value={this.state.username} onChange={this.handleChange} className="form-control"  placeholder="Username *"/>
                            </div>

                            <div className="form-group">
                                <input type="password" name="password" required onChange={this.handleChange} value={this.state.password} className="form-control" placeholder="Password *"/>
                            </div>

                            <div className="form-group">
                                <input type="password" name="confirmPassword" required onChange={this.handleChange} value={this.state.confirmPassword} className="form-control" placeholder="Confirm Password *"/>
                            </div>

                            <button type="submit" className="btn btn-block btn-primary">Create Account</button>

                            <hr />
                            <span>Already a Freelancer.com member? </span>
                            <a href="/login" >Log In</a>
                            <p>{this.errorAlert()}</p>
                        </form>
                    </div>
                </div>
            </div>

        )
    }
}

export default Register;
