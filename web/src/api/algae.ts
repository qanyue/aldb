import http from "~/utils/request";

export interface Tag {
    name: string,
    resourceName: string
}
export interface Operator{
    name:string,
    email:string,
    dataSet: string[]
}

export interface Annotation {
    description: string,
    createAt?: string,
    updateAt?: string,
    tag: Tag,
}
export interface Alga {
    name: string
    src: string
    annotations: Annotation[]
}
export interface River {
    name:string,
    address:string
}

//TODO gwtalgaData 逻辑修改
export const getAlgaData = (algaId:string) => {
    return http.request({
        url: "/api/alga/get",
        method: "post",
        headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
        data:{
            "algaId":algaId
        },
    });
};

export const getAllAlga = (riverId:string) => {
    return http.request({
        url: "/api/alga/get?riverId="+riverId,
        method: "get",
    });
};

export const addAlga = (data: object) => {
    return http.request({
        url: "/api/alga/add",
        method: "post",
        headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
        data,
    });
};

export const addAlgaMore = (data: object) => {
    return http.request({
        url: "/api/alga/addMore",
        method: "post",
        headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
        data,
    });
};
//TODO 逻辑为修改为 let formdata = new FormData()
export const searchAlga = (key: string, rId: string) => {
    return http.request({
        url: "/api/alga/search",
        method: "post",
        headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
        data: {
            "key": key,
            "riverId": rId
        }
    });
};
//
export const getAnno = (query: string) => {
    return http.request({
        url: "/api/alga/anno?algaId=" + query,
        method: "get",
    });
};

//TODO修改逻辑
export const addAnno = (algaId:string, anno:Annotation) => {
    return http.request({
        url: "/api/anno/add",
        headers: {'Content-Type': 'application/form-x-www-form-urlencoded'},
        method: "post",
        data: {
            "algaId": algaId,
            "tag":anno.tag,
            "description":anno.description
        }
    });
};

//TODO 修改数据格式
export const deleteAnno = (algaId: string, annotation: Annotation) => {
    return http.request({
        url: "/api/anno/delete",
        method: "post",
        data: {
            "algaId": algaId,
            "tag":annotation.tag,
            "description":annotation.description
        }
    });
};

//弃用
/*export const updateAnno = (data: object) => {
    return http.request({
        url: "/api/anno/update",
        method: "post",
        data,
    });
};*/

//TODO 修改data格式
export const addRiver = (data: object) => {
    return http.request({
        url: "/api/river/add",
        method: "post",
        data,
    });
};

//TODO 修改使用的参数设计
export const getRivers = (userEmail: string) => {
    return http.request({
        url: "/api/river/all",
        method: "post",
        data: {
            "userEmail": userEmail
        },
    });
};
