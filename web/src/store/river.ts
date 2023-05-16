import {defineStore} from 'pinia'

export const useRiverStore = defineStore('river', {
    state: () => {
        return{
            id:"",
            name: "",
            address: "",
            algae: [] as string[],
        }
    },
    getters: {
    },
    actions: {
    },
})