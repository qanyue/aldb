import {defineStore} from 'pinia'
import {Operator} from "~/api/algae";
export const useOperatorStore = defineStore('operator', {
    state: () => {
        return{
            name:"" ,
            email:"",
            dataSet:[] as string[]
        }
    },
    getters: {
    },
    actions: {

    },
})