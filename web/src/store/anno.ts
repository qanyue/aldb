import {defineStore} from 'pinia'
import {Annotation, Tag} from "~/api/algae";

export const useAnnotationStore = defineStore('anno', {
    state: () => {
        return {
            algaId: "",
            anno: {
                tag: {
                    name: "",
                    resourceName: ""
                } as Tag
            }as Annotation
        }
    },
    getters: {},
    actions: {},
})