import React, {Component} from 'react';
import '../css/login.css'
import logo from '../images/logo.png';
import {doCheckout} from "../API/api";

class Checkout extends Component{

    constructor(){
        super();
        this.state={
            cardNumber:"",
            name:"",
            expMonth:"",
            expYear:"",
            cvv:"",
        }
        this.handleChange =this.handleChange.bind(this);
        this.handlePay =this.handlePay.bind(this);

    }

    handlePay(){
        console.log(this.state);
        doCheckout(this.state)
            .then((data)=>{
                this.history.push('/viewtransaction')
            })
            .catch((err)=>{
                console.log(err);
            })
    }

    handleChange(e){

        this.setState({[e.target.name]:e.target.value});
    }

    render(){
        return(
            <div className="back">

                <div className="div-center">

                    <div className="content">

                        <form onSubmit={this.handlePay}>

                            <div className="form-group">
                                <input type="text" name="cardnumber" required value={this.state.cardNumber} onChange={this.handleChange} className="form-control"  placeholder="Card Number *"/>
                            </div>

                            <div className="form-group">
                                <input type="text" name="name" required onChange={this.handleChange} value={this.state.name} className="form-control" placeholder="Name on Card *"/>
                            </div>

                            <div className="row form-group">
                                <div className="col-md-6">
                                    <input type="text"  name="expMonth" required onChange={this.handleChange} value={this.state.expMonth} className="form-control" placeholder="Expiry Month *"/>
                                </div>
                                <div className="col-md-6">
                                    <input type="text" name="expYear" required onChange={this.handleChange} value={this.state.expYear} className="form-control" placeholder="Expiry Year *"/>
                                </div>
                            </div>

                            <div className="form-group">
                                <input type="text" name="cvv" required onChange={this.handleChange} value={this.state.cvv} className="form-control" placeholder="CVV *"/>
                            </div>

                            <button type="submit" className="btn btn-block btn-primary">Pay</button>

                        </form>
                    </div>
                </div>
            </div>

        )
    }
}

export default Checkout;
