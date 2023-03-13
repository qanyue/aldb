import http from "~/utils/request";

export const uploadFile = (data: object) => {
    return http.request({
        url: "/api/upload",
        method: "post",
        data,
    });
};