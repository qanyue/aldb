import { stringify } from "querystring";
import http from "~/utils/request";
interface Tag{
    name:string
    resourceName:string
}

export const addTag = (data: Tag) => {
    return http.request({
        url: "/api/tag/add",
        method: "post",
        data:{
            name:data.name,
            resourceName:data.resourceName
        },
    });
};

export const deleteTag = (data: Tag) => {
    return http.request({
        url: "/api/tag/delete",
        method: "post",
        data:{
            name:data.name,
            resourceName:data.resourceName
        },
    });
}

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