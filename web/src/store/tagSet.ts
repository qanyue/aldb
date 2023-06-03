import {defineStore} from 'pinia'
import { Tag } from '~/api/algae'

export const useTagStore = defineStore('tags', {
    state: () => {
        return{
            tags: [] as Tag[],
        }
    },
    getters: {
    },
    actions: {
    },
})