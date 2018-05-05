const headers = {
    'Accept': 'application/json',
    'Credentials':'include'
};

kong_url=process.env.KONG_URL || "54.183.59.125:8000";

export const doLogin = (payload) =>
    fetch(kong_url+`user/login`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
            console.log(error);
            return error;
    });

export const doSignUp = (payload) =>
    fetch(kong_url+`user/signup`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log(error);
        console.log("This is error");
        return error;
    });

export const doCheckout = (payload) =>
    return fetch(kong_url+`/checkout/checkoutCart`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    }).then(res => {
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });

export const getOrder = (payload) => {
    console.log(payload);
    return fetch(kong_url+`cart/order`+UserId, {
        method: 'GET',
        headers: {
            ...headers,
        },
        body: JSON.stringify(payload)
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log(error);
        return error;
    });
}

export const updateOrder = () =>
    fetch(kong_url+`/cart/order/`, {
        method: 'PUT',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload)
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });

export const addOrder = (payload) =>
    fetch(kong_url+`/cart/order/`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload)
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });