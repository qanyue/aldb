import http from "~/utils/request";

export const getData = (data: object) => {
    return http.request({
        url: "/api/data",
        method: "get",
        data,
    });
};

export const addAlga = (data: object) => {
    return http.request({
        url: "/api/alga/add",
        method: "post",
        data,
    });
};

export const searchAlga = (query: string) => {
    return http.request({
        url: "/api/alga/search?k=" + query,
        method: "get",
    });
};

export const getAnno = (query: string) => {
    return http.request({
        url: "/api/alga/anno?alga=" + query,
        method: "get",
    });
};

export const addAnno = (data: object) => {
    return http.request({
        url: "/api/anno/add",
        method: "post",
        data,
    });
};

export const deleteAnno = (query: string) => {
    return http.request({
        url: "/api/anno/delete?id=" + query,
        method: "get",
    });
};

export const updateAnno = (data: object) => {
    return http.request({
        url: "/api/anno/update",
        method: "post",
        data,
    });
};

export const addRiver = (data: object) => {
    return http.request({
        url: "/api/river/add",
        method: "post",
        data,
    });
};

export const getRivers = (data: object) => {
    return http.request({
        url: "/api/river/all",
        method: "get",
        data,
    });
};
