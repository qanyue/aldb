import http from "~/utils/request";

export const registerUser = (data: object) => {
    return http.request({
        url: "/register",
        method: "post",
        data,
    });
};

export const getUser = (query: string) => {
    return http.request({
        url: "/api/user/info?email=" + query,
        method: "get"
    });
};

export const getUsers = () => {
    return http.request({
        url: "/api/user/all",
        method: "get"
    });
};


export const updateUser = (data: object) => {
    return http.request({
        url: "/api/user/update",
        method: "post",
        data,
    });
};

export const getAnno = (query: string) => {
    return http.request({
        url: "/api/user/anno?user=" + query,
        method: "get"
    });
};

export const deleteUser = (query: string) => {
    return http.request({
        url: "/api/user/delete?email=" + query,
        method: "get"
    });
};

export const changePassword = (data: object) => {
    return http.request({
        url: "/api/user/pwd",
        method: "post",
        data,
    });
};
