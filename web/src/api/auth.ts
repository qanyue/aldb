import http from "~/utils/request";

export const getCaptcha = (data: object) => {
    return http.request({
        url: "/captcha",
        method: "get",
        data,
    });
};

export const loginSubmit = (data: object) => {
    return http.request({
        url: "/auth",
        method: "post",
        data,
    });
};
