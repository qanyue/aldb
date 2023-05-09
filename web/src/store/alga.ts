import {defineStore} from 'pinia'
import {Annotation} from "~/api/algae";

export const useAlgaStore = defineStore('alga', {
    state: () => {
        return{
            algaId: "",
            algaInfo:{
                name: "",
                url:"",
                annotations:[] as Annotation[]
            }
        }
    },
    getters: {
    },
    actions: {
    },
})