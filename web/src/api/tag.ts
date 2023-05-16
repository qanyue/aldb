import http from "~/utils/request";
interface Tag{
    name:string
    resourceName:string
}

export const addTag = (data: Tag) => {
    return http.request({
        url: "/api/tag/add",
        method: "post",
        data,
    });
};

export const getAllTags = () => {
    var bodyFormData = new FormData();
    return http.request({
        url: "/api/tag/all",
        method: "get",
        // headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
        data:{

        },
    });
};