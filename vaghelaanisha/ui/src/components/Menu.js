import React, {Component} from 'react';
import '../css/project.css';
import '../css/userBidList.css';
import {addOrder} from "../API/api";

class Menu extends Component {

    constructor(){
        super();
        this.state={
            menu:[
                {
                    _id: 1,
                    name: "mocha",
                    description: "mochhaaaaaaaaaaaaaaaaa",
                    price: 10,
                    qty: 0
                },
                {
                    _id:2,
                    name:"latte",
                    description:"lattteeeeeeeee",
                    price:10,
                    qty:0
                }            ],
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleAddCart = this.handleAddCart.bind(this);
    }

    handleChange(e,item){
        e.preventDefault();
        let {menu}=this.state;
        menu.forEach(function(menuItem){
            if(menuItem._id===item._id){
                menuItem.qty=parseInt(e.target.value);
            }
        });
        this.setState({menu});
    }

    handleAddCart(e,item){
        e.preventDefault();
        console.log("item:",item);

        addOrder(item)
            .then((data)=>{
                // localStorage.setItem("UserId",data.UserId);
                // this.history.push('/login')
            })
            .catch((err)=>{
                this.setState({error:"There is some error!"})
                console.log(err);
            })
    }

    render() {
        return (
            <div className={"container"}>

                <div className="row table-header">

                    <div className="col-md-6 col-xs-6">
                        ITEM
                    </div>

                    <div className="col-md-2 col-xs-2">
                        PRICE
                    </div>

                    <div className="col-md-2 col-xs-2">
                        QUANTITY
                    </div>

                    <div className="col-md-2 col-xs-2">
                    </div>

                </div>

                {this.state.menu.map((item,key)=>{
                    const {name,description,price,qty} = item;
                    return (
                        <div className="row project-item">

                            <div className="col-md-6">

                                <div className="row">
                                    <div className="col-md-12">
                                        <div className="project-name">{name ? name : ""}</div>
                                    </div>
                                </div>

                                <div className="row">
                                    <div className="col-md-12">
                                        <div className="project-description">{description ? description : ""}</div>
                                    </div>
                                </div>


                            </div>

                            <div className="col-md-2">

                                <div className="project-employer">{price ? price:""}</div>

                            </div>

                            <div className="col-md-2">

                                <div className="form-group">
                                    <input type="number" name="qty" onChange={(e)=>this.handleChange(e,item)}  value={qty} className="form-control" id="exampleInputEmail1" placeholder="Email or Username"/>
                                </div>
                            </div>

                            <div className="col-md-2">
                                <button className="btn btn-sm btn-primary" onClick={(e)=>this.handleAddCart(e,item)}>Add</button>
                            </div>

                        </div>
                    );
                } )}

            </div>

        );
    }
}

export default Menu;
