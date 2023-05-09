import {defineStore} from 'pinia'

export const useRiverStore = defineStore('river', {
    state: () => {
        return{
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