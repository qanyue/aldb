import {defineStore} from 'pinia'
import {Operator, River, RiverInfo} from "~/api/algae";
export const useOperatorStore = defineStore('operator', {
    state: () => {
        return{
            name:"" ,
            email:"",
            dataSet:[] as RiverInfo[]
        }
    },
    getters: {
    },
    actions: {
        

    },
})