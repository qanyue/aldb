import http from "~/utils/request";

export interface Tag {
    name: string,
    resourceName: string
}

export interface Operator {
    name: string,
    email: string,
    dataSet: string[]
}

export interface Annotation {
    description: string,
    createAt?: string,
    updateAt?: string,
    segmentation: number[],
    tag: Tag,
}
export interface Alga {
    name: string
    src: string
    annotations: Annotation[]
}
export interface River {
    name: string,
    address: string
    algae?: string[]
}


//和RiverWIthid仅作命名区分
export interface RiverInfo {
    name: string,
    address: string
    id: string
}

export interface RiverWithId {
    id: string, //mongodb id
    name: string, //名称
    address: string //描述信息
    //search专用，无藻类数据
}


//TODO gwtalgaData 逻辑修改
export const getAlgaData = (algaId: string) => {
    return http.request({
        url: "/api/alga/get",
        method: "post",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        data: {
            "algaId": algaId
        },
    });
};

export const getAllAlga = (riverId: string) => {
    return http.request({
        url: "/api/alga/all?riverId=" + riverId,
        method: "get",
    });
};

export const addAlga = (data: FormData) => {
    return http.request({
        url: "/api/alga/add",
        method: "post",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        data,
    });
};

export const addAlgaMore = (data: FormData, riverId: string, preSuffix: string) => {
    data.append("riverId", riverId);
    data.append("preSuffix", preSuffix)
    return http.request({
        url: "/api/alga/addMore",
        method: "post",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        data
    });
};
export const searchAlga = (key: string, rId: string) => {
    var bodyFormData = new FormData();
    bodyFormData.append('key', key);
    bodyFormData.append('riverId', rId);
    return http.request({
        url: "/api/alga/search",
        method: "post",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        data: bodyFormData
    });
};
export const getAnno = (algaId: string) => {
    return http.request({
        url: "/api/alga/anno?algaId=" + algaId,
        method: "get",
    });
};

export const addAnno = (algaId: string, segmentation: number[], tag: Tag, description: string) => {
    return http.request({
        url: "/api/anno/add",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        method: "post",
        data: {
            "algaId": algaId,
            "tag": tag,
            "segmentation": segmentation,
            "description": description
        }
    });
};

export const deleteAnno = (algaId: string, annotation: Annotation) => {
    return http.request({
        url: "/api/anno/delete",
        method: "post",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        data: {
            "algaId": algaId,
            "tag": annotation.tag,
            "description": annotation.description
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

export const addRiver = (userEmail: string, River: River) => {
    return http.request({
        url: "/api/river/add",
        method: "post",
        data: {
            userEmail: userEmail,
            river: River
        }
    });
};

export const shareRiver = (userEmail: string, riverId: string) => {
    let formData = new FormData()
    formData.append("userEmail", userEmail)
    formData.append("riverId", riverId)
    return http.request({
        url: "/api/river/share",
        method: "post",
        data: formData
    })

}


export const getRivers = (userEmail: string) => {
    var bodyFormData = new FormData();
    bodyFormData.append("userEmail", userEmail)
    return http.request({
        url: "/api/river/all",
        method: "post",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        data: bodyFormData,
    });
};
export const searchRiverByKey = (userEmail: string, key: string) => {
    let bodyFormData = new FormData();

    bodyFormData.append("userEmail", userEmail)
    bodyFormData.append("key", key)
    return http.request({
        url: "/api/river/search",
        method: "post",
        headers: { 'Content-Type': 'application/form-x-www-form-urlencoded' },
        data: bodyFormData,
    });
}

export const getRiverInfo = (riverId: string) => {
    return http.request({
        url: "/api/river/info?riverId=" + riverId,
        method: "get",
        data: {
        },
    });
};
