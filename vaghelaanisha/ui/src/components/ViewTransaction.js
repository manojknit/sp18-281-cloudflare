import React, {Component} from 'react';
import '../css/project.css';
import '../css/userBidList.css';

class ViewTransaction extends Component {

    constructor(){
        super();
        this.state={
            orders:[
                [{
                    _id:1,
                    name:"mocha",
                    description:"mochhaaaaaaaaaaaaaaaaa",
                    price:10,
                    qty:10
                },
                {
                    _id:1,
                    name:"latte",
                    description:"lattteeeeeeeeee",
                    price:10,
                    qty:10
                }],
                [{
                    _id:1,
                    name:"mocha",
                    description:"mochhaaaaaaaaaaaaaaaaa",
                    price:10,
                    qty:10
                }]
            ],
        }
        this.handleChange = this.handleChange.bind(this);
        this.handlePlaceOrder = this.handlePlaceOrder.bind(this);
        this.handleTotalPrice = this.handleTotalPrice.bind(this);
    }

    handleChange(e,item){
        e.preventDefault();
        let {order}=this.state;
        order.forEach(function(orderItem){
            if(orderItem._id===item._id){
                orderItem.qty=parseInt(e.target.value);
            }
        });
        this.setState({order});
    }

    handlePlaceOrder(){
        console.log(this.state);
    }

    handleTotalPrice(order){
        let totalPrice=0;
        order.forEach(function(orderItem){
            totalPrice+=orderItem.price*orderItem.qty;
        });
        console.log("totalPrice",totalPrice);
        return totalPrice;
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
                        COST
                    </div>



                </div>

                {this.state.orders.map((order)=>{
                    console.log(order);
                    {var temp =order.map((item,key)=>{
                        console.log(item);
                        const {name,description,price,qty} = item;
                        console.log("name:",name);
                        return (
                            <div className="row project-item">

                                <div className="col-md-6">

                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="project-name">{name ? name : ""}</div>
                                            {console.log("name1:",name)}
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
                                    <div className="project-employer">{qty ? qty:""}</div>
                                </div>


                                <div className="col-md-2">
                                    <div className="project-employer">{price && qty ? price*qty:""}</div>
                                </div>

                            </div>
                        );
                    } )}
                    return(
                        <div className="col-md-12">
                            {temp}
                            <div className="col-md-offset-10 col-md-2" style={{fontSize:15+"px"}}><strong>
                                Total Price : {this.handleTotalPrice(order)}</strong>
                            </div>
                        </div>
                    );
                })}
            </div>

        );
    }
}

export default ViewTransaction;
