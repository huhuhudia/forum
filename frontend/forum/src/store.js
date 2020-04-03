
import Vuex from 'vuex'
import Vue from 'vue'
Vue.use(Vuex)

const store = new Vuex.Store({
    state:{
        isLogin:false
    },
    mutations:{
        login(state){
            state.isLogin = true
        }
    }
})

export default store