import { url } from "inspector";
import { encode } from "punycode";
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
        url: "/api/user/info?email=" + encodeURIComponent(query),
        headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
        method: "get"
    });
};

export const getUsers = (userEmail:string) => {
    return http.request({
        url: "/api/user/all?UserEmail="+encodeURIComponent(userEmail),
        headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
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

//弃用
/*
export const getAnno = (query: string) => {
    return http.request({
        url: "/api/user/anno?user=" + query,
        method: "get"
    });
};

*/
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
