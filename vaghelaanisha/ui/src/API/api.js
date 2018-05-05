const headers = {
    'Accept': 'application/json',
    'Credentials':'include'
};

export const doLogin = (payload) =>
    fetch(`/userlogin`, {
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

export const doRegister = (payload) =>
    fetch(`/userregister`, {
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

export const userProfile = (token) =>
    fetch(`/authToken/profile`, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'x-access-token':token
        },
    }).then(res => {
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });

export const postProject = (token,payload) => {
    console.log(payload);
    return fetch(`/authToken/postProject`, {
        method: 'POST',
        headers: {
            ...headers,
            'x-access-token': token
        },
        body: payload
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log(error);
        return error;
    });
}

export const getProjectDetails = (token,pid) =>
    fetch(`/authToken/projectDetails/`+pid, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'x-access-token':token
        },
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });

export const placeBid = (token,payload) => {
    console.log(payload);
    return fetch(`/authToken/placeBid`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'x-access-token': token
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

export const getHome = (token) =>
    fetch(`/authToken/home/`, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'x-access-token':token
        },
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });

export const getEmployerProjects = (token) =>
    fetch(`/authToken/employerProjects/`, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'x-access-token':token
        },
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });

export const getFreelancerProjects = (token) =>
    fetch(`/authToken/freelancerProjects/`, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'x-access-token':token
        },
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log("This is error");
        return error;
    });

export const profileEdit = (token,payload) => {
    console.log(payload);
    return fetch(`/authToken/profileEdit`, {
        method: 'POST',
        headers: {
            ...headers,
            'x-access-token': token
        },
        body: payload
    }).then(res => {
        console.log(res);
        return res.json();
    }).catch(error => {
        console.log(error);
        return error;
    });
}
